package voucher

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// ByPartySoft is commented
func ByPartySoft(w http.ResponseWriter, r *http.Request) {
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	id := chi.URLParam(r, "partyId")
	vouchers, err := vouchersCollection.GetByParty(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	userClub, ok := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	if !ok {
		err := errors.New("asser bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(len(vouchers))
	result := []types.SoftVoucher{}
	for _, voucher := range vouchers {
		instance := types.SoftVoucher{
			Price:                  voucher.Price.Value,
			Product:                voucher.Product,
			Name:                   voucher.Product.Name,
			Status:                 voucher.Status,
			Type:                   voucher.Type,
			VoucherID:              voucher.ID,
			CreationDate:           voucher.CreationDate,
			TransferedFrom:         &voucher.TransferedFrom,
			VoucherUseDate:         voucher.VoucherUseDate,
			VoucherUseUserClubID:   voucher.VoucherUseUserClubID,
			VoucherUseUserClubName: voucher.VoucherUseUserClubName,
			CustomerID:             voucher.CustomerID,
			CustomerName:           voucher.CustomerName,
		}
		if voucher.Customer != nil {
			instance.CustomerMail = voucher.Customer.Mail
		}
		if voucher.Responsable != nil {
			instance.ResponsableUserClubID = &voucher.Responsable.ID
			instance.ResponsableUserClubName = &voucher.Responsable.Name
		}
		if userClub.Profile == "PROMOTER" {
			log.Println("eu sou promoter")
			if voucher.Responsable != nil {
				if voucher.Responsable.ID == userClub.ID {
					log.Println("got a voucher")
					result = append(result, instance)
				}
			}
			continue
		}
		if instance.Status == "ERROR" {
			if userClub.Profile == "ONNI" {
				result = append(result, instance)
			}
			continue
		}
		result = append(result, instance)

	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
