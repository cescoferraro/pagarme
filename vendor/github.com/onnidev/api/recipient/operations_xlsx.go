package recipient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
)

// Xlsx is commented
func Xlsx(w http.ResponseWriter, r *http.Request) {
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
	URL.Path += "/1/recipients/" + recipient.RecipientID + "/balance/operations.xlsx"
	parameters := url.Values{}
	parameters.Add("start_date", strconv.FormatInt(filter.From.Time().Unix()*1000, 10))
	parameters.Add("end_date", strconv.FormatInt(filter.Till.Time().Unix()*1000, 10))
	parameters.Add("api_key", "ak_live_iSZM4oGkTcBmVhzGysL9BE2QP6ZAIz")
	URL.RawQuery = parameters.Encode()
	fmt.Println(URL.String())
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	io.Copy(w, resp.Body)
	resp.Body.Close()
}
