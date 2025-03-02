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

package util

import (
	"fmt"
	"net/http"
)

// userAgentRoundTripper adds a User-Agent header on top of the default transport.
type userAgentRoundTripper struct{}

func (u *userAgentRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("User-Agent", userAgent)
	return http.DefaultTransport.RoundTrip(request)
}

var Website = "https://catgirl.codes/slurp"

var version = "0.0.0"

var userAgent string

var HttpClient *http.Client

func init() {
	userAgent = fmt.Sprintf("slurp/%s (+%s)", version, Website)
	HttpClient = &http.Client{
		Transport: &userAgentRoundTripper{},
	}
}
