package recipient

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// PagarMeRecipient is commented
func PagarMeRecipient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	recipient, err := onni.GetRecipientByToken(ctx, chi.URLParam(r, "recipientID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	api := pagarme.New(viper.GetString("PAGARME"))
	pagarmerecipient, code, err := api.Recipient(r.Context(), recipient.RecipientID)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	j, _ := json.MarshalIndent(pagarmerecipient, "", "    ")
	log.Println(string(j))
	render.Status(r, http.StatusOK)
	render.JSON(w, r, pagarmerecipient)
}
