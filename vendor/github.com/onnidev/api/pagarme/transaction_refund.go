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

// Refund dsjkfndsfkfjng
func (api *API) Refund(ctx context.Context, id string) error {
	URL, err := api.getURL()
	if err != nil {
		return err
	}
	URL.Path += "/1/transactions/" + id + "/refund"
	parameters := url.Values{}
	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()
	fmt.Println(URL.String())
	req, err := http.NewRequest("POST", URL.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	if err != nil {
		return err
	}
	log.Println(resp.StatusCode)
	return nil
}

// SplitRefund dsjkfndsfkfjng
func (api *API) SplitRefund(ctx context.Context, tid string, request types.PagarMeTransactionRefundRequest) (types.PagarMeTransactionResponse, error) {
	var response types.PagarMeTransactionResponse
	URL, err := api.getURL()
	if err != nil {
		return response, err
	}
	URL.Path += "/1/transactions/" + tid + "/refund"
	parameters := url.Values{}
	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()
	j, err := json.MarshalIndent(request, "", "    ")
	if err != nil {
		return response, err
	}
	log.Println("Request sent to PagarMe")
	log.Println(string(j))
	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(j))
	if err != nil {
		return response, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	if resp.StatusCode >= 400 {
		err := errors.New("bad boy")
		var erroResponse types.PagarMeTransactionError
		err = json.Unmarshal(byt, &erroResponse)
		if err != nil {
			return response, err
		}
		for _, pgerr := range erroResponse.Errors {
			err := errors.New(pgerr.Message)
			return response, err
		}
		return response, err
	}
	log.Println(string(byt))
	err = json.Unmarshal(byt, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
