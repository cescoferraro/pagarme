package party

import (
	"net/http"

	"github.com/bradfitz/slice"
	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// ReadClubPartiesSite is the shit
func ReadClubPartiesSite(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clubId")
	partiesCollection := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	parties, err := partiesCollection.GetByClubIDSite(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	if id == "5b31b3ab08d1250001d0d000" {
		// coolture sobe a serrra
		party1, err := partiesCollection.GetByID("5af8c25acc922d4af19fb9c5")
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		if party1.Status != "CLOSED" {
			parties = append(parties, party1)
		}
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, SortPartiesSite(parties))

}

// SortPartiesSite TODO: NEEDS COMMENT INFO
func SortPartiesSite(parties []types.Party) []types.Party {

	slice.Sort(parties[:], func(i, j int) bool {
		return (parties[i].StartDate.Time().Before(parties[j].StartDate.Time()))
	})
	return parties
}
