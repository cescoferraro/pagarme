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

// Antecipations dsjkfndsfkfjng
func (api *API) AntecipationsLimit(ctx context.Context, recipientID, date string) (types.PagarMeLimits, error, int) {
	var balance types.PagarMeLimits
	URL, err := api.getURL()
	if err != nil {
		return balance, err, 400
	}
	URL.Path += "/1/recipients/" + recipientID + "/bulk_anticipations/limits"
	parameters := url.Values{}
	parameters.Add("api_key", api.Key)
	parameters.Add("timeframe", "start")
	parameters.Add("payment_date", date)
	URL.RawQuery = parameters.Encode()
	fmt.Println(URL.String())
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return balance, err, 400
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	if err != nil {
		return balance, err, 400
	}
	log.Println(resp.StatusCode)
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return balance, err, 400
	}
	log.Println(string(byt))
	if resp.StatusCode >= 400 {
		err := errors.New(string(byt))
		return balance, err, resp.StatusCode
	}
	err = json.Unmarshal(byt, &balance)
	if err != nil {
		return balance, err, 400
	}
	return balance, nil, resp.StatusCode
}
