package party

import (
	"net/http"

	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// ListPartiesFiltered Parties sdjknfdskjf
func AppListPartiesFiltered(w http.ResponseWriter, r *http.Request) {
	allParties, err := onni.GetPartiesFromFilter(r.Context())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allParties)
}
