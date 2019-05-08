package voucher

import (
	"log"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// AppByCustomer dskfjnsd
func AppByCustomer(w http.ResponseWriter, r *http.Request) {
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	customer := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	vouchers, err := vouchersCollection.GetActualAppVoucherByCustomer(customer.ID.Hex())
	if err != nil {
		log.Println(err.Error())
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	for _, voucher := range vouchers {
		if voucher.Responsable != nil {
			voucher.Responsable.Clubs = nil
		}
		if voucher.Customer != nil {
			if voucher.Customer.DocumentNumber != nil {
				if *voucher.Customer.DocumentNumber == "" {
					zero := "00000000000"
					voucher.Customer.DocumentNumber = &zero
					continue
				}
				continue
			}
			zero := "00000000000"
			voucher.Customer.DocumentNumber = &zero
		}
	}
	result := []types.AppCompleteVoucher{}
	for _, voucher := range vouchers {
		if voucher.Party.Status == "ACTIVE" {
			// if voucher.Club != nil {
			// 	if voucher.Club.Status == "ACTIVE" {
			result = append(result, voucher)
			// 	}
			// }
		}
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
