package recipient

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// PagarMeRecipientTweaks is commented
func PagarMeRecipientTweaks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	post, ok := r.Context().Value(middlewares.RecipientTweaksPatchKey).(types.RecipientTweaksPatch)
	if !ok {
		err := errors.New("assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	recipient, err := onni.GetRecipientByToken(ctx, chi.URLParam(r, "recipientID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	api := pagarme.New(viper.GetString("PAGARME"))
	pagarMeRecipient, code, err := api.Recipient(r.Context(), recipient.RecipientID)
	if err != nil {
		shared.MakeONNiError(w, r, code, err)
		return
	}
	tweakedRecipient, code, err := api.RecipientTweaks(r.Context(), pagarMeRecipient, post)
	if err != nil {
		shared.MakeONNiError(w, r, code, err)
		return
	}
	j, _ := json.MarshalIndent(tweakedRecipient, "", "    ")
	log.Println(string(j))
	render.Status(r, http.StatusOK)
	render.JSON(w, r, tweakedRecipient)
}
