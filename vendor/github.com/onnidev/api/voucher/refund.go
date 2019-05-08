package voucher

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Refund TODO: NEEDS COMMENT INFO
func Refund(w http.ResponseWriter, r *http.Request) {
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	id := chi.URLParam(r, "voucherId")
	voucher, err := vouchersCollection.GetSimpleByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if shared.Contains([]string{"ERROR", "CANCELED"}, voucher.Status) {
		err := errors.New("not a plausible voucher status")
		shared.MakeONNiError(w, r, 520, err)
		return
	}
	// if voucher.InvoiceID != nil {
	// 	invoice, err := onni.Invoice(r.Context(), voucher.InvoiceID.Hex())
	// 	if err != nil {
	// 		shared.MakeONNiError(w, r, 400, err)
	// 		return
	// 	}
	// 	log.Println(invoice.ID.Hex())
	// 	api := pagarme.New(viper.GetString("PAGARME"))
	// 	err = api.Refund(r.Context(), invoice.ID.Hex())
	// 	if err != nil {
	// 		shared.MakeONNiError(w, r, 400, err)
	// 		return
	// 	}
	// 	render.Status(r, http.StatusOK)
	// 	render.JSON(w, r, invoice)

	// 	return
	// }
	used, err := vouchersCollection.UseVoucher(voucher, userClub)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, used)
}
