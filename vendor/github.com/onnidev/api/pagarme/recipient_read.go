package pagarme

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/onnidev/api/types"
	"golang.org/x/net/context/ctxhttp"
)

// Recipient dsjkfndsfkfjng
func (api *API) Recipient(ctx context.Context, recipientID string) (types.PagarMeRecipient, int, error) {
	var recipient types.PagarMeRecipient
	URL, err := api.getURL()
	if err != nil {
		return recipient, 400, err
	}
	URL.Path += "/1/recipients/" + recipientID
	parameters := url.Values{}
	parameters.Add("api_key", api.Key)
	parameters.Add("count", "1000300")
	URL.RawQuery = parameters.Encode()
	fmt.Println(URL.String())
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return recipient, 400, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	if err != nil {
		return recipient, 400, err
	}
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recipient, 400, err
	}
	if resp.StatusCode >= 400 {
		err := errors.New(string(byt))
		return recipient, resp.StatusCode, err
	}
	err = json.Unmarshal(byt, &recipient)
	if err != nil {
		return recipient, 400, err
	}
	return recipient, resp.StatusCode, nil
}
