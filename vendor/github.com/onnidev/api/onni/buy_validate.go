package onni

import (
	"context"
	"errors"

	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// ValidateBuyPostRequest TODO: NEEDS COMMENT INFO
func ValidateBuyPostRequest(ctx context.Context, buy types.BuyPost, club types.Club, party types.Party, customer types.Customer) error {
	if len(buy.Products) != 0 {
		for _, iten := range buy.Products {
			if !bson.IsObjectIdHex(iten.PartyProductID) {
				return errors.New("you need a valid partyP")
			}
		}
		ids, err := GetCustomerPartyPromotions(ctx, party, customer)
		if err != nil {
			return err
		}
		for _, iten := range buy.Products {
			if bson.IsObjectIdHex(iten.PromotionID) {
				err := ValidatePromotionBuy(ctx, customer, iten, club.ID, ids)
				if err != nil {
					return err
				}
				continue
			}
			if !bson.IsObjectIdHex(iten.PromotionID) {
				_, err := ValidatePartyProductBuy(ctx, iten.PartyProductID)
				if err != nil {
					return err
				}
				continue
			}
			err := errors.New("buggued")
			return err
		}
		return nil
	}
	return errors.New("you need to want to buy something")
}
