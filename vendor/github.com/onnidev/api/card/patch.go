package card

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Patch is commented
// swagger:route PATCH /cards/{id} app patchCard
//
// Patches a card document on MongoDB of a given id.
//
// This will enables to update a card with such id.
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
//       200: dbCard
//       404: error
func Patch(w http.ResponseWriter, r *http.Request) {
	cardsRepo := r.Context().Value(middlewares.CardsRepoKey).(interfaces.CardsRepo)
	cardPatch := r.Context().Value(middlewares.CardUpdateKey).(types.CardPatch)
	customer := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	log.Println(cardPatch)
	id := chi.URLParam(r, "id")
	if cardPatch.Default {
		err := makeCustomerCardsDull(cardsRepo, customer)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
	}
	result := types.Card{}
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"defaultCard": cardPatch.Default,
			}},
		ReturnNew: true,
	}
	_, err := cardsRepo.
		Collection.
		Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &result)
	if err != nil {
		log.Println(err.Error())
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
