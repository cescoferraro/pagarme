package types

import "gopkg.in/mgo.v2/bson"

// PromotionPatchRequest TODO: NEEDS COMMENT INFO
type PromotionPatchRequest struct {
	Name                   string   `json:"name" bson:"name"`
	PartyProductID         string   `json:"partyProductId" bson:"partyProductId"`
	MakePublic             bool     `json:"makePublic" bson:"makePublic"`
	IgnoreActiveBatchRules bool     `json:"ignoreActiveBatchRules" bson:"ignoreActiveBatchRules"`
	AvailableToFollowers   bool     `json:"availableToFollowers" bson:"availableToFollowers"`
	Promoters              []string `json:"promoters" bson:"promoters"`

	StartDate           *Timestamp `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate             *Timestamp `bson:"endDate,omitempty" json:"endDate,omitempty"`
	QuantityTotal       *int64     `json:"quantityTotal,omitempty" bson:"quantityTotal,omitempty"`
	QuantityPerCustomer *int64     `json:"quantityPerCustomer,omitempty" bson:"quantityPerCustomer,omitempty"`
}

// PromotionPostRequest TODO: NEEDS COMMENT INFO
type PromotionPostRequest struct {
	Name                   string   `json:"name" bson:"name"`
	PartyProductID         string   `json:"partyProductId" bson:"partyProductId"`
	Price                  float64  `json:"price" bson:"price"`
	MakePublic             bool     `json:"makePublic" bson:"makePublic"`
	IgnoreActiveBatchRules bool     `json:"ignoreActiveBatchRules" bson:"ignoreActiveBatchRules"`
	AvailableToFollowers   bool     `json:"availableToFollowers" bson:"availableToFollowers"`
	Promoters              []string `json:"promoters" bson:"promoters"`

	StartDate           *Timestamp `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate             *Timestamp `bson:"endDate,omitempty" json:"endDate,omitempty"`
	QuantityTotal       *int64     `json:"quantityTotal,omitempty" bson:"quantityTotal,omitempty"`
	QuantityPerCustomer *int64     `json:"quantityPerCustomer,omitempty" bson:"quantityPerCustomer,omitempty"`
}

// Promotion TODO: NEEDS COMMENT INFO
type Promotion struct {
	ID                     bson.ObjectId `json:"id" bson:"id"`
	Name                   string        `json:"name" bson:"name"`
	StartDate              *Timestamp    `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate                *Timestamp    `bson:"endDate,omitempty" json:"endDate,omitempty"`
	Price                  Price         `json:"price" bson:"price"`
	QuantityTotal          *int64        `json:"quantityTotal" bson:"quantityTotal"`
	QuantityPurchased      int64         `json:"quantityPurchased" bson:"quantityPurchased"`
	QuantityPerCustomer    *int64        `json:"quantityPerCustomer" bson:"quantityPerCustomer"`
	MakePublic             bool          `json:"makePublic" bson:"makePublic"`
	IgnoreActiveBatchRules bool          `json:"ignoreActiveBatchRules" bson:"ignoreActiveBatchRules"`
	AvailableToFollowers   bool          `json:"availableToFollowers" bson:"availableToFollowers"`
	Promoters              *[]Promoter   `json:"promoters" bson:"promoters"`
}

// Valid sdkfjndsf
func (promo Promotion) Valid() bool {
	if bson.IsObjectIdHex(promo.ID.Hex()) {
		return true
	}
	return false
}

// SoftPromotion TODO: NEEDS COMMENT INFO
type SoftPromotion struct {
	ID                     bson.ObjectId   `json:"id" bson:"id"`
	Name                   string          `json:"name" bson:"name"`
	StartDate              *Timestamp      `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate                *Timestamp      `bson:"endDate,omitempty" json:"endDate,omitempty"`
	PartyName              string          `json:"partyName" bson:"partyName"`
	ProductName            string          `json:"productName" bson:"productName"`
	Price                  Price           `json:"price" bson:"price"`
	QuantityTotal          *int64          `json:"quantityTotal" bson:"quantityTotal"`
	QuantityPurchased      int64           `json:"quantityPurchased" bson:"quantityPurchased"`
	QuantityPerCustomer    *int64          `json:"quantityPerCustomer" bson:"quantityPerCustomer"`
	MakePublic             bool            `json:"makePublic" bson:"makePublic"`
	IgnoreActiveBatchRules bool            `json:"ignoreActiveBatchRules" bson:"ignoreActiveBatchRules"`
	AvailableToFollowers   bool            `json:"availableToFollowers" bson:"availableToFollowers"`
	Promoters              *[]SoftPromoter `json:"promoters" bson:"promoters"`
}
