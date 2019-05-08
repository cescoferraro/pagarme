package promotionalCustomer

import (
	"errors"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Invite is a comented function
func Invite(w http.ResponseWriter, r *http.Request) {
	req, ok := r.Context().Value(middlewares.PromotionalCustomerPostQueryKey).(types.PromotionalCustomerPost)
	if !ok {
		err := errors.New("1bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	userClub, ok := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	if !ok {
		err := errors.New("3bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	customersRepo, ok := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	if !ok {
		err := errors.New("4bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	productsCollection, ok := r.Context().Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("4bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	inviterepo, ok := r.Context().Value(middlewares.InvitedCustomerRepoKey).(interfaces.InvitedCustomerRepo)
	if !ok {
		err := errors.New("4bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	// nocustomer := []string{}
	// already := []string{}
	// promotions := []types.PromotionalCustomer{}
	for _, promotion := range req.PromotionIDS {
		partyProduct, promo, err := productsCollection.GetPromotion(promotion)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		party, err := onni.Party(r.Context(), partyProduct.PartyID.Hex())
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		for _, mail := range req.Mails {
			email := shared.NormalizeEmail(mail)
			invite, err := inviterepo.GetByMail(email)
			if err != nil {
				if err.Error() == "not found" {
					log.Println("not found invite for: ", email)
					id := bson.NewObjectId()
					now := types.Timestamp(time.Now())
					err := inviterepo.Collection.Insert(types.InvitedCustomer{
						CreationDate: &now,
						Mail:         email,
						ID:           id,
					})
					if err != nil {
						shared.MakeONNiError(w, r, 400, err)
						return
					}
					// create invitePromotion
					prom := types.PromotionalCustomer{
						ID:           bson.NewObjectId(),
						CreationDate: &now,
						PromotionID:  promo.ID,
						PromoterID:   userClub.ID,
						PromoterName: userClub.Name,
						CustomerID:   id,
						CustomerMail: email,
						CustomerName: "[INVITED] " + email,
					}
					err = onni.PersistPromotionalCustomer(r.Context(), party, partyProduct, prom, email)
					if err != nil {
						shared.MakeONNiError(w, r, 400, err)
						return
					}
					continue
				}
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			if invite.LinkedCustomer != nil {
				id := *invite.LinkedCustomer
				customer, err := customersRepo.GetByID(id.Hex())
				if err != nil {
					shared.MakeONNiError(w, r, 400, err)

					return
				}
				now := types.Timestamp(time.Now())
				prom := types.PromotionalCustomer{
					ID:           bson.NewObjectId(),
					CreationDate: &now,
					PromotionID:  promo.ID,
					PromoterID:   userClub.ID,
					PromoterName: userClub.Name,
					CustomerID:   customer.ID,
					CustomerMail: customer.Mail,
					CustomerName: customer.Name(),
				}
				err = onni.PersistPromotionalCustomer(r.Context(), party, partyProduct, prom, customer.Mail)
				if err != nil {
					shared.MakeONNiError(w, r, 400, err)
					return
				}
				continue
			}
			if invite.Done {
				customer, err := customersRepo.GetByID(invite.ID.Hex())
				if err != nil {
					shared.MakeONNiError(w, r, 400, err)
					return
				}
				now := types.Timestamp(time.Now())
				prom := types.PromotionalCustomer{
					ID:           bson.NewObjectId(),
					CreationDate: &now,
					PromotionID:  promo.ID,
					PromoterID:   userClub.ID,
					PromoterName: userClub.Name,
					CustomerID:   customer.ID,
					CustomerMail: customer.Mail,
					CustomerName: customer.Name(),
				}
				err = onni.PersistPromotionalCustomer(r.Context(), party, partyProduct, prom, customer.Mail)
				if err != nil {
					shared.MakeONNiError(w, r, 400, err)
					return
				}
				continue
			}
			// create invitePromotion
			now := types.Timestamp(time.Now())
			prom := types.PromotionalCustomer{
				ID:           bson.NewObjectId(),
				CreationDate: &now,
				PromotionID:  promo.ID,
				PromoterID:   userClub.ID,
				PromoterName: userClub.Name,
				CustomerID:   invite.ID,
				CustomerMail: invite.Mail,
				CustomerName: "[INVITED] " + email,
			}
			err = onni.PersistPromotionalCustomer(r.Context(), party, partyProduct, prom, email)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}

		}
	}

	render.Status(r, 200)
	render.JSON(w, r, 33)
}
