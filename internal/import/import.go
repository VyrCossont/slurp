package _import

import (
	"errors"

	"github.com/go-openapi/runtime"

	apiclient "github.com/VyrCossont/slurp/client"
	"github.com/VyrCossont/slurp/client/search"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

var NoResults = errors.New("no search results")

func Account(client *apiclient.GoToSocialSwaggerDocumentation, auth runtime.ClientAuthInfoWriter, accountOrURL string) (*models.Account, error) {
	results, err := search1(client, auth, accountOrURL, "accounts")
	if err != nil {
		return nil, err
	}
	if len(results.Accounts) == 0 {
		return nil, NoResults
	}
	return results.Accounts[0], nil
}

func Status(client *apiclient.GoToSocialSwaggerDocumentation, auth runtime.ClientAuthInfoWriter, url string) (*models.Status, error) {
	results, err := search1(client, auth, url, "statuses")
	if err != nil {
		return nil, err
	}
	if len(results.Statuses) == 0 {
		return nil, NoResults
	}
	return results.Statuses[0], nil
}

func search1(client *apiclient.GoToSocialSwaggerDocumentation, auth runtime.ClientAuthInfoWriter, q string, searchType string) (*models.SearchResult, error) {
	resp, err := client.Search.SearchGet(&search.SearchGetParams{
		APIVersion: "v2",
		Limit:      util.Ptr(int64(1)),
		Q:          q,
		Resolve:    util.Ptr(true),
		Type:       util.Ptr(searchType),
	}, auth)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}
