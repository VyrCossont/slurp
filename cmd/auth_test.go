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

package cmd_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	neturl "net/url"
	"regexp"
	"strings"
	"testing"

	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/util"
	gokeyring "github.com/zalando/go-keyring"
	"golang.org/x/net/publicsuffix"
)

// gtsTestrigAuthorizer logs into a GTS testrig and returns an OOB auth code.
func gtsTestrigAuthorizer(oauthAuthorizeURL string) (code string, err error) {
	// Create a client with a cookie jar.
	var client = &http.Client{}
	*client = *util.HttpClient
	client.Jar, err = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return
	}

	// Load the authorization URL.
	resp, err := client.Get(oauthAuthorizeURL)
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// Log in.
	resp, err = client.PostForm("http://localhost:8080/auth/sign_in", neturl.Values{
		"username": []string{"zork@example.org"},
		"password": []string{"password"},
	})
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// Agree to the authorization prompt.
	resp, err = client.PostForm("http://localhost:8080/oauth/authorize", neturl.Values{})
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// Get the authorization code.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	authCodePattern, err := regexp.Compile(`<code>(\w+)</code>`)
	if err != nil {
		return
	}
	match := authCodePattern.FindStringSubmatch(string(body))
	if len(match) < 2 {
		err = errors.New("couldn't find OOB auth code")
		return
	}

	code = match[1]
	return
}

// Checks that you have a GTS testrig running.
func TestGtsTestrigPresence(t *testing.T) {
	resp, err := util.HttpClient.Get("http://localhost:8080/")
	if err != nil {
		t.Error("failed to get landing page")
		t.FailNow()
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "GoToSocial Testrig Instance") {
		t.Error("whatever's running on localhost:8080 isn't a GTS testrig")
	}
}

// Test that we can log into it.
func TestGtsTestrigLogin(t *testing.T) {
	gokeyring.MockInit()
	if err := auth.Login(
		"the_mighty_zork@localhost:8080",
		true,
		util.SystemKeyring,
		gtsTestrigAuthorizer,
	); err != nil {
		t.Errorf("error logging GTS testrig: %+v", err)
	}
}
