package customer

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Activate TODO: NEEDS COMMENT INFO
func Activate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "customerId")
	userclub, err := onni.Customer(ctx, id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	mode := chi.URLParam(r, "mode")
	if !shared.Contains([]string{"active", "inactive"}, mode) {
		err := errors.New("not a plausible mode")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	userClubADMIN := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	now := types.Timestamp(time.Now())
	trusted := strings.ToUpper(mode)
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate": &now,
			"updatedBy":  &userClubADMIN.ID,
			"trusted":    &trusted,
		}},
		ReturnNew: true,
	}
	repo, ok := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	var patchedCustomer types.Customer
	_, err = repo.Collection.Find(bson.M{"_id": userclub.ID}).Apply(change, &patchedCustomer)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, true)
}
