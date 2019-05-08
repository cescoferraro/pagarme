package onni

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// CompleteVoucher siksjdnf
func CompleteVoucher(ctx context.Context, partyID string) (types.CompleteVoucher, error) {
	var voucher types.CompleteVoucher
	repo, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		return voucher, err
	}
	voucher, err := repo.GetByID(partyID)
	if err != nil {
		return voucher, err
	}
	return voucher, nil
}

// CancelVoucher TODO: NEEDS COMMENT INFO
func CancelVoucher(ctx context.Context, voucher types.Voucher, userClub types.UserClub) error {
	repo, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	_, err := repo.CancelVoucher(voucher, userClub)
	if err != nil {
		return err
	}
	return nil
}

// Voucher siksjdnf
func Voucher(ctx context.Context, partyID string) (types.Voucher, error) {
	var voucher types.Voucher
	repo, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		return voucher, err
	}
	voucher, err := repo.GetSimpleByID(partyID)
	if err != nil {
		return voucher, err
	}
	return voucher, nil
}

// FutureCustomerVouchers TODO: NEEDS COMMENT INFO
func FutureCustomerVouchers(ctx context.Context, customer string) ([]types.CompleteVoucher, error) {
	voucherRepo := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	voucherCount, err := voucherRepo.FutureByCustomer(customer)
	if err != nil {
		return voucherCount, err
	}
	return voucherCount, nil
}

// CustomerVouchers TODO: NEEDS COMMENT INFO
func CustomerVouchers(ctx context.Context, customer types.Customer, partyID string) ([]types.CompleteVoucher, error) {
	voucherRepo := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	voucherCount, err := voucherRepo.ByPartyAndCustomer(partyID, customer.ID.Hex())
	if err != nil {
		return voucherCount, err
	}
	return voucherCount, nil
}

// HowManyVoucheraCustomerHas TODO: NEEDS COMMENT INFO
func HowManyVoucheraCustomerHas(ctx context.Context, customer types.Customer, partyID string) (int, error) {
	voucherRepo := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	voucherCount, err := voucherRepo.CountByPartyAndCustomer(partyID, customer.ID.Hex())
	if err != nil {
		return 0, err
	}
	return voucherCount, nil
}

// CreateVoucher TODO: NEEDS COMMENT INFO
func CreateVoucher(
	ctx context.Context,
	club types.Club,
	party types.Party,
	req types.VoucherPostRequest,
	partyProduct types.PartyProduct,
	customers []types.Customer,
) ([]types.Voucher, error) {
	userClub := ctx.Value(middlewares.UserClubKey).(types.UserClub)
	var vouchers []types.Voucher
	vouchersCollection, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		return vouchers, err
	}
	horario := types.Timestamp(time.Now())
	zero := types.Timestamp(time.Time{})
	for _, customer := range customers {
		go MailNewVoucher(customer.Mail, party, partyProduct, req, customer.ID.Hex(), false)
		for i := 1; i <= req.Quantity; i++ {
			vouch := types.Voucher{
				ID:                    bson.NewObjectId(),
				CreationDate:          &horario,
				CustomerID:            customer.ID,
				PartyID:               bson.ObjectIdHex(req.PartyID),
				ClubID:                club.ID,
				PartyProductID:        bson.ObjectIdHex(req.PartyProductID),
				Type:                  req.Type,
				ClubName:              club.Name,
				PartyName:             party.Name,
				CustomerName:          customer.FirstName + " " + customer.LastName,
				StartDate:             party.StartDate,
				VoucherUseDate:        &zero,
				EndDate:               party.EndDate,
				Status:                "AVAILABLE",
				ResponsableUserClubID: &userClub.ID,
				Price: partyProduct.MoneyAmount,
				Product: types.VoucherProduct{
					// Image: partyProduct.Image,
					Type: partyProduct.Type,
					Name: partyProduct.Name,
				},
			}
			if partyProduct.Image != nil {
				vouch.Product.Image = *partyProduct.Image
			}
			if partyProduct.Type == "TICKET" {
				now := types.Timestamp(time.Now())
				vouch.Product.Image = types.Image{
					FileID:       bson.ObjectIdHex("5849540db80dff3e46d8e7ab"),
					MimeType:     "IMAGE_PNG",
					CreationDate: &now,
				}
			}
			vouchers = append(vouchers, vouch)
		}
	}
	// if viper.GetBool("verbose") {
	j, _ := json.MarshalIndent(vouchers, "", "    ")
	log.Println("before inserting")
	log.Println(string(j))
	// }
	for _, voucher := range vouchers {
		err := vouchersCollection.Collection.Insert(voucher)
		if err != nil {
			return vouchers, err
		}
	}

	return vouchers, nil
}
