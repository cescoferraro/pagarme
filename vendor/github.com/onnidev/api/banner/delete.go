package banner

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
	"gopkg.in/mgo.v2/bson"
)

// Delete sdfkjn
func Delete(w http.ResponseWriter, r *http.Request) {
	bannerRepo := r.Context().Value(middlewares.BannerRepoKey).(interfaces.BannerRepo)
	id := chi.URLParam(r, "bannerID")
	err := bannerRepo.Collection.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, id)
}
