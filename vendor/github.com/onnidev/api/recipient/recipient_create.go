package recipient

import (
	"errors"
	"log"
	"net/http"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Create is commented
func Create(w http.ResponseWriter, r *http.Request) {
	recipient, ok := r.Context().Value(middlewares.ReadRecipientPostKey).(types.RecipientPost)
	if !ok {
		err := errors.New("bug")
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println(recipient.ClubID)
	api := pagarme.New(viper.GetString("PAGARME"))
	pagarmeRecipient, status, err := api.RecipientCreate(r.Context(), recipient)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	onniRecipient, err := onni.CreateONNiRecipient(r.Context(), recipient, pagarmeRecipient)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	render.Status(r, status)
	render.JSON(w, r, onniRecipient)
}
