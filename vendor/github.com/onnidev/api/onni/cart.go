package onni

import (
	"context"
	"errors"
	"log"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// GetFinalPromotions TODO: NEEDS COMMENT INFO
func GetFinalPromotions(
	ctx context.Context,
	partyProducts []types.PartyProduct,
	party types.Party,
	club types.Club,
	customer types.Customer,
) ([]types.CartPartyProduct, error) {
	result := []types.CartPartyProduct{}
	ids, err := GetCustomerPartyPromotions(ctx, party, customer)
	if err != nil {
		return result, err
	}
	voucherRepo, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("facebook assert errro")
		return result, err
	}
	for _, partyP := range partyProducts {
		if partyP.Status != "INACTIVE" {
			if partyP.PromotionalPrices != nil {
				log.Println("================================================================== ")
				log.Println("============= PartyProduct " + partyP.ID.Hex() + " " + partyP.Name + "================ ")
				log.Println("================================================================== ")
				log.Println("got promotions", len(*partyP.PromotionalPrices))
				for _, promotion := range *partyP.PromotionalPrices {
					log.Println("********** Promotion " + promotion.ID.Hex() + "****************")
					err := ValidatePromotion(ctx, customer, promotion)
					if err != nil {
						log.Println("erro de validaçao na promo:", err.Error())
						continue
					}
					vouchers, err := voucherRepo.PromotionPurchasedbyCustomer(promotion.ID.Hex(), customer.ID.Hex())
					if err != nil {
						continue
					}
					log.Println("eu comporei", len(vouchers))
					log.Println("passed promotion validation over promotionId")
					isCustomer := shared.ContainsObjectID(ids, promotion.ID)
					isFollower := shared.ContainsObjectID(customer.FavoriteClubs, club.ID)
					log.Println("isCustomer", isCustomer, "follower", promotion.AvailableToFollowers && isFollower, "public", promotion.MakePublic)
					if isCustomer || promotion.MakePublic || promotion.AvailableToFollowers && isFollower {
						log.Println("its a valid kind of promotion")
						exists, index := shared.ContainsObjectIDIndex(types.IDS(result), partyP.ID)
						log.Println("exists", exists, "index", index)
						if promotion.IgnoreActiveBatchRules {
							log.Println("ignore batch rules")
							if exists {
								isSmaller := result[index].Price.Value > promotion.Price.Value
								log.Println("existing", result[index].Price.Value)
								log.Println("next", promotion.Price.Value)
								log.Println("isSmaller", isSmaller)
								if isSmaller {
									newPrice := promotion.Price.Value
									fee := partyP.GetFee(true, party, club, promotion)
									result[index].Price.Value = newPrice
									result[index].PromotionID = &promotion.ID
									result[index].Price.Fee = fee
									result[index].Price.TotalValue = newPrice + fee
								}
								continue
							}
							log.Println("GOT IT")
							result = append(result, partyP.Discount(party, club, promotion, len(vouchers)))
							continue
						}
						log.Println("does not ignore batch rules")
						_, err := ValidatePartyProduct(partyP)
						if err != nil {
							log.Println("not valid over PartyProduct")
							log.Println(err.Error())
							continue
						}
						log.Println("validate over PartyProduct")
						if exists {
							isSmaller := result[index].Price.Value > promotion.Price.Value
							log.Println("existing", result[index].Price.Value)
							log.Println("next", promotion.Price.Value)
							log.Println("isSmaller", isSmaller)
							if isSmaller {
								newPrice := promotion.Price.Value
								fee := partyP.GetFee(true, party, club, promotion)
								result[index].Price.Value = newPrice
								result[index].PromotionID = &promotion.ID
								result[index].Price.Fee = fee
								result[index].Price.TotalValue = newPrice + fee
							}
							continue
						}
						log.Println("GOT IT")
						result = append(result, partyP.Discount(party, club, promotion, len(vouchers)))
						continue
					}
					log.Println("**** END" + partyP.ID.Hex() + "*****")
				}
			}
		}

	}
	return result, nil
}

