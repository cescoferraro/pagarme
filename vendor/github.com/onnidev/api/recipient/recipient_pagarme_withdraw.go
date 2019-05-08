package recipient

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// WithDraw skdjfn
func WithDraw(w http.ResponseWriter, r *http.Request) {
	recipientRepo := r.Context().Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	recipientPatch := r.Context().Value(middlewares.ReadRecipientWithDrawKey).(types.RecipientWithDraw)
	_, err := recipientRepo.GetByToken(chi.URLParam(r, "recipientID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	api := pagarme.New(viper.GetString("PAGARME"))
	code, err := api.RecipientWithDraw(r.Context(), recipientPatch)
	if err != nil {
		shared.MakeONNiError(w, r, code, err)
		return
	}
	render.Status(r, code)
	render.JSON(w, r, true)
}
