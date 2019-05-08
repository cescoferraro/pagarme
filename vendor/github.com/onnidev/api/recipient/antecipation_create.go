package recipient

import (
	"errors"
	"net/http"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// CreateAnteciapation is commented
func CreateAnteciapation(w http.ResponseWriter, r *http.Request) {
	antecipation, ok := r.Context().Value(middlewares.ReadAntecipationPostKey).(types.AntecipationPostRequest)
	if !ok {
		err := errors.New("bug")
		http.Error(w, err.Error(), 400)
		return
	}

	api := pagarme.New(viper.GetString("PAGARME"))
	pagarmeAntecipation, err, code := api.AntecipationsCreate(r.Context(), antecipation)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	render.Status(r, code)
	render.JSON(w, r, pagarmeAntecipation)
}
