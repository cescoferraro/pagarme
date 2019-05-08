package product

import (
	"net/http"
	"strings"
	"time"

	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"gopkg.in/mgo.v2/bson"
)

// Add TODO: NEEDS COMMENT INFO
func Add(w http.ResponseWriter, r *http.Request) {
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
	horario := types.Timestamp(time.Now())
	product := types.Product{
		ID:           bson.NewObjectId(),
		CreationDate: &horario,
		Name:         r.FormValue("name"),
		Deprecated:   false,
		NameSort:     strings.ToLower(r.FormValue("name")),
		Type:         r.FormValue("type"),
		Image:        image,
		Category:     r.FormValue("category"),
	}
	err = onni.CreateProduct(r.Context(), product)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, product)
}
