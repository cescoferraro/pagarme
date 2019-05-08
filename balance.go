package getnet

import (
	"context"
)

// RecipientBalance dsjkfndsfkfjng
func (api *API) Auth(ctx context.Context, recipientID string) (interface{}, error) {
	// var balance types.PagarMeTransactionBalance
	// URL, err := api.getURL()
	// if err != nil {
	// 	return balance, err
	// }
	// URL.Path += "/1/recipients/" + recipientID + "/balance"
	// parameters := url.Values{}
	// parameters.Add("api_key", api.Key)
	// URL.RawQuery = parameters.Encode()
	// req, err := http.NewRequest("GET", URL.String(), nil)
	// if err != nil {
	// 	return balance, err
	// }
	// if api.Verbose {
	// 	fmt.Println(URL.String())
	// }
	// req.Header.Set("Content-Type", "application/json")
	// resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	// if err != nil {
	// 	return balance, err
	// }
	// defer resp.Body.Close()
	// byt, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return balance, err
	// }
	// err = json.Unmarshal(byt, &balance)
	// if err != nil {
	// 	return balance, err
	// }
	return nil, nil
}
