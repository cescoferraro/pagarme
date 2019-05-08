package invitedCustomer

import (
	"errors"
	"log"
	"net/http"
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

// Link skdjfn
func Link(w http.ResponseWriter, r *http.Request) {
	req, ok := r.Context().Value(middlewares.InvitedLinkCustomerPostKey).(types.InvitedLinkCustomerPost)
	if !ok {
		err := errors.New("bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(req)
	inviterepo, ok := r.Context().Value(middlewares.InvitedCustomerRepoKey).(interfaces.InvitedCustomerRepo)
	if !ok {
		err := errors.New("bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	id := chi.URLParam(r, "inviteId")
	invite, err := inviterepo.GetByID(id)
	if err != nil {
		log.Println("not found invite")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if invite.LinkedCustomer != nil {
		err := errors.New("inviteCustomer mail already linked")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	customer, err := onni.FindCustomerByMail(r.Context(), req.Mail)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	var patchedInvite types.InvitedCustomer
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate":     &now,
				"linkedCustomer": &customer.ID,
				"done":           true,
				"assignedMail":   &customer.Mail,
			}},
		ReturnNew: true,
	}
	_, err = inviterepo.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &patchedInvite)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	repo := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	vouchers, err := repo.GetByCustomer(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	horario := types.Timestamp(time.Now())
	for _, voucher := range vouchers {
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"updateDate":   &horario,
					"status":       "AVAILABLE",
					"customerId":   customer.ID,
					"customerName": customer.Name(),
				}},
			ReturnNew: true,
		}
		result := types.Voucher{}
		_, err := repo.Collection.Find(bson.M{"_id": voucher.ID}).Apply(change, &result)
		if err != nil {
			log.Println("nao achou o voucher")
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
	promos, err := promoRepo.ByCustomer(id)
	if err != nil {
		log.Println("nao achou promos by log")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	for _, promo := range promos {
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"updateDate":   &horario,
					"customerId":   customer.ID,
					"customerName": customer.Name(),
					"customerMail": customer.Mail,
				}},
			ReturnNew: true,
		}
		result := types.PromotionalCustomer{}
		_, err := promoRepo.Collection.Find(bson.M{"_id": promo.ID}).Apply(change, &result)
		if err != nil {
			log.Println("nao achou a prmo")
			shared.MakeONNiError(w, r, 400, err)
			return
		}
	}
	render.JSON(w, r, customer)
}
