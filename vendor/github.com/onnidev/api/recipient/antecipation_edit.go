package recipient

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// EditAnteciapation is commented
func EditAnteciapation(w http.ResponseWriter, r *http.Request) {
	log.Println("kerpen")
	antecipation, ok := r.Context().
		Value(middlewares.ReadAntecipationPostKey).(types.AntecipationPostRequest)
	if !ok {
		err := errors.New("bug")
		http.Error(w, err.Error(), 400)
		return
	}
	api := pagarme.New(viper.GetString("PAGARME"))
	pagarmeAntecipation, err, code := api.
		AntecipationsEdit(
			r.Context(),
			chi.URLParam(r, "bulkID"),
			antecipation)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	render.Status(r, code)
	render.JSON(w, r, pagarmeAntecipation)
}
