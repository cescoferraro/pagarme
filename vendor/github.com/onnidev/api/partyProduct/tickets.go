package partyProduct

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

// Products carinho
func TicketProducts(w http.ResponseWriter, r *http.Request) {
	productsCollection := r.Context().Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	id := chi.URLParam(r, "partyId")
	allUser, err := productsCollection.GetTicketsByPartyID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(allUser, "", "    ")
		log.Println("All products from club")
		log.Println(string(j))
		log.Println("Tamanho", len(allUser))
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allUser)
}
