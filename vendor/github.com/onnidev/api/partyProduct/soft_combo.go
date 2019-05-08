package partyProduct

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// SoftCombo carinho
func SoftCombo(w http.ResponseWriter, r *http.Request) {
	productsCollection := r.Context().Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	id := chi.URLParam(r, "partyId")
	allUser, err := productsCollection.GetByPartyID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	result := []types.SoftComboPartyProduct{}
	for _, partyP := range allUser {
		if partyP.Status == "ACTIVE" {
			if len(partyP.Batches) == 0 {
				result = append(result, partyP.SoftCombo())
			}
		}
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
