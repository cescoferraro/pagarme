package recipient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/bradfitz/slice"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// PayablesTimeline is commented
func PayablesTimeline(w http.ResponseWriter, r *http.Request) {
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
	URL.Path += "/1/payables/days"
	parameters := url.Values{}
	parameters.Add("start_date", strconv.FormatInt(filter.From.Time().Unix()*1000, 10))
	parameters.Add("end_date", strconv.FormatInt(filter.Till.Time().Unix()*1000, 10))
	parameters.Add("recipient_id", recipient.RecipientID)
	parameters.Add("api_key", "ak_live_iSZM4oGkTcBmVhzGysL9BE2QP6ZAIz")
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
	var balance []types.PagarmePayablesTimeline
	err = json.Unmarshal(byt, &balance)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	var existence []time.Time
	for _, Payable := range balance {
		existence = append(existence, Payable.Date)
	}
	if len(balance) != 0 {
		space := int(filter.Till.Time().Sub(filter.From.Time()).Hours() / 24)
		bal := make([]types.PagarmePayablesTimeline, space)
		for index := range bal {
			stretch := time.Duration((index)*24) * time.Hour
			gap := filter.From.Time().Add(stretch)
			i, contain := (ContainDate(existence, gap))
			var timeline types.PagarmePayablesTimeline
			if contain {
				timeline = balance[i]
				timeline.Real = true
			} else {
				timeline = types.PagarmePayablesTimeline{Date: gap}
			}
			bal[index] = timeline
		}
		balance = bal
	}
	slice.Sort(balance[:], func(i, j int) bool {
		return balance[i].Date.Before(balance[j].Date)
	})
	render.Status(r, http.StatusOK)
	render.JSON(w, r, balance)
}

func ContainDate(all []time.Time, the time.Time) (int, bool) {
	for i, this := range all {
		y1, m1, d1 := this.Date()
		y2, m2, d2 := the.Date()
		if y1 == y2 && m1 == m2 && d1 == d2 {
			return i, true
		}
	}
	return 0, false
}
