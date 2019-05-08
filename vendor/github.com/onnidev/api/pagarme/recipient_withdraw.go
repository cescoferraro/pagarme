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

// RecipientWithDraw dsjkfndsfkfjng
func (api *API) RecipientWithDraw(ctx context.Context, withdraw types.RecipientWithDraw) (int, error) {
	URL, err := api.getURL()
	if err != nil {
		return 400, err
	}
	URL.Path += "/1/transfers"
	parameters := url.Values{}

	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()
	fmt.Println(URL.String())
	j, _ := json.MarshalIndent(withdraw, "", "    ")
	log.Println("Request sent to PagarMe")
	log.Println(string(j))
	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(j))
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
