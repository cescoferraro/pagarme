package party

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// ReadPartySoft TODO: NEEDS COMMENT INFO
func ReadPartySoft(w http.ResponseWriter, r *http.Request) {
	partiesCollection := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	id := chi.URLParam(r, "partyId")
	party, err := partiesCollection.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	softparty, err := onni.SoftParty(r.Context(), party)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, softparty)
}
