package recipient

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// WithDraws skdjfn
func WithDraws(w http.ResponseWriter, r *http.Request) {
	recipientRepo := r.Context().Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	recipient, err := recipientRepo.GetByToken(chi.URLParam(r, "recipientID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	api := pagarme.New(viper.GetString("PAGARME"))
	withdraws, code, err := api.Transfers(r.Context(), recipient.RecipientID)
	if err != nil {
		shared.MakeONNiError(w, r, code, err)
		return
	}
	render.Status(r, code)
	render.JSON(w, r, withdraws)
}
