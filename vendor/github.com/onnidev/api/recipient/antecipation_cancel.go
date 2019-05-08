package recipient

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/pagarme"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// CancelAnteciapation is commented
func CancelAnteciapation(w http.ResponseWriter, r *http.Request) {
	api := pagarme.New(viper.GetString("PAGARME"))
	antecipation, err, code := api.AntecipationsCancel(
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
