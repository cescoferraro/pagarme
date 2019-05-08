package appclub

import (
	"errors"
	"log"
	"net/http"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Read TODO: NEEDS COMMENT INFO
func Read(w http.ResponseWriter, r *http.Request) {
	log.Println(">>>>>> adentrei o read endpoint")
	req, ok := r.Context().Value(middlewares.VoucherSoftReadReqKey).(types.VoucherSoftReadReq)
	if !ok {
		err := errors.New("bug assert")
		render.Status(r, 520)
		render.JSON(w, r, err.Error())
		return
	}
	userClub, err := onni.UserClub(r.Context(), req.UserClubID)
	if err != nil {
		render.Status(r, 520)
		render.JSON(w, r, err.Error())
		return
	}
	log.Println("clubs")
	for _, club := range userClub.Clubs {
		log.Println(club)
	}
	log.Println("clubs")
	log.Println("clubs")
	log.Println("clubs")
	voucher, err := onni.VoucherUseComplete(r.Context(), req.VoucherID, true, onni.AllKindsConstrain, userClub)
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