// GetCustomerPartyPromotions TODO: NEEDS COMMENT INFO
func GetCustomerPartyPromotions(ctx context.Context, party types.Party, customer types.Customer) ([]bson.ObjectId, error) {
	var result []bson.ObjectId
	promotions, err := GetPartyPromotions(ctx, party.ID.Hex())
	if err != nil {
		return result, err
	}
	repo, ok := ctx.
		Value(middlewares.PromotionalCustomerRepoKey).(interfaces.PromotionalCustomerRepo)
	if !ok {
		err := errors.New("bug")
		return result, err
	}
	return repo.GetAllFromCustomer(promotions, customer.ID.Hex())
}

// GetPartyPromotions TODO: NEEDS COMMENT INFO
func GetPartyPromotions(ctx context.Context, partyID string) ([]bson.ObjectId, error) {
	var result []bson.ObjectId
	products, err := PartyPartyProducts(ctx, partyID)
	if err != nil {
		return result, err
	}
	for _, partyP := range products {
		if partyP.PromotionalPrices != nil {
			// byt, _ := json.MarshalIndent(*partyP.PromotionalPrices, "", "    ")
			// log.Println(string(byt))
			for _, promotion := range *partyP.PromotionalPrices {
				result = append(result, promotion.ID)

			}
		}
	}
	return result, nil
}

// GroupedPartyProducts TODO: NEEDS COMMENT INFO
func GroupedPartyProducts(ctx context.Context, club types.Club, party types.Party) (map[string][]types.CartPartyProduct, error) {
	result := make(map[string][]types.CartPartyProduct)
	partyProductCollection := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	partyProducts, err := partyProductCollection.GetByPartyID(party.ID.Hex())
	if err != nil {
		return result, err
	}
	availablePartyProducts := []types.PartyProduct{}
	for _, partyP := range partyProducts {
		validatedPartyP, err := ValidatePartyProduct(partyP)
		if err != nil {
			continue
		}
		availablePartyProducts = append(availablePartyProducts, validatedPartyP)
	}
	for _, product := range availablePartyProducts {
		log.Println("**********", product.Name)
		if product.Type == "TICKET" {
			log.Println("tIcKET")
			log.Println(product.Batches)
			log.Println(product.OwnerBatchID)
			if len(product.Batches) != 0 || product.OwnerBatchID != nil {
				continue
			}
			next := product.NoDiscount(party, club)
			result["_empty"] = append(result["_empty"], next)
		}
		if product.Category != nil {
			result[*product.Category] = append(result[*product.Category], product.NoDiscount(party, club))
		}
		if product.Featured {
			result["_featured"] = append(result["_featured"], product.NoDiscount(party, club))
		}
	}
	for _, batch := range Masters(partyProducts) {
		log.Println("99999999")
		log.Println("99999999")
		log.Println(mySlaves(partyProducts, batch.ID.Hex()))
		best, err := bestSlave(partyProducts, batch.ID.Hex())
		if err != nil {
			continue
		}
		log.Println("nas que era")
		log.Println(best)
		next := best.NoDiscount(party, club)
		result["_empty"] = append(result["_empty"], next)
		log.Println("99999999")
		log.Println("99999999")
		log.Println("99999999")
	}
	return result, nil
}

func mySlaves(batches []types.PartyProduct, id string) []types.PartyProduct {
	result := []types.PartyProduct{}
	for _, batch := range batches {
		if batch.OwnerBatchID != nil {
			if batch.OwnerBatchID.Hex() == id {
				result = append(result, batch)
			}
		}
	}
	return result
}
func bestSlave(batches []types.PartyProduct, id string) (types.PartyProduct, error) {
	result := types.PartyProduct{}
	err := errors.New("no best slave")
	for _, batch := range batches {
		if batch.OwnerBatchID != nil {
			if batch.OwnerBatchID.Hex() == id {
				validated, err := ValidatePartyProduct(batch)
				if err != nil {
					log.Println("9898989")
					log.Println(err.Error())
					continue
				}
				log.Println("venceu na vida")
				return validated, nil
				// err = nil
				// result = validated
			}
		}
	}
	return result, err
}

// Masters TODO: NEEDS COMMENT INFO
func Masters(all []types.PartyProduct) []types.PartyProduct {
	result := []types.PartyProduct{}
	for _, product := range all {
		if len(product.Batches) != 0 {
			result = append(result, product)
		}
	}
	return result
}
