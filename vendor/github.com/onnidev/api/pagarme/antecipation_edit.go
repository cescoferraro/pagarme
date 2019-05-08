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

// AntecipationsEdit dsjkfndsfkfjng
func (api *API) AntecipationsEdit(ctx context.Context, id string,
	antecipation types.AntecipationPostRequest) (
	types.PagarMeAntecipationResponse, error, int) {
	var response types.PagarMeAntecipationResponse
	URL, err := api.getURL()
	if err != nil {
		return response, err, 400
	}
	URL.Path += "/1/recipients/" + antecipation.RecipientID + "/bulk_anticipations/" + id
	parameters := url.Values{}

	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()
	fmt.Println(URL.String())
	slimAntecipation := types.PagarMeAntecipationPostRequest{
		APIKey:          api.Key,
		Build:           antecipation.Build,
		Timeframe:       antecipation.Timeframe,
		PaymentDay:      antecipation.PaymentDay,
		RequestedAmount: int32(antecipation.RequestedAmount),
	}
	j, _ := json.MarshalIndent(slimAntecipation, "", "    ")
	log.Println("Request sent to PagarMe")
	log.Println(string(j))
	req, err := http.NewRequest("PUT", URL.String(), bytes.NewBuffer(j))
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
