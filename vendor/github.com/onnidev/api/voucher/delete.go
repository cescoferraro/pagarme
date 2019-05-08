package voucher

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
	"gopkg.in/mgo.v2/bson"
)

// Delete carinho
func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "voucherId")
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	err := vouchersCollection.Collection.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, id)
}
