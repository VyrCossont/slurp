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
	"github.com/zalando/go-keyring"
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
	return c.limiter.Wait(c.ctx)
}

func NewAuthClient(user string) (*Client, error) {
	instance, err := keyring.Get(keyringServiceInstance, user)
	if err != nil {
		slog.Error("couldn't get user's instance from keychain (did you log in first?)", "user", user)
		return nil, err
	}

	accessToken, err := keyring.Get(keyringServiceAccessToken, user)
	if err != nil {
		slog.Error("couldn't find user's access token (did you log in first?)", "user", user)
		return nil, err
	}

	return &Client{
		Client:  clientForInstance(instance),
		Auth:    httptransport.BearerToken(accessToken),
		limiter: rate.NewLimiter(1.0, 300),
		ctx:     context.Background(),
	}, nil
}

const (
	keyringServiceAccessToken  = "codes.catgirl.slurp.access-token"
	keyringServiceInstance     = "codes.catgirl.slurp.instance"
	keyringServiceClientID     = "codes.catgirl.slurp.client-id"
	keyringServiceClientSecret = "codes.catgirl.slurp.client-secret"
)

const (
	oauthRedirect = "urn:ietf:wg:oauth:2.0:oob"
	oauthScopes   = "read write"
)

// Login authenticates the user and saves the credentials in the system keychain.
func Login(user string) error {
	var err error

	// TODO: store a default user in prefs after successful authentication
	// TODO: try load it here if the user is empty
	if user == "" {
		return errors.WithStack(errors.New("a user is required"))
	}
	if !strings.ContainsRune(user, '@') {
		return errors.WithStack(errors.New("a fully qualified user with a domain is required"))
	}
	if user[0] == '@' {
		return errors.WithStack(errors.New("take the leading @ off the user and try again"))
	}

	if _, err = keyring.Get(keyringServiceAccessToken, user); err == nil {
		slog.Warn("already logged in", "user", user)
		return nil
	}

	instance, err := ensureInstance(user)
	if err != nil {
		slog.Error("couldn't get user's instance", "user", user, "error", err)
		return err
	}

	client := clientForInstance(instance)
	clientID, clientSecret, err := ensureAppCredentials(user, instance, client)
	if err != nil {
		slog.Error("OAuth2 app setup failed", "user", user, "instance", instance, "error", err)
		return err
	}

	code := promptForOAuthCode(instance, clientID)

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

	slog.Info("login successful", "user", user, "instance", instance)

	return nil
}

// ensureInstance finds a user's instance or retrieves a previously cached instance for them.
func ensureInstance(user string) (string, error) {
	if instance, err := keyring.Get(keyringServiceInstance, user); err == nil {
		return instance, nil
	}

	instance, err := findInstance(user)
	if err != nil {
		slog.Error("WebFinger lookup failed", "user", user, "error", err)
		return "", err
	}

	err = keyring.Set(keyringServiceInstance, user, instance)
	if err != nil {
		slog.Error("couldn't set instance in keychain", "user", user, "instance", instance, "error", err)
		return "", err
	}

	return instance, nil
}

// findInstance does a WebFinger lookup to find the domain of the instance API for a given user.
func findInstance(user string) (string, error) {
	webfingerClient := webfinger.NewClient(nil)
	jrd, err := webfingerClient.Lookup(user, nil)
	if err != nil {
		return "", err
	}

	var href string
	for _, link := range jrd.Links {
		if link.Rel == "self" && link.Type == "application/activity+json" {
			href = link.Href
			break
		}
	}
	if href == "" {
		return "", errors.New("no link with rel=\"self\" and type=\"application/activity+json\"")
	}

	url, err := neturl.Parse(href)
	if err != nil {
		return "", err
	}

	if url.Scheme != "https" || !(url.Port() == "" || url.Port() == "443") || url.Hostname() == "" {
		return "", errors.New("unexpected URL format")
	}

	return url.Hostname(), nil
}

func clientForInstance(instance string) *apiclient.GoToSocialSwaggerDocumentation {
	return apiclient.New(httptransport.New(instance, "", []string{"https"}), strfmt.Default)
}

// ensureAppCredentials retrieves or creates and stores app credentials.
func ensureAppCredentials(user string, instance string, client *apiclient.GoToSocialSwaggerDocumentation) (string, string, error) {
	shouldCreateNewApp := false

	clientID, err := keyring.Get(keyringServiceClientID, user)
	if clientID == "" || errors.Is(err, keyring.ErrNotFound) {
		shouldCreateNewApp = true
	} else if err != nil {
		slog.Error("couldn't get client ID from keychain", "user", user, "error", err)
		return "", "", err
	}

	clientSecret, err := keyring.Get(keyringServiceClientSecret, user)
	if clientSecret == "" || errors.Is(err, keyring.ErrNotFound) {
		shouldCreateNewApp = true
	} else if err != nil {
		slog.Error("couldn't get client secret from keychain", "user", user, "error", err)
		return "", "", err
	}

	if !shouldCreateNewApp {
		return clientID, clientSecret, nil
	}

	app, err := createApp(client)
	if err != nil {
		slog.Error("couldn't create OAuth2 app", "user", user, "instance", instance, "error", err)
		return "", "", err
	}
	clientID = app.ClientID
	clientSecret = app.ClientSecret

	err = keyring.Set(keyringServiceClientID, user, clientID)
	if err != nil {
		slog.Error("couldn't set client ID in keychain", "user", user, "error", err)
		return "", "", err
	}
	err = keyring.Set(keyringServiceClientSecret, user, clientSecret)
	if err != nil {
		slog.Error("couldn't set client secret in keychain", "user", user, "error", err)
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

func promptForOAuthCode(instance string, clientID string) string {
	oauthAuthorizeURL := (&neturl.URL{
		Scheme: "https",
		Host:   instance,
		Path:   "/oauth/authorize",
		RawQuery: neturl.Values{
			"response_type": []string{"code"},
			"client_id":     []string{clientID},
			"redirect_uri":  []string{oauthRedirect},
			"scope":         []string{oauthScopes},
		}.Encode(),
	}).String()
	err := browser.OpenURL(oauthAuthorizeURL)
	if err != nil {
		slog.Warn("couldn't open browser to authorize", "error", err)
		print("Please open this URL in your browser:", oauthAuthorizeURL)
	}

	print("Enter authorization code: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	code := strings.TrimSpace(scanner.Text())

	return code
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
	oauthTokenURL := (&neturl.URL{
		Scheme: "https",
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

func Whoami() error {
	// TODO: look up current user from prefs
	return errors.New("NOT IMPLEMENTED")
}
