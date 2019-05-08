package menuTicket

import (
	"errors"
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

// Delete is commented
func Delete(w http.ResponseWriter, r *http.Request) {
	repo, ok := r.Context().Value(middlewares.ClubMenuTicketRepoKey).(interfaces.ClubMenuTicketRepo)
	if !ok {
		err := errors.New("assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if bson.IsObjectIdHex(chi.URLParam(r, "menuId")) {
		id := bson.ObjectIdHex(chi.URLParam(r, "menuId"))
		now := types.Timestamp(time.Now())
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"updateDate": &now,
					"status":     "INACTIVE",
				}},
			ReturnNew: true,
		}
		patchedTicket := types.ClubMenuTicket{}
		_, err := repo.Collection.Find(bson.M{"_id": id}).Apply(change, &patchedTicket)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		render.JSON(w, r, id.Hex())
		return
	}
	err := errors.New("not a valid objectid")
	shared.MakeONNiError(w, r, 400, err)
	return
}
