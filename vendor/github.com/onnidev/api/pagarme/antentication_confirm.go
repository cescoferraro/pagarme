package pagarme

import (
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

// AntecipationsConfirm dsjkfndsfkfjng
func (api *API) AntecipationsConfirm(ctx context.Context, recipientID, bulkID string) (
	types.PagarMeAntecipationResponse,
	error, int) {
	var response types.PagarMeAntecipationResponse
	URL, err := api.getURL()
	if err != nil {
		return response, err, 400
	}
	URL.Path += "/1/recipients/" + recipientID + "/bulk_anticipations/" + bulkID + "/confirm"
	parameters := url.Values{}

	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()
	req, err := http.NewRequest("POST", URL.String(), nil)
	if err != nil {
		return response, err, 400
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	if err != nil {
		return response, err, 400
	}
	log.Println(resp.StatusCode)
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err, resp.StatusCode
	}
	log.Println(string(byt))
	if resp.StatusCode >= 400 {
		err := errors.New(string(byt))
		return response, err, resp.StatusCode
	}
	err = json.Unmarshal(byt, &response)
	if err != nil {
		return response, err, resp.StatusCode
	}
	return response, nil, 200
}
