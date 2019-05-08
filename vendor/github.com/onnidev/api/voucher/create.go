package voucher

import (
	"log"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Create carinho
func Create(w http.ResponseWriter, r *http.Request) {
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
	log.Println(voucherReq.PartyProductID)
	partyProduct, err := partyProductCollection.GetByID(voucherReq.PartyProductID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	customersRepo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	customers := []types.Customer{}
	nocustomers := []string{}
	for _, mail := range voucherReq.Emails {
		email := shared.NormalizeEmail(mail)
		user, err := customersRepo.GetByEmail(email)
		if err != nil {
			if err.Error() == "not found" {
				log.Println("not found  by email requested")
				inviterepo, ok := r.Context().Value(middlewares.InvitedCustomerRepoKey).(interfaces.InvitedCustomerRepo)
				if !ok {
					log.Println("bug problem")
					nocustomers = append(nocustomers, email)
					continue
				}
				invite, err := inviterepo.GetByMail(shared.NormalizeEmail(mail))
				if err != nil {
					log.Println("nbow achei por emIl")
					nocustomers = append(nocustomers, email)
					continue
				}
				if invite.Done {
					if invite.LinkedCustomer != nil {
						id := *invite.LinkedCustomer
						customer, err := customersRepo.GetByID(id.Hex())
						if err != nil {
							log.Println("nao peguei por idd linked")
							nocustomers = append(nocustomers, email)
							continue
						}
						customers = append(customers, customer)
						continue
					}
					customer, err := customersRepo.GetByID(invite.ID.Hex())
					if err != nil {
						log.Println("nao peguei por idd full")
						nocustomers = append(nocustomers, email)
						continue
					}
					customers = append(customers, customer)
					continue
				}
				nocustomers = append(nocustomers, email)
				continue
			}
			nocustomers = append(nocustomers, email)
			continue
		}
		customers = append(customers, user)
	}
	log.Println(customers)
	log.Println(nocustomers)
	vouchers, err := onni.CreateVoucher(r.Context(), club, party, voucherReq, partyProduct, customers)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if len(nocustomers) > 0 {
		render.Status(r, 206)
		render.JSON(w, r, nocustomers)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, vouchers)
}
