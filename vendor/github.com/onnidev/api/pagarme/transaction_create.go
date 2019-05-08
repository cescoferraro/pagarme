package pagarme

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/onnidev/api/types"
)

// TransactionCreate dsjkfndsfkfjng
func (api *API) TransactionCreate(ctx context.Context, request types.PagarMeTransactionRequest) (types.PagarMeTransactionResponse, error) {
	var response types.PagarMeTransactionResponse
	URL, err := api.getURL()
	if err != nil {
		log.Println("erro pegando api key", err.Error())
		return response, err
	}
	URL.Path += "/1/transactions"
	parameters := url.Values{}
	parameters.Add("api_key", api.Key)
	URL.RawQuery = parameters.Encode()
	j, err := json.MarshalIndent(request, "", "    ")
	if err != nil {
		log.Println("erro printing req body", err.Error())
		return response, err
	}
	log.Println(string(j))
	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(j))
	if err != nil {
		log.Println("erro na criando request buy", err.Error())
		return response, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := api.defaultHTTPClient().Do(req)
	if err != nil {
		log.Println("erro na execucao da  request buy", err.Error())
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
			log.Println("erro unmarshalling pagarme error", err.Error())
			return response, err
		}
		log.Println(string(byt))
		for _, pgerr := range erroResponse.Errors {
			err := errors.New(pgerr.Message)
			log.Println("erro normal do pagarme ", err.Error())
			return response, err
		}
		log.Println("erro louco do pg", err.Error())
		return response, errors.New("crcazy pg")
	}
	err = json.Unmarshal(byt, &response)
	if err != nil {
		log.Println("erro unmarshalling successo no pagarme", err.Error())
		return response, err
	}
	log.Println("deu tudo certo")
	return response, nil
}
