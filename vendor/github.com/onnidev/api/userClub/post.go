package userClub

import (
	"log"
	"net/http"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"gopkg.in/mgo.v2/bson"
)

// Post carinho
func Post(w http.ResponseWriter, r *http.Request) {
	productsCollection := r.Context().Value(middlewares.ProductsRepoKey).(interfaces.ProductsRepo)
	productReq := r.Context().Value(middlewares.ReadProductPostKey).(types.ProductPostRequest)
	horario := types.Timestamp(time.Now())
	product := types.Product{
		ID:           bson.ObjectIdHex(productReq.ID),
		CreationDate: &horario,
		Name:         productReq.Name,
		NameSort:     productReq.Name,
		Type:         productReq.Type,
		Image: types.Image{
			FileID:       bson.ObjectIdHex(productReq.Image),
			MimeType:     "IMAGE_PNG",
			CreationDate: &horario,
		},
	}
	err := productsCollection.Collection.Insert(product)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	log.Println(productReq, productsCollection)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, 33)
}
