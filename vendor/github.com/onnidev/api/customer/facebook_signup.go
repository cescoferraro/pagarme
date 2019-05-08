package customer

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"gopkg.in/mgo.v2/bson"
)

// FacebookSignUp sdkjfn
func FacebookSignUp(w http.ResponseWriter, r *http.Request) {
	facebookSignUpReq := r.Context().
		Value(middlewares.FacebookSignUpRequestKey).(types.FacebookSignUpRequest)
	validation, err := onni.FacebookAppValidate(facebookSignUpReq.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	customerCollection := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	cust, err := customerCollection.FacebookIDExists(validation.Data.UserID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if cust != 0 {
		err := errors.New("user already")
		http.Error(w, err.Error(), 400)
		return
	}
	id := bson.NewObjectId()
	horario := types.Timestamp(time.Now())
	birth := types.Timestamp(facebookSignUpReq.BirthDate.Time())
	customer := types.Customer{
		ID:             id,
		CreationDate:   &horario,
		BirthDate:      &birth,
		FacebookID:     validation.Data.UserID,
		FirstName:      facebookSignUpReq.FirstName,
		LastName:       facebookSignUpReq.LastName,
		Mail:           strings.ToLower(facebookSignUpReq.Mail),
		Phone:          strings.Replace(facebookSignUpReq.Phone, " ", "", -1),
		UserName:       facebookSignUpReq.UserName,
		DocumentNumber: &facebookSignUpReq.DocumentNumber,
		Password:       shared.EncryptPassword2(shared.RangeIn(100000, 999999)),
		FavoriteClubs:  []bson.ObjectId{},
	}
	err = customerCollection.Collection.Insert(customer)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	response, err := customer.LogInCustomer()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
