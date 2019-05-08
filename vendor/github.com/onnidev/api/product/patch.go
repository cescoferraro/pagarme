package product

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"gopkg.in/mgo.v2/bson"
)

// Patch carinho
func Patch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productId")
	productsCollection := r.Context().Value(middlewares.ProductsRepoKey).(interfaces.ProductsRepo)
	productReq := r.Context().Value(middlewares.ReadProductPatchKey).(types.ProductPatchRequest)
	gridfs := r.Context().Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	product, err := productsCollection.GetByID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	if product.Image.FileID.Hex() != productReq.Image {
		log.Println("have changed")
		file, err := gridfs.FS.OpenId(bson.ObjectIdHex(productReq.Image))
		if err != nil {
			render.Status(r, http.StatusExpectationFailed)
			render.JSON(w, r, err.Error())
			return
		}
		err = productsCollection.PatchWithImage(id, file, productReq)
		if err != nil {
			render.Status(r, http.StatusExpectationFailed)
			render.JSON(w, r, err.Error())
			return
		}
		return

	}
	log.Println("havent changed")
	err = productsCollection.Patch(id, productReq)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, product)
}
