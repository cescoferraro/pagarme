package voucher

import (
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Invite carinho
func Invite(w http.ResponseWriter, r *http.Request) {
	log.Println("creating voucher endpoint")
	voucherReq := r.Context().Value(middlewares.VouchersPostRequestKey).(types.VoucherPostRequest)
	partiesCollection := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	party, err := partiesCollection.GetByID(voucherReq.PartyID)
	if err != nil {
		log.Println("party")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	clubCollection := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	club, err := clubCollection.GetByID(party.ClubID.Hex())
	if err != nil {
		log.Println("club")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	partyProductCollection := r.Context().Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	partyProduct, err := partyProductCollection.GetByID(voucherReq.PartyProductID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	inviterepo, ok := r.Context().Value(middlewares.InvitedCustomerRepoKey).(interfaces.InvitedCustomerRepo)
	if !ok {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	voucher := types.Voucher{}
	for _, email := range voucherReq.Emails {
		invite, err := inviterepo.GetByMail(email)
		if err != nil {
			if err.Error() == "not found" {
				id := bson.NewObjectId()
				now := types.Timestamp(time.Now())
				err := inviterepo.Collection.Insert(types.InvitedCustomer{
					CreationDate: &now,
					Mail:         strings.ToLower(email),
					ID:           id,
				})
				if err != nil {
					shared.MakeONNiError(w, r, 400, err)
					return
				}
				voucher, err = onni.CreateInviteVoucher(r.Context(), email, club, party, voucherReq, partyProduct, id.Hex(), email, "PENDING")
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
			customersRepo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
			customer, err := customersRepo.GetByID(id.Hex())
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			voucher, err = onni.CreateInviteVoucher(r.Context(), customer.Mail, club, party, voucherReq, partyProduct, id.Hex(), customer.Name(), "AVAILABLE")
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			continue
		}
		if invite.Done {

			customersRepo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
			customer, err := customersRepo.GetByID(invite.ID.Hex())
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			voucher, err = onni.CreateInviteVoucher(r.Context(), customer.Mail, club, party, voucherReq, partyProduct, invite.ID.Hex(), customer.Name(), "AVAILABLE")
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			continue
		}
		voucher, err = onni.CreateInviteVoucher(r.Context(), email, club, party, voucherReq, partyProduct, invite.ID.Hex(), email, "PENDING")
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, voucher)
}
