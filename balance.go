package pagarme

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/net/context/ctxhttp"
)

// RecipientBalance dsjkfndsfkfjng
func (api *API) RecipientBalance(ctx context.Context, recipientID string) (RecipientBalance, error) {
	var balance RecipientBalance
	URL, err := api.getURL()
	if err != nil {
		return balance, err
	}
	URL.Path += "/1/recipients/" + recipientID + "/balance"
	parameters := url.Values{}
	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return balance, err
	}
	if api.Verbose {
		fmt.Println(URL.String())
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	if err != nil {
		return balance, err
	}
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return balance, err
	}
	err = json.Unmarshal(byt, &balance)
	if err != nil {
		return balance, err
	}
	return balance, nil
}
