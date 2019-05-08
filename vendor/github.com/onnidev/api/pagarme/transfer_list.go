package pagarme

import (
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

// Transfers dsjkfndsfkfjng
func (api *API) Transfers(ctx context.Context, recipientID string) ([]types.PagarMeTransfer, int, error) {
	var balance []types.PagarMeTransfer
	URL, err := api.getURL()
	if err != nil {
		return balance, 400, err
	}
	URL.Path += "/1/transfers"
	parameters := url.Values{}
	parameters.Add("recipient_id", recipientID)
	parameters.Add("api_key", api.Key)
	parameters.Add("count", "1000300")
	URL.RawQuery = parameters.Encode()
	fmt.Println(URL.String())
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return balance, 400, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	if err != nil {
		return balance, 400, err
	}
	log.Println(resp.StatusCode)
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return balance, 400, err
	}
	log.Println(string(byt))
	if resp.StatusCode >= 400 {
		err := errors.New(string(byt))
		return balance, resp.StatusCode, err
	}
	err = json.Unmarshal(byt, &balance)
	if err != nil {
		return balance, 400, err
	}
	return balance, resp.StatusCode, nil
}
