package followers

import (
	"strings"

	"github.com/VyrCossont/slurp/client/accounts"
	"github.com/VyrCossont/slurp/internal/api"
	"github.com/VyrCossont/slurp/internal/auth"
	"github.com/VyrCossont/slurp/internal/own"
	"github.com/VyrCossont/slurp/internal/util"
	"github.com/VyrCossont/slurp/models"
)

func Export(authClient *auth.Client, file string) error {
	ownAccount, err := own.Account(authClient)
	if err != nil {
		return err
	}
	pagedRequester := &accountFollowersPagedRequester{accountID: ownAccount.ID}

	ownDomain, err := own.Domain(authClient)
	if err != nil {
		return err
	}

	followedAccounts, err := api.ReadAllPaged(authClient, pagedRequester, nil, nil)
	if err != nil {
		return err
	}

	followers := make([]*followersListEntry, 0, len(followedAccounts))
	for _, account := range followedAccounts {
		followers = append(followers, newFollowersListEntry(ownDomain, account))
	}

	csvRows := make([][]string, 0, 1+len(followers))
	csvRows = append(csvRows, csvHeader)
	for _, follower := range followers {
		csvRows = append(csvRows, follower.csvFields())
	}

	return util.WriteCSV(file, csvRows)
}

type accountFollowersPagedRequester struct {
	accountID     string
	forwardPaging bool
}

func (pagedRequester *accountFollowersPagedRequester) Request(authClient *auth.Client, maxID *string, minID *string) (*accountFollowersPagedResponse, error) {
	resp, err := authClient.Client.Accounts.AccountFollowers(&accounts.AccountFollowersParams{
		ID:    pagedRequester.accountID,
		MaxID: maxID,
		MinID: minID,
	}, authClient.Auth)
	if err != nil {
		return nil, err
	}
	return &accountFollowersPagedResponse{resp}, nil
}

func (pagedRequester *accountFollowersPagedRequester) ForwardPaging() bool {
	return pagedRequester.forwardPaging
}

type accountFollowersPagedResponse struct {
	resp *accounts.AccountFollowersOK
}

func (pagedResponse *accountFollowersPagedResponse) Link() string {
	return pagedResponse.resp.Link
}

func (pagedResponse *accountFollowersPagedResponse) Elements() []*models.Account {
	return pagedResponse.resp.GetPayload()
}

var csvHeader = []string{
	"Account address",
}

type followersListEntry struct {
	accountAddress string
}

func newFollowersListEntry(ownDomain string, account *models.Account) *followersListEntry {
	e := &followersListEntry{
		accountAddress: account.Acct,
	}
	if !strings.ContainsRune(e.accountAddress, '@') {
		e.accountAddress += "@" + ownDomain
	}
	return e
}

func (e *followersListEntry) csvFields() []string {
	return []string{
		e.accountAddress,
	}
}
