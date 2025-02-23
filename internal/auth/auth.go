// Slurp
// Copyright (C) Vyr Cossont
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package auth

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	neturl "net/url"
	"os"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	gokeyring "github.com/zalando/go-keyring"
	"golang.org/x/time/rate"
	"webfinger.net/go/webfinger"

	apiclient "github.com/VyrCossont/slurp/client"
	"github.com/VyrCossont/slurp/client/apps"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

// Client is a GtS API client with attached authentication credentials and rate limiter.
// Credentials may be no-op.
type Client struct {
	Client  *apiclient.GoToSocialSwaggerDocumentation
	Auth    runtime.ClientAuthInfoWriter
	limiter *rate.Limiter
	ctx     context.Context
}

func (c *Client) Wait() error {
	if err := c.limiter.Wait(c.ctx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func NewAuthClient(user string) (*Client, error) {
	var err error

	keyring := util.SystemKeyring
	useCleartextFileKeyring, err := util.GetUseCleartextFileKeyring()
	if err != nil {
		slog.Error("couldn't get keyring prefs")
		return nil, err
	}
	if useCleartextFileKeyring {
		keyring = util.CleartextFileKeyring
	}

	if user == "" {
		user, err = util.GetDefaultUser()
		if err != nil {
			slog.Error("no user provided, couldn't get default user from prefs (did you log in first?)")
			return nil, err
		}
	}

	instance, err := util.GetUserInstance(user)
	if err != nil {
		slog.Error("couldn't get user's instance from prefs (did you log in first?)", "user", user)
		return nil, err
	}

	accessToken, err := keyring.Get(keyringServiceAccessToken, user)
	if err != nil {
		slog.Error("couldn't find user's access token (did you log in first?)", "user", user, "instance", instance)
		return nil, err
	}

	client, err := clientForInstance(instance)
	if err != nil {
		slog.Error("couldn't get client for user's instance", "user", user, "instance", instance)
		return nil, err
	}

	limiter, err := rateLimiterForInstance(instance)
	if err != nil {
		slog.Error("couldn't get rate limiter for user's instance", "user", user, "instance", instance)
		return nil, err
	}

	return &Client{
		Client:  client,
		Auth:    httptransport.BearerToken(accessToken),
		limiter: limiter,
		ctx:     context.Background(),
	}, nil
}

const (
	keyringServiceAccessToken  = "codes.catgirl.slurp.access-token"
	keyringServiceClientSecret = "codes.catgirl.slurp.client-secret"
)

const (
	oauthRedirect = "urn:ietf:wg:oauth:2.0:oob"
	oauthScopes   = "read write"
)

// Login authenticates the user and saves the credentials in the system keychain.
func Login(user string, allowHTTP bool, useCleartextFileKeyring bool) error {
	var err error

	keyring := util.SystemKeyring
	if useCleartextFileKeyring {
		err = util.SetUseCleartextFileKeyring(true)
		if err != nil {
			slog.Error("couldn't set keyring prefs")
			return err
		}
	} else {
		useCleartextFileKeyring, err = util.GetUseCleartextFileKeyring()
		if err != nil {
			slog.Error("couldn't get keyring prefs")
			return err
		}
	}
	if useCleartextFileKeyring {
		keyring = util.CleartextFileKeyring
	}

	if user == "" {
		user, err = util.GetDefaultUser()
		if err != nil {
			slog.Error("no user provided, couldn't get default user from prefs (have you logged in before?)")
			return err
		}
	}

	if user == "" {
		return errors.WithStack(errors.New("a user is required"))
	}
	if !strings.ContainsRune(user, '@') {
		return errors.WithStack(errors.New("a fully qualified user with a domain is required"))
	}
	if user[0] == '@' {
		return errors.WithStack(errors.New("take the leading @ off the user and try again"))
	}

	if _, err := keyring.Get(keyringServiceAccessToken, user); err == nil {
		slog.Warn("already logged in, will log in again", "user", user)
	}

	instance, err := ensureInstance(user, allowHTTP)
	if err != nil {
		slog.Error("couldn't get user's instance", "user", user, "error", err)
		return err
	}

	client, err := clientForInstance(instance)
	if err != nil {
		slog.Error("couldn't get client for user's instance", "user", user, "instance", instance, "error", err)
		return err
	}
	clientID, clientSecret, err := ensureAppCredentials(instance, keyring, client)
	if err != nil {
		slog.Error("OAuth2 app setup failed", "user", user, "instance", instance, "error", err)
		return err
	}

	code, err := promptForOAuthCode(instance, clientID)
	if err != nil {
		slog.Error("couldn't prompt for OAuth2 authorization code", "user", user, "instance", instance, "error", err)
		return err
	}

	accessToken, err := exchangeCodeForToken(instance, clientID, clientSecret, code)
	if err != nil {
		slog.Error("couldn't exchange OAuth2 authorization code for access token", "user", user, "instance", instance, "error", err)
		return err
	}

	err = keyring.Set(keyringServiceAccessToken, user, accessToken)
	if err != nil {
		slog.Error("couldn't set access token in keychain", "user", user, "instance", instance, "error", err)
		return err
	}

	err = util.SetDefaultUser(user)
	if err != nil {
		slog.Error("couldn't set default user in prefs", "user", user, "instance", instance, "error", err)
		return err
	}

	slog.Info("login successful", "user", user, "instance", instance)

	return nil
}

// ensureInstance finds a user's instance or retrieves a previously cached instance for them.
func ensureInstance(user string, allowHTTP bool) (string, error) {
	if instance, err := util.GetUserInstance(user); err == nil {
		return instance, nil
	}

	instance, scheme, err := findInstance(user, allowHTTP)
	if err != nil {
		slog.Error("WebFinger lookup failed", "user", user, "error", err)
		return "", err
	}

	err = util.SetInstanceScheme(instance, scheme)
	if err != nil {
		slog.Error("couldn't set instance scheme in prefs", "user", user, "instance", instance, "error", err)
		return "", err
	}

	err = util.SetUserInstance(user, instance)
	if err != nil {
		slog.Error("couldn't set instance in prefs", "user", user, "instance", instance, "error", err)
		return "", err
	}

	return instance, nil
}

// findInstance does a WebFinger lookup to find the domain (with optional port)
// and scheme of the instance API for a given user.
func findInstance(user string, allowHTTP bool) (string, string, error) {
	webfingerClient := webfinger.NewClient(nil)
	webfingerClient.AllowHTTP = allowHTTP
	jrd, err := webfingerClient.Lookup("acct:"+user, nil)
	if err != nil {
		return "", "", fmt.Errorf("WebFinger lookup failed: %w", err)
	}

	var href string
	for _, link := range jrd.Links {
		if link.Rel == "self" && link.Type == "application/activity+json" {
			href = link.Href
			break
		}
	}
	if href == "" {
		return "", "", errors.New("no link with rel=\"self\" and type=\"application/activity+json\"")
	}

	url, err := neturl.Parse(href)
	if err != nil {
		return "", "", err
	}
	if url.Hostname() == "" {
		return "", "", errors.New("unexpected URL format")
	}

	switch url.Scheme {
	case "https":
		if url.Port() == "" || url.Port() == "443" {
			return url.Hostname(), url.Scheme, nil
		} else {
			return url.Hostname() + ":" + url.Port(), url.Scheme, nil
		}

	case "http":
		if !allowHTTP {
			return "", "", errors.Errorf("unexpected URL scheme: %s", url.Scheme)
		}

		if url.Port() == "" || url.Port() == "80" {
			return url.Hostname(), url.Scheme, nil
		} else {
			return url.Hostname() + ":" + url.Port(), url.Scheme, nil
		}

	default:
		return "", "", errors.Errorf("unexpected URL scheme: %s", url.Scheme)
	}
}

// clientForInstance returns an OpenAPI client for the instance.
func clientForInstance(instance string) (*apiclient.GoToSocialSwaggerDocumentation, error) {
	scheme, err := util.GetInstanceScheme(instance)
	if err != nil {
		slog.Error("couldn't get instance URL scheme from prefs", "instance", instance, "error", err)
		return nil, err
	}

	return apiclient.New(httptransport.New(instance, "", []string{scheme}), strfmt.Default), nil
}

// rateLimiterForInstance returns a rate limiter for the instance.
func rateLimiterForInstance(instance string) (*rate.Limiter, error) {
	rateLimit, err := util.GetInstanceRateLimit(instance)
	if err != nil {
		slog.Error("couldn't get instance rate limit from prefs", "instance", instance)
		return nil, err
	}

	burstCap, err := util.GetInstanceBurstCap(instance)
	if err != nil {
		slog.Error("couldn't get instance burst capacity from prefs", "instance", instance)
		return nil, err
	}

	return rate.NewLimiter(rate.Limit(rateLimit), burstCap), nil
}

// ensureAppCredentials retrieves or creates and stores app credentials.
func ensureAppCredentials(instance string, keyring gokeyring.Keyring, client *apiclient.GoToSocialSwaggerDocumentation) (string, string, error) {
	shouldCreateNewApp := false

	clientID, err := util.GetInstanceClientID(instance)
	if clientID == "" || errors.Is(err, gokeyring.ErrNotFound) {
		shouldCreateNewApp = true
	} else if err != nil {
		slog.Error("couldn't get client ID from prefs", "instance", instance, "error", err)
		return "", "", err
	}

	clientSecret, err := keyring.Get(keyringServiceClientSecret, instance)
	if clientSecret == "" || errors.Is(err, gokeyring.ErrNotFound) {
		shouldCreateNewApp = true
	} else if err != nil {
		slog.Error("couldn't get client secret from keychain", "instance", instance, "error", err)
		return "", "", err
	}

	if !shouldCreateNewApp {
		return clientID, clientSecret, nil
	}

	app, err := createApp(client)
	if err != nil {
		slog.Error("couldn't create OAuth2 app", "instance", instance, "error", err)
		return "", "", err
	}
	clientID = app.ClientID
	clientSecret = app.ClientSecret

	err = util.SetInstanceClientID(instance, clientID)
	if err != nil {
		slog.Error("couldn't set client ID in prefs", "instance", instance, "error", err)
		return "", "", err
	}
	err = keyring.Set(keyringServiceClientSecret, instance, clientSecret)
	if err != nil {
		slog.Error("couldn't set client secret in keychain", "instance", instance, "error", err)
		return "", "", err
	}

	return clientID, clientSecret, nil
}

// createApp registers a new OAuth2 application.
func createApp(client *apiclient.GoToSocialSwaggerDocumentation) (*models.Application, error) {
	resp, err := client.Apps.AppCreate(
		&apps.AppCreateParams{
			ClientName:   "slurp",
			RedirectURIs: oauthRedirect,
			Scopes:       util.Ptr(oauthScopes),
			Website:      util.Ptr("https://catgirl.codes/slurp"),
		},
		func(op *runtime.ClientOperation) {
			op.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
		},
	)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

func promptForOAuthCode(instance string, clientID string) (string, error) {
	scheme, err := util.GetInstanceScheme(instance)
	if err != nil {
		slog.Error("couldn't get instance URL scheme from prefs", "instance", instance, "error", err)
		return "", err
	}

	oauthAuthorizeURL := (&neturl.URL{
		Scheme: scheme,
		Host:   instance,
		Path:   "/oauth/authorize",
		RawQuery: neturl.Values{
			"response_type": []string{"code"},
			"client_id":     []string{clientID},
			"redirect_uri":  []string{oauthRedirect},
			"scope":         []string{oauthScopes},
		}.Encode(),
	}).String()
	err = browser.OpenURL(oauthAuthorizeURL)
	if err != nil {
		slog.Warn("couldn't open browser to authorize", "error", err)
		print("Please open this URL in your browser:", oauthAuthorizeURL)
	}

	print("Enter authorization code: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	code := strings.TrimSpace(scanner.Text())

	return code, nil
}

type oauthTokenOK struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	CreatedAt   int64  `json:"created_at"`
}

type oauthTokenError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// exchangeCodeForToken exchanges an authorization code for an access token.
func exchangeCodeForToken(instance string, clientID string, clientSecret string, code string) (string, error) {
	scheme, err := util.GetInstanceScheme(instance)
	if err != nil {
		slog.Error("couldn't get instance URL scheme from prefs", "instance", instance, "error", err)
		return "", err
	}

	oauthTokenURL := (&neturl.URL{
		Scheme: scheme,
		Host:   instance,
		Path:   "/oauth/token",
	}).String()

	// TODO: add this to GtS Swagger doc
	resp, err := http.Post(oauthTokenURL, "application/x-www-form-urlencoded", strings.NewReader(neturl.Values{
		"grant_type":    []string{"authorization_code"},
		"code":          []string{code},
		"client_id":     []string{clientID},
		"client_secret": []string{clientSecret},
		"redirect_uri":  []string{oauthRedirect},
		"scope":         []string{oauthScopes},
	}.Encode()))
	if err != nil {
		slog.Error("call to OAuth2 token endpoint failed", "instance", instance, "error", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		var payload oauthTokenError
		err = json.NewDecoder(resp.Body).Decode(&payload)
		if err != nil {
			slog.Error("couldn't decode OAuth2 token endpoint error response", "instance", instance, "error", err)
			return "", err
		}
		return "", errors.WithStack(errors.New(payload.ErrorDescription))
	}

	var payload oauthTokenOK
	err = json.NewDecoder(resp.Body).Decode(&payload)
	if err != nil {
		slog.Error("couldn't decode OAuth2 token endpoint success response", "instance", instance, "error", err)
		return "", err
	}

	if payload.TokenType != "Bearer" {
		err = errors.WithStack(errors.New("unknown access token type"))
		slog.Error("unexpected response from OAuth2 token endpoint", "instance", instance, "token_type", payload.TokenType)
		return "", err
	}

	if payload.Scope != oauthScopes {
		err = errors.WithStack(errors.New("scopes are not what we asked for"))
		slog.Error("unexpected response from OAuth2 token endpoint", "instance", instance, "scopes", payload.Scope)
		return "", err
	}

	return payload.AccessToken, nil
}

func Logout(user string) error {
	// TODO: revoke token
	// TODO: remove token from keychain
	// TODO: remove client ID from keychain
	// TODO: remove client secret from keychain
	return errors.New("NOT IMPLEMENTED")
}

func Whoami() (string, error) {
	user, err := util.GetDefaultUser()
	if err != nil {
		slog.Error("no user provided, couldn't get default user from prefs (have you logged in before?)")
		return "", err
	}

	return user, nil
}

func Switch(user string) error {
	if user == "" {
		return errors.New("no user provided")
	}

	return util.SetDefaultUser(user)
}
