package product

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Products carinho
func Products(w http.ResponseWriter, r *http.Request) {
	productsCollection := r.Context().Value(middlewares.ProductsRepoKey).(interfaces.ProductsRepo)
	allUser, err := productsCollection.List()
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(allUser, "", "    ")
		log.Println("All onni products")
		log.Println(string(j))
		log.Println("Tamanho", len(allUser))
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allUser)
}
