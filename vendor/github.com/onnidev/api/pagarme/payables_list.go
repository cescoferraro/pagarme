package pagarme

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/bradfitz/slice"
	"github.com/onnidev/api/types"
	"golang.org/x/net/context/ctxhttp"
)

// Payables dsjkfndsfkfjng
func (api *API) Payables(ctx context.Context, recipientID, filter string) ([]types.PagarMePayable, error) {
	var balance []types.PagarMePayable
	URL, err := api.getURL()
	if err != nil {
		return balance, err
	}
	URL.Path += "/1/payables"
	parameters := url.Values{}
	parameters.Add("api_key", api.Key)
	parameters.Add("recipient_id", recipientID)
	parameters.Add("status", "waiting_funds")
	parameters.Add("payment_date", filter)
	parameters.Add("count", "1000300")
	parameters.Add("page", "1")
	URL.RawQuery = parameters.Encode()
	fmt.Println(URL.String())
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return balance, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	if err != nil {
		return balance, err
	}
	log.Println(resp.StatusCode)
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return balance, err
	}
	err = json.Unmarshal(byt, &balance)
	if err != nil {
		return balance, err
	}
	slice.Sort(balance[:], func(i, j int) bool {
		return balance[i].DateCreated.Before(balance[j].DateCreated)
	})
	return balance, nil
}
