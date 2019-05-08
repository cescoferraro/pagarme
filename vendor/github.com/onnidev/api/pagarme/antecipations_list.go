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

	"github.com/bradfitz/slice"
	"github.com/onnidev/api/types"
	"golang.org/x/net/context/ctxhttp"
)

// Antecipations dsjkfndsfkfjng
func (api *API) Antecipations(ctx context.Context, recipientID string) ([]types.PagarMeAntecipation, int, error) {
	var balance []types.PagarMeAntecipation
	URL, err := api.getURL()
	if err != nil {
		return balance, 400, err
	}
	URL.Path += "/1/recipients/" + recipientID + "/bulk_anticipations/"
	parameters := url.Values{}
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

	slice.Sort(balance[:], func(i, j int) bool {
		return balance[i].DateCreated.After(balance[j].DateCreated)
	})
	return balance, resp.StatusCode, nil
}
