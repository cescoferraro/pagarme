package product

import (
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// SoftPatch carinho
func SoftPatch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productId")
	repo := r.Context().Value(middlewares.ProductsRepoKey).(interfaces.ProductsRepo)
	productReq := r.Context().Value(middlewares.ReadProductSoftPatchKey).(types.ProductSoftPatchRequest)
	product, err := repo.GetByID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate": &now,
				"type":       shared.OrBlank(product.Type, productReq.Type),
				"name":       shared.OrBlank(product.Name, productReq.Name),
				"category":   shared.OrBlank(product.Category, productReq.Category),
			}},
		ReturnNew: true,
	}
	patchedProduct := types.Product{}
	_, err = repo.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &patchedProduct)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, patchedProduct)
}
