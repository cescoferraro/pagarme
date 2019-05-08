package card

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// ListCards is dskjfn
// swagger:route GET /cards app listCards
//
// Lists cards of a given ser customer.
//
// This will show all available cards by default.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//
//     Security:
//       X-AUTH-APPLICATION-TOKEN: "sdfdfs"
//       X-AUTH-ID: "asdasdsad"
//       X-CLIENT-ID: "asd2e2e3"
//
//     Responses:
//       200: cardsList
func ListCards(w http.ResponseWriter, r *http.Request) {
	log.Println(322)
	cardsRepo := r.Context().Value(middlewares.CardsRepoKey).(interfaces.CardsRepo)
	customer := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	allUser, err := cardsRepo.GetByCustomerID(customer.ID.Hex())
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(allUser, "", "    ")
		log.Println("All customers cards")
		log.Println(string(j))
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allUser)
}
