package recipient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bradfitz/slice"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Operations is commented
func Operations(w http.ResponseWriter, r *http.Request) {
	filter := r.Context().
		Value(middlewares.FinanceQueryReq).(types.FinanceQuery)
	recipientRepo := r.Context().
		Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	recipient, err := recipientRepo.GetByToken(filter.RecipientID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	var URL *url.URL
	URL, err = url.Parse("https://api.pagar.me")
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	URL.Path += "/1/recipients/" + recipient.RecipientID + "/balance/operations"
	parameters := url.Values{}
	parameters.Add("start_date", strconv.FormatInt(filter.From.Time().Unix()*1000, 10))
	parameters.Add("end_date", strconv.FormatInt(filter.Till.Time().Unix()*1000, 10))
	parameters.Add("api_key", "ak_live_iSZM4oGkTcBmVhzGysL9BE2QP6ZAIz")
	parameters.Add("status", "available")
	parameters.Add("count", "300")
	URL.RawQuery = parameters.Encode()
	fmt.Println(URL.String())
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(987)
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	log.Println(resp.StatusCode)
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	var balance []types.PagameOperation
	err = json.Unmarshal(byt, &balance)
	if err != nil {
		log.Println(err.Error())
		return
	}
	slice.Sort(balance[:], func(i, j int) bool {
		return balance[i].DateCreated.Before(balance[j].DateCreated)
	})
	render.Status(r, http.StatusOK)
	render.JSON(w, r, balance)
}
