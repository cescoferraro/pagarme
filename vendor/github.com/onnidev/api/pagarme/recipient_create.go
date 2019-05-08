package pagarme

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/onnidev/api/types"
	"golang.org/x/net/context/ctxhttp"
)

// RecipientCreate dsjkfndsfkfjng
func (api *API) RecipientCreate(ctx context.Context, recipient types.RecipientPost) (
	types.PagarMeRecipient, int, error) {
	var response types.PagarMeRecipient
	URL, err := api.getURL()
	if err != nil {
		return response, 400, err
	}
	URL.Path += "/1/recipients"
	parameters := url.Values{}

	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()

	name := recipient.BankAccountName
	if len(name) > 12 {
		name = name[:11]
	}
	card := types.PagarMeRecipientPostRequest{
		AutomaticAntecipationsEnabled: "false",
		AntecipatableVolumePercentage: "100",
		TransferDay:                   "3",
		TransferEnabled:               "true",
		TransferInterval:              "weekly",
		BankAccount: types.PagarMeRecipientPostBankAccount{
			BankCode:       recipient.BankCode,
			Branch:         recipient.BankBranch,
			Account:        recipient.BankAccount,
			AccountVC:      recipient.BankAccountVC,
			Type:           recipient.BankAccountType,
			DocumentNumber: recipient.DocumentNumber,
			LegalName:      name,
		},
	}
	if recipient.BankAccountVC == "" {
		card.BankAccount.AccountVC = "0"
	}
	if recipient.BankBranchVC != "" {
		card.BankAccount.BranchVC = &recipient.BankBranchVC
	}
	j, _ := json.MarshalIndent(card, "", "    ")
	log.Println("Request sent to PagarMe")
	log.Println(string(j))
	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(j))
	if err != nil {
		return response, 400, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	if err != nil {
		return response, 400, err
	}
	log.Println(resp.StatusCode)
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, 400, err
	}
	log.Println(string(byt))
	if resp.StatusCode >= 400 {
		err := errors.New(string(byt))
		return response, resp.StatusCode, err
	}
	err = json.Unmarshal(byt, &response)
	if err != nil {
		return response, 400, err
	}
	log.Println("antes ded acaber pagar.me")
	return response, resp.StatusCode, nil
}
