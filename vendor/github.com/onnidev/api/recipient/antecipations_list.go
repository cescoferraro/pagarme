package recipient

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Anteciapations is commented
func Anteciapations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	recipient, err := onni.GetRecipientByToken(ctx, chi.URLParam(r, "recipientID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	api := pagarme.New(viper.GetString("PAGARME"))
	payables, code, err := api.Antecipations(r.Context(), recipient.RecipientID)
	if err != nil {
		shared.MakeONNiError(w, r, code, err)

		return
	}
	render.Status(r, code)
	render.JSON(w, r, payables)
}
