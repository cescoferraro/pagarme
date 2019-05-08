package product

import (
	"errors"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// UpdateImage TODO: NEEDS COMMENT INFO
func UpdateImage(w http.ResponseWriter, r *http.Request) {
	const mB = 1 << 20
	err := r.ParseMultipartForm(2 * mB)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	rfile, _, err := r.FormFile("image")
	fileHeader := make([]byte, r.ContentLength)
	_, err = rfile.Read(fileHeader)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	image, err := onni.CreateImage(r.Context(), fileHeader)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	repo, ok := r.Context().Value(middlewares.ProductsRepoKey).(interfaces.ProductsRepo)
	if !ok {
		err := errors.New("bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate": &now,
				"image":      image,
			}},
		ReturnNew: true,
	}
	patchedProduct := types.Product{}
	_, err = repo.Collection.Find(bson.M{"_id": bson.ObjectIdHex(chi.URLParam(r, "productId"))}).Apply(change, &patchedProduct)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, patchedProduct)
}
