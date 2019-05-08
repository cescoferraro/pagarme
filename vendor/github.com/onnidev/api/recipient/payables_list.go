package recipient

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Payables  is commented
func Payables(w http.ResponseWriter, r *http.Request) {
	filter := r.Context().
		Value(middlewares.FinanceQueryReq).(types.FinanceQuery)
	recipientRepo, ok := r.Context().
		Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	recipient, err := recipientRepo.GetByToken(filter.RecipientID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	api := pagarme.New(viper.GetString("PAGARME"))
	payables, err := api.Payables(
		r.Context(),
		recipient.RecipientID,
		strconv.FormatInt(filter.From.Time().Unix()*1000, 10))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, payables)
}
