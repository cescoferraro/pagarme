package types

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

// CartPartyProduct fknjsd
type CartPartyProduct struct {
	ID                bson.ObjectId      `json:"id" bson:"_id,omitempty"`
	ProductID         *bson.ObjectId     `json:"productId" bson:"productId"`
	PromotionID       *bson.ObjectId     `json:"promotionId" bson:"promotionId"`
	Name              string             `json:"name" bson:"name"`
	Type              string             `json:"type" bson:"type"`
	QuantityAvailable int64              `json:"quantityAvailable" bson:"quantityAvailable"`
	Image             Image              `bson:"image" json:"image"`
	Price             CartPromotionPrice `json:"price" bson:"price"`
}

// IDS TODO: NEEDS COMMENT INFO
func IDS(products []CartPartyProduct) []bson.ObjectId {
	var result []bson.ObjectId
	for _, product := range products {
		result = append(result, product.ID)
	}
	return result

}

// Discount TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) Discount(party Party, club Club, promotion Promotion, count int) CartPartyProduct {
	fee := partyP.GetFee(true, party, club, promotion)
	var productID bson.ObjectId
	if partyP.Product != nil {
		productID = partyP.Product.ID
	}
	log.Println("promotion")
	log.Println("comprados pelo customer", count)
	if promotion.QuantityTotal != nil {
		log.Println("promotion.QuantityTotal", *promotion.QuantityTotal)
	}
	log.Println("promotion.QuantityPurchased", promotion.QuantityPurchased)
	if promotion.QuantityPerCustomer != nil {
		log.Println("promotion.QuantityPerCustomer", *promotion.QuantityPerCustomer)
	}
	log.Println("promotion.IgnoreActiveBatchRules", promotion.IgnoreActiveBatchRules)
	log.Println("partyP.DoNotSellMoreThan", partyP.DoNotSellMoreThan)
	log.Println("partyP.QuantityPurchased", partyP.QuantityPurchased)
	log.Println("partyP.QuantityFree", partyP.QuantityFree)

	available := partyP.DoNotSellMoreThan - partyP.QuantityPurchased - partyP.QuantityFree
	if promotion.IgnoreActiveBatchRules {
		available = 9999999999
	}
	if promotion.QuantityTotal != nil {
		log.Println("tem toptal")
		available = *promotion.QuantityTotal - promotion.QuantityPurchased
	}
	if promotion.QuantityPerCustomer != nil {
		log.Println("tem toptal per customer")
		available = *promotion.QuantityPerCustomer - int64(count)
	}
	if promotion.QuantityTotal != nil && promotion.QuantityPerCustomer != nil {
		log.Println("tem as duas")
		totalA := *promotion.QuantityTotal - promotion.QuantityPurchased
		custA := *promotion.QuantityPerCustomer - int64(count)
		available = totalA
		if custA < totalA {
			available = custA
		}
	}
	product := CartPartyProduct{
		ID:                partyP.ID,
		ProductID:         &productID,
		PromotionID:       &promotion.ID,
		Name:              partyP.Name,
		Type:              partyP.Type,
		QuantityAvailable: available,
		Price: CartPromotionPrice{
			Value:                promotion.Price.Value,
			Currency:             promotion.Price.CurrentIsoCode,
			Fee:                  fee,
			TotalValue:           fee + promotion.Price.Value,
			ValueBeforePromotion: &partyP.MoneyAmount.Value,
		},
	}
	if partyP.Image != nil {
		product.Image = *partyP.Image
	}
	return product
}

// NoDiscount TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) NoDiscount(party Party, club Club) CartPartyProduct {
	fee := partyP.GetFee(false, party, club, Promotion{})
	var productID bson.ObjectId
	if partyP.Product != nil {
		productID = partyP.Product.ID
	}
	product := CartPartyProduct{
		ID:                partyP.ID,
		ProductID:         &productID,
		Name:              partyP.Name,
		Type:              partyP.Type,
		QuantityAvailable: partyP.DoNotSellMoreThan - partyP.QuantityPurchased - partyP.QuantityFree,
		Price: CartPromotionPrice{
			Value:      partyP.MoneyAmount.Value,
			Currency:   partyP.MoneyAmount.CurrentIsoCode,
			TotalValue: fee + partyP.MoneyAmount.Value,
			Fee:        fee,
		},
	}
	if partyP.Image != nil {
		product.Image = *partyP.Image
	}
	return product
}

// CartPromotionPrice TODO: NEEDS COMMENT INFO
type CartPromotionPrice struct {
	Value                float64  `json:"value" bson:"value"`
	Currency             string   `json:"currency" bson:"currency"`
	Fee                  float64  `json:"fee" bson:"fee"`
	TotalValue           float64  `json:"totalValue" bson:"totalValue"`
	ValueBeforePromotion *float64 `json:"valueBeforePromotion" bson:"valueBeforePromotion"`
}
