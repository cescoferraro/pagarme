package pagarme

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/onnidev/api/types"
	"golang.org/x/net/context/ctxhttp"
)

// RecipientEditBankingInfo dsjkfndsfkfjng
func (api *API) RecipientEditBankingInfo(ctx context.Context, post types.Recipient) (
	types.PagarMeRecipient, int, error) {
	var response types.PagarMeRecipient
	URL, err := api.getURL()
	if err != nil {
		return response, 400, err
	}
	URL.Path += "/1/recipients/" + post.RecipientID
	parameters := url.Values{}

	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()
	fmt.Println(URL.String())
	slimAntecipation := types.PagarMeRecipient{
		BankAccount: types.PagarMeBankAccount{
			BankCode:       post.BankingInfo.BankCode,
			Agencia:        post.BankingInfo.BankBranch,
			Conta:          post.BankingInfo.BankAccount,
			ContaDV:        post.BankingInfo.BankAccountVC,
			LegalName:      post.BankingInfo.BankAccountName,
			DocumentNumber: post.BankingInfo.DocumentNumber,
		},
	}

	if *post.BankingInfo.BankBranchVC != "" {
		slimAntecipation.BankAccount.AgenciaDV = post.BankingInfo.BankBranchVC
	}
	j, _ := json.MarshalIndent(slimAntecipation, "", "    ")
	log.Println("Request sent to PagarMe")
	log.Println(string(j))
	req, err := http.NewRequest("PUT", URL.String(), bytes.NewBuffer(j))
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
		return response, resp.StatusCode, err
	}
	log.Println(string(byt))
	if resp.StatusCode >= 400 {
		err := errors.New(string(byt))
		return response, resp.StatusCode, err
	}
	err = json.Unmarshal(byt, &response)
	if err != nil {
		return response, resp.StatusCode, err
	}
	return response, 200, nil
}
