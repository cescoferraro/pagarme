package pagarme

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/context/ctxhttp"

	"github.com/onnidev/api/types"
)

// TransactionRead dsjkfndsfkfjng
func (api *API) TransactionRead(ctx context.Context, tid string) (types.PagarMeTransactionResponse, error) {
	var response types.PagarMeTransactionResponse
	URL, err := api.getURL()
	if err != nil {
		return response, err
	}
	URL.Path += "/1/transactions/" + tid
	parameters := url.Values{}
	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()
	log.Println("Request sent to PagarMe")
	req, err := http.NewRequest("GET", URL.String(), nil)
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
