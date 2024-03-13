package own

import (
	"github.com/go-openapi/runtime"

	apiclient "github.com/VyrCossont/slurp/client"
	"github.com/VyrCossont/slurp/models"
)

// Account returns the currently authenticated account.
func Account(client *apiclient.GoToSocialSwaggerDocumentation, auth runtime.ClientAuthInfoWriter) (*models.Account, error) {
	resp, err := client.Accounts.AccountVerify(nil, auth)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

// Instance returns the instance of the currently authenticated account.
func Instance(client *apiclient.GoToSocialSwaggerDocumentation) (*models.InstanceV2, error) {
	resp, err := client.Instance.InstanceGetV2(nil)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}
