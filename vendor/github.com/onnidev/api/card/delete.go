package card

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Delete is commented
// swagger:route DELETE /cards/{id} app deleteCard
//
// Deletes a card document on MongoDB of a given id.
//
// This will enables to delete a card with such id.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Security:
//       X-AUTH-APPLICATION-TOKEN:
//       X-AUTH-ID:
//       X-CLIENT-ID:
//
//
//     Responses:
//       200: cardsList
//       404: error
func Delete(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().Value(middlewares.CardsRepoKey).(interfaces.CardsRepo)
	customer := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	id := chi.URLParam(r, "id")
	cards, err := repo.GetByCustomerID(customer.ID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	for _, card := range cards {
		if card.ID.Hex() == id {
			err := onni.DeleteCard(r.Context(), customer, id)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			render.Status(r, http.StatusOK)
			render.JSON(w, r, id)
			return
		}
	}
	http.Error(w, err.Error(), http.StatusForbidden)
}
