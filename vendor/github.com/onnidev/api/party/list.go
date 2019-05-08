package party

import (
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// ListParties sdjknfdskjf
func ListParties(w http.ResponseWriter, r *http.Request) {
	partiesCollection := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	// lista festas que comecam em 1 hora
	allUser, err := partiesCollection.List()
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allUser)
}
