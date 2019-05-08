package recipient

import (
	"net/http"
	"strconv"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// AnteciapationsLimit is commented
func AnteciapationsLimit(w http.ResponseWriter, r *http.Request) {
	filter := r.Context().
		Value(middlewares.FinanceQueryReq).(types.FinanceQuery)
	ctx := r.Context()
	recipient, err := onni.GetRecipientByToken(ctx, filter.RecipientID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}

	api := pagarme.New(viper.GetString("PAGARME"))
	antecipations, status, err := api.Antecipations(r.Context(), recipient.RecipientID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	for _, antecipation := range antecipations {
		if antecipation.Status == "building" {
			_, _ = api.AntecipationsDelete(r.Context(), recipient.RecipientID, antecipation.ID)
		}
	}

	payables, err, status := api.AntecipationsLimit(
		r.Context(),
		recipient.RecipientID,
		strconv.Itoa(int(filter.From.Time().Unix())*1000))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, status)
	render.JSON(w, r, payables)
}
