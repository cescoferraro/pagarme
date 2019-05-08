package onni

import (
	"context"
	"errors"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
)

// ValidatePromotionBuy TODO: NEEDS COMMENT INFO
func ValidatePromotionBuy(ctx context.Context, customer types.Customer, req types.BuyPartyProductsItemRequest, clubID bson.ObjectId, ids []bson.ObjectId) error {
	_, promotion, err := Promotion(ctx, req.PromotionID)
	if err != nil {
		return err
	}
	err = ValidatePromotion(ctx, customer, promotion)
	if err != nil {
		return err
	}
	isCustomer := shared.ContainsObjectID(ids, promotion.ID)
	isFollower := shared.ContainsObjectID(customer.FavoriteClubs, clubID)
	log.Println("isCustomer", isCustomer, "follower", promotion.AvailableToFollowers && isFollower, "public", promotion.MakePublic)
	if isCustomer || promotion.MakePublic || promotion.AvailableToFollowers && isFollower {
		if promotion.IgnoreActiveBatchRules {
			return nil
		}
		partyP, err := PartyProduct(ctx, req.PartyProductID)
		if err != nil {
			return err
		}
		_, err = ValidatePartyProduct(partyP)
		if err != nil {
			log.Println("not valid over PartyProduct")
			return err
		}
		return nil
	}
	err = errors.New("buggued")
	return err
}

// ValidatePromotion TODO: NEEDS COMMENT INFO
func ValidatePromotion(ctx context.Context, customer types.Customer, promotion types.Promotion) error {
	if promotion.QuantityTotal != nil {
		if *promotion.QuantityTotal-promotion.QuantityPurchased < 1 {
			err := errors.New("sold out")
			return err
		}
	}
	if promotion.QuantityPerCustomer != nil {
		repo, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
		if !ok {
			err := errors.New("bug")
			return err
		}
		quantyPurchased, err := repo.PromotionCountPurchasedbyCustomer(promotion.ID.Hex(), customer.ID.Hex())
		if err != nil {
			err := errors.New("bug")
			return err
		}
		if quantyPurchased >= int(*promotion.QuantityPerCustomer) {
			err := errors.New("bought too much")
			return err
		}
	}
	if !promotion.StartDate.Time().IsZero() && !promotion.EndDate.Time().IsZero() {
		onTime := shared.
			InTimeSpan(promotion.StartDate.Time(), promotion.EndDate.Time(), time.Now())
		if !onTime {
			err := errors.New("not on time")
			return err
		}
	}
	return nil
}

// ValidatePartyProductBuy TODO: NEEDS COMMENT INFO
func ValidatePartyProductBuy(ctx context.Context, id string) (types.PartyProduct, error) {
	log.Printf("##### validating party product %v \n", id)
	partyP, err := PartyProduct(ctx, id)
	if err != nil {
		return types.PartyProduct{}, err
	}
	log.Printf("##### validating %v \n", partyP.Name)
	partyP, err = ValidatePartyProduct(partyP)
	if err != nil {
		return types.PartyProduct{}, err
	}
	return partyP, nil
}

// ValidatePartyProduct TODO: NEEDS COMMENT INFO
func ValidatePartyProduct(partyP types.PartyProduct) (types.PartyProduct, error) {
	err := errors.New("partyPropduct not valid")
	result := types.PartyProduct{}
	if partyP.Status == "ACTIVE" {
		if !partyP.Deprecated {
			if partyP.DoNotSellMoreThan-(partyP.QuantityPurchased+partyP.QuantityFree) > 0 {
				if !partyP.DoNotSellBefore.Time().IsZero() && !partyP.DoNotSellAfter.Time().IsZero() {
					if shared.InTimeSpan(partyP.DoNotSellBefore.Time(), partyP.DoNotSellAfter.Time(), time.Now()) {
						return partyP, nil
					}
					return result, err
				}
				if !partyP.DoNotSellBefore.Time().IsZero() {
					if time.Now().Before(partyP.DoNotSellBefore.Time()) {
						return result, err
					}
					return partyP, nil
				}
				if !partyP.DoNotSellAfter.Time().IsZero() {
					if time.Now().After(partyP.DoNotSellAfter.Time()) {
						return result, err
					}
					return partyP, nil
				}
				return partyP, nil
			}
			return result, err
		}
	}
	return result, err
}
