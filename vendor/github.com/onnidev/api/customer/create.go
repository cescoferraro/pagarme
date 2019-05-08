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

// CreateEndpoint TODO: NEEDS COMMENT INFO
func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
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

	inviterepo, ok := r.Context().Value(middlewares.InvitedCustomerRepoKey).(interfaces.InvitedCustomerRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	invite, err := inviterepo.GetByID(req.ID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	mail := shared.NormalizeEmail(req.Mail)
	invite, err = inviterepo.ChangeEmail(req.ID, mail)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	now := types.Timestamp(time.Now())
	customer := types.Customer{
		ID:             bson.ObjectIdHex(req.ID),
		CreationDate:   &now,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Mail:           mail,
		Password:       req.Password,
		Phone:          req.Phone,
		UserName:       strings.Replace(req.UserName, "@", "", -1),
		DocumentNumber: &req.DocumentNumber,
		BirthDate:      req.BirthDate,
		FacebookID:     *invite.FBID,
		FavoriteClubs:  []bson.ObjectId{},
	}
	err = repo.Collection.Insert(customer)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(invite, "", "    ")
		log.Println("Customer created on MongoDB")
		log.Println(string(j))
	}
	vouchersCollection, ok := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	_, err = inviterepo.Done(req.ID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	vouchers, err := vouchersCollection.GetAllCustomerVouchers(req.ID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	for _, voucher := range vouchers {
		log.Println(55655)
		_, err := vouchersCollection.SetInvitedNameOnVoucher(voucher.ID.Hex(), req.FirstName+" "+req.LastName)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
	}
	promoRepo, ok := r.Context().Value(middlewares.PromotionalCustomerRepoKey).(interfaces.PromotionalCustomerRepo)
	if !ok {
		err := errors.New("2bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	promos, err := promoRepo.ByCustomer(req.ID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	for _, promo := range promos {
		_, err := promoRepo.SetInvitedNameOnPromotion(promo.ID.Hex(), customer)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
	}
	render.Status(r, 200)
	render.JSON(w, r, customer)
}
