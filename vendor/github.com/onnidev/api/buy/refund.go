package buy

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Refund is a comented function
func Refund(w http.ResponseWriter, r *http.Request) {
	voucherID := chi.URLParam(r, "voucherId")
	log.Println(voucherID)
	voucher, err := onni.Voucher(r.Context(), voucherID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	party, err := onni.Party(r.Context(), voucher.PartyID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	club, err := onni.Club(r.Context(), voucher.ClubID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	userClub, ok := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	if !ok {
		err := errors.New("assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	err = onni.ValidateVoucherRefund(r.Context(), voucher, userClub)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(">>>>>> separar os produtos")
	products, err := onni.BuyPostProductsTyped(r.Context(), voucher.BuyPartyProductsItemRequest())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Printf(">>>>>> ONNi %v %v\n", products.SumONNi(party, club), products.SumClub(party, club))
	log.Printf(">>>>>> ONNi %v %v\n", products.SumONNiString(party, club), products.SumClubString(party, club))
	if voucher.InvoiceID != nil {
		err = onni.RefundVoucher(r.Context(), voucher, products)
		if err != nil {
			if err.Error() != "Transação já estornada" {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
		}
	}
	repo, ok := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	used, err := repo.CancelVoucher(voucher, userClub)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, 200)
	render.JSON(w, r, used)

}
