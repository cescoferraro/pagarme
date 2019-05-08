package userClub

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/bradfitz/slice"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// VoucherReadByUserClub sdkjfn
func VoucherReadByUserClub(w http.ResponseWriter, r *http.Request) {
	vouchersRepo := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	user := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	if !shared.Contains([]string{"ATTENDANT", "ADMIN"}, user.Profile) {
		err := errors.New("usuário não faz leitura")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	vouchers, err := vouchersRepo.ReadByUserClubID(user.ID.Hex())
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	log.Println(len(vouchers))
	result := make(map[string][]types.CompleteVoucher)
	for _, voucher := range vouchers {
		result[voucher.PartyID.Hex()] = append(result[voucher.PartyID.Hex()], voucher)
	}
	type SectionListVouchers struct {
		Party     string                  `json:"party" bson:"party"`
		StartDate *types.Timestamp        `json:"startDate" bson:"startDate"`
		Data      []types.CompleteVoucher `json:"data" bson:"data"`
	}
	var trueR []SectionListVouchers
	for key, party := range result {
		log.Println(key)
		slice.Sort(party[:], func(i, j int) bool {
			return party[i].VoucherUseDate.Time().After(party[j].VoucherUseDate.Time())
		})
		empty := types.Timestamp(time.Time{})
		startDate := &empty
		if len(party) > 0 {
			festa := *party[0].Party
			startDate = festa.StartDate
		}
		trueR = append(trueR, SectionListVouchers{
			Party:     key,
			StartDate: startDate,
			Data:      party,
		})
	}

	slice.Sort(trueR[:], func(i, j int) bool {
		return trueR[i].StartDate.Time().After(trueR[j].StartDate.Time())
	})
	render.Status(r, http.StatusOK)
	render.JSON(w, r, trueR)
}

// SimpleVoucherReadByUserClub sdkjfn
func SimpleVoucherReadByUserClub(w http.ResponseWriter, r *http.Request) {
	vouchersRepo := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	user, ok := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	if !ok {
		err := errors.New("assert error")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if !shared.Contains([]string{"ATTENDANT", "ADMIN"}, user.Profile) {
		err := errors.New("usuário não faz leitura")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	vouchers, err := vouchersRepo.ReadByUserClubID(user.ID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	result := []types.VoucherHistoryResume{}
	for _, voucher := range vouchers {
		party := *voucher.Party
		club := *voucher.Club
		customer := *voucher.Customer
		uri := "https://user-images.githubusercontent.com/36003926/39221162-d154a218-480b-11e8-90c2-490cb5993985.png"
		if customer.FacebookID != "" {
			uri = `https://graph.facebook.com/` + customer.FacebookID + `/picture?type=square&width=300`
		}
		result = append(result, types.VoucherHistoryResume{
			ID:                     voucher.ID,
			CreationDate:           voucher.CreationDate,
			UpdateDate:             voucher.UpdateDate,
			StartDate:              voucher.StartDate,
			EndDate:                voucher.EndDate,
			VoucherUseDate:         voucher.VoucherUseDate,
			PartyName:              party.Name,
			ClubName:               club.Name,
			CustomerName:           voucher.CustomerName,
			VoucherUseUserClubID:   voucher.VoucherUseUserClubID,
			ResponsableUserClubID:  voucher.ResponsableUserClubID,
			VoucherUseUserClubName: voucher.VoucherUseUserClubName,
			Status:                 voucher.Status,
			Price:                  voucher.Price,
			Product:                voucher.Product,
			ProductName:            voucher.Product.Name,
			Type:                   voucher.Type,
			CustomerImage:          uri,
			ProductImage:           "https://s3-sa-east-1.amazonaws.com/onni-medium-images/" + voucher.Product.Image.FileID.Hex(),
			PartyImage:             "https://s3-sa-east-1.amazonaws.com/onni-medium-images/" + party.BackgroundImage.FileID.Hex(),
			ClubImage:              "https://s3-sa-east-1.amazonaws.com/onni-medium-images/" + club.BackgroundImage.FileID.Hex(),
		})

	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
