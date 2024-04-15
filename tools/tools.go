//go:build tools

// tools exists to pull in command-line tools that we need to go get,
// and is behind a build tag that is otherwise unused and thus only visible
// to dependency management commands. See https://stackoverflow.com/a/54028731.
package tools

import (
	// Provides swagger command used to update GtS client library
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
)
