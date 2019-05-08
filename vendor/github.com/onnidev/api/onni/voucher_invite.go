package onni

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// CreateInviteVoucher TODO: NEEDS COMMENT INFO
func CreateInviteVoucher(
	ctx context.Context,
	email string,
	club types.Club,
	party types.Party,
	req types.VoucherPostRequest,
	partyProduct types.PartyProduct,
	id, name, status string,
) (types.Voucher, error) {
	userClub := ctx.Value(middlewares.UserClubKey).(types.UserClub)
	voucher := types.Voucher{}
	vouchersCollection, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		return voucher, err
	}
	horario := types.Timestamp(time.Now())
	zero := types.Timestamp(time.Time{})
	vouchers := []types.Voucher{}
	for i := 1; i <= req.Quantity; i++ {
		voucher = types.Voucher{
			ID:                    bson.NewObjectId(),
			CreationDate:          &horario,
			CustomerID:            bson.ObjectIdHex(id),
			PartyID:               bson.ObjectIdHex(req.PartyID),
			ClubID:                club.ID,
			PartyProductID:        bson.ObjectIdHex(req.PartyProductID),
			Type:                  req.Type,
			ClubName:              club.Name,
			PartyName:             party.Name,
			CustomerName:          name,
			StartDate:             party.StartDate,
			VoucherUseDate:        &zero,
			EndDate:               party.EndDate,
			Status:                status,
			ResponsableUserClubID: &userClub.ID,
			Price: partyProduct.MoneyAmount,
			Product: types.VoucherProduct{
				// Image: partyProduct.Image,
				Type: partyProduct.Type,
				Name: partyProduct.Name,
			},
		}
		if partyProduct.Image != nil {
			voucher.Product.Image = *partyProduct.Image
		}
		if partyProduct.Type == "TICKET" {
			now := types.Timestamp(time.Now())
			voucher.Product.Image = types.Image{
				FileID:       bson.ObjectIdHex("5849540db80dff3e46d8e7ab"),
				MimeType:     "IMAGE_PNG",
				CreationDate: &now,
			}
		}
		vouchers = append(vouchers, voucher)
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(voucher, "", "    ")
		log.Println("before inserting")
		log.Println(string(j))
	}
	for _, voucher := range vouchers {
		err := vouchersCollection.Collection.Insert(voucher)
		if err != nil {
			return voucher, err
		}
	}
	go MailNewVoucher(email, party, partyProduct, req, id, true)
	return voucher, nil
}
