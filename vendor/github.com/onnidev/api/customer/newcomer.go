package customer

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// NewComerEndpoint TODO: NEEDS COMMENT INFO
func NewComerEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("entrei no endpoint")
	log.Println("entrei no endpoint")
	log.Println("entrei no endpoint")
	repo, ok := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	req, ok := r.Context().Value(middlewares.CustomerRequestKey).(types.CustomerPostRequest)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	mail := shared.NormalizeEmail(req.Mail)
	now := types.Timestamp(time.Now())
	customer := types.Customer{
		ID:             bson.NewObjectId(),
		CreationDate:   &now,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Mail:           mail,
		Password:       req.Password,
		Phone:          req.Phone,
		UserName:       strings.Replace(req.UserName, "@", "", -1),
		DocumentNumber: &req.DocumentNumber,
		BirthDate:      req.BirthDate,
		FacebookID:     req.FacebookID,
		FavoriteClubs:  []bson.ObjectId{},
	}
	err := repo.Collection.Insert(customer)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(customer, "", "    ")
		log.Println("Customer created on MongoDB")
		log.Println(string(j))
	}

	token, err := customer.GenerateToken()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, 200)
	render.JSON(w, r, token)
}
