package card

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Read is commented
// swagger:route GET /cards/{id} app getCard
//
// Reads a card document on MongoDB of a given id.
//
// This will enables to get all card information
// from such id.
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
func Read(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().Value(middlewares.CardsRepoKey).(interfaces.CardsRepo)
	id := chi.URLParam(r, "id")
	card, err := repo.GetByID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, 22)
		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(card, "", "    ")
		log.Println("Card you requested from MongoDB")
		log.Println(string(j))
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, card)
}
