package recipient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// ListTransactions is commented
func ListTransactions(w http.ResponseWriter, r *http.Request) {
	filter := r.Context().
		Value(middlewares.FinanceQueryReq).(types.FinanceQuery)
	recipientRepo := r.Context().
		Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	recipient, err := recipientRepo.GetByToken(filter.RecipientID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	term := types.ESTransactionTerm{
		Term: types.ESTransactionSplitRules{SplitRules: recipient.RecipientID}}
	rangeI := types.ESTransactionRange{
		Range: types.ESTransactiosDate{
			Date: types.ESTransactionDateCreated{
				Lte: filter.Till.Time().Format("2006-01-02"),
				Gte: filter.From.Time().Format("2006-01-02"),
			}}}
	transactionReq := types.PagarmeTransactionRequest{
		From: 0,
		Size: 10,
		Query: types.ESFinanceQuery{
			Filtered: types.ESTransactionFiltered{
				Filter: types.ESTransactionFilter{
					And: []interface{}{
						term,
						rangeI,
					}},
			},
		},
	}
	j, err := json.MarshalIndent(transactionReq, "", "    ")
	if err != nil {
		log.Println(23444)
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	b, _ := json.Marshal(transactionReq)
	if viper.GetBool("verbose") {
		log.Println("Query sent to PagarMe")
		log.Println(string(j))
	}
	var URL *url.URL
	URL, err = url.Parse("https://api.pagar.me")
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	URL.Path += "/1/search"
	parameters := url.Values{}
	parameters.Add("api_key", "ak_live_iSZM4oGkTcBmVhzGysL9BE2QP6ZAIz")
	parameters.Add("type", "transaction")
	parameters.Add("query", (string(b)))
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
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	var result types.PagarmeTransaction
	err = json.Unmarshal(byt, &result)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
