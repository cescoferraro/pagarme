package party

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// ReadParty TODO: NEEDS COMMENT INFO
func ReadParty(w http.ResponseWriter, r *http.Request) {
	partiesCollection := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	id := chi.URLParam(r, "partyId")
	allUser, err := partiesCollection.GetByID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allUser)

}
