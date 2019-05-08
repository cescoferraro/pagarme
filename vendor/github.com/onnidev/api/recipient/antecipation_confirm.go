package recipient

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/pagarme"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// ConfirmAnteciapation is commented
func ConfirmAnteciapation(w http.ResponseWriter, r *http.Request) {
	api := pagarme.New(viper.GetString("PAGARME"))
	antecipation, err, code := api.AntecipationsConfirm(
		r.Context(),
		chi.URLParam(r, "recipientID"),
		chi.URLParam(r, "bulkID"),
	)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	render.Status(r, code)
	render.JSON(w, r, antecipation)
}
