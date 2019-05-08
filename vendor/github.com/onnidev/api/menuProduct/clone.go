package menuProduct

import (
	"errors"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// Clone is commented
func Clone(w http.ResponseWriter, r *http.Request) {
	if bson.IsObjectIdHex(chi.URLParam(r, "menuId")) {
		id := bson.ObjectIdHex(chi.URLParam(r, "menuId"))
		repo, ok := r.Context().Value(middlewares.ClubMenuProductRepoKey).(interfaces.ClubMenuProductRepo)
		if !ok {
			err := errors.New("assert")
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		menu, err := repo.GetByID(id.Hex())
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		menu.ID = bson.NewObjectId()
		err = repo.Collection.Insert(menu)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, menu.ID.Hex())
		return
	}
	err := errors.New("not a valid objectid")
	shared.MakeONNiError(w, r, 400, err)
	return
}
