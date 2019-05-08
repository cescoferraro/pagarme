package customer

import (
	"errors"
	"net/http"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UpdatePassword sdkjfn
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	req := r.Context().Value(middlewares.UserClubReq).(types.ChangePasswordRequest)
	customer := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	if shared.EncryptPassword2(req.Password) != customer.Password {
		err := errors.New("password does not match")
		shared.MakeONNiError(w, r, 400, err)

		return
	}

	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"password":   shared.EncryptPassword2(req.NewPassword),
				"updateDate": &now,
			}},
		ReturnNew: true,
	}
	var result types.Customer
	_, err := repo.Collection.Find(bson.M{"_id": customer.ID}).Apply(change, &result)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
