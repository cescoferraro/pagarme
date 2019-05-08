package appclub

import (
	"errors"
	"log"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Validate TODO: NEEDS COMMENT INFO
func Validate(w http.ResponseWriter, r *http.Request) {
	log.Println(">>>>>> adentrei o validate endpoint")
	req, ok := r.Context().Value(middlewares.VoucherSoftValidateReqKey).(types.VoucherSoftValidateReq)
	if !ok {
		err := errors.New("bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	vouchersCollection, ok := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	voucher, err := vouchersCollection.GetByID(req.VoucherID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	err = onni.VoucherCheckStatus(voucher.Status)
	if err != nil {
		render.Status(r, 520)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, struct {
		CustomerName string `json:"customerName"`
		ProductName  string `json:"productName"`
	}{
		CustomerName: voucher.CustomerName,
		ProductName:  voucher.Product.Name,
	})
}
