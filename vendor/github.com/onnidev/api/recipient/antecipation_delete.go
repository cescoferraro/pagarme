package recipient

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// DeleteAnteciapation is commented
func DeleteAnteciapation(w http.ResponseWriter, r *http.Request) {
	api := pagarme.New(viper.GetString("PAGARME"))
	code, err := api.AntecipationsDelete(
		r.Context(),
		chi.URLParam(r, "recipientID"),
		chi.URLParam(r, "bulkID"),
	)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, code)
	render.JSON(w, r, true)
}
