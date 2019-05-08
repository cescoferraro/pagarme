package pagarme

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/context/ctxhttp"
)

// AntecipationsEdit dsjkfndsfkfjng
func (api *API) AntecipationsDelete(ctx context.Context, recipientID, bulkID string) (
	int, error) {
	URL, err := api.getURL()
	if err != nil {
		return 400, err
	}
	URL.Path += "/1/recipients/" + recipientID + "/bulk_anticipations/" + bulkID
	parameters := url.Values{}

	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()
	req, err := http.NewRequest("DELETE", URL.String(), nil)
	if err != nil {
		return 400, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	if err != nil {
		return 400, err
	}
	log.Println(resp.StatusCode)
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}
	log.Println(string(byt))
	if resp.StatusCode >= 400 {
		err := errors.New(string(byt))
		return resp.StatusCode, err
	}
	return 200, nil
}
