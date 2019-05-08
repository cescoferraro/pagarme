package types

import (
	"gopkg.in/mgo.v2/bson"
)

// Product fknjsd
type Product struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	Name         string        `json:"name" bson:"name"`
	Deprecated   bool          `json:"deprecated" bson:"deprecated"`
	NameSort     string        `json:"nameSort" bson:"nameSort"`
	Type         string        `json:"type" bson:"type"`
	Image        Image         `json:"image" bson:"image"`
	Category     string        `json:"category" bson:"category"`
}

// VoucherProduct fknjsd
type VoucherProduct struct {
	Name  string `json:"name" bson:"name"`
	Type  string `json:"type" bson:"type"`
	Image Image  `json:"image" bson:"image"`
}

// ProductItem type for the above middleware
type ProductItem struct {
	Quantity       int64         `json:"quantity" bson:"quantity"`
	PartyProductID bson.ObjectId `json:"partyProductId" bson:"partyProductId,omitempty"`
	PromotionID    bson.ObjectId `json:"promotionId" bson:"promotionId,omitempty"`
	Product        Product       `json:"product" bson:"product"`
	UnitPrice      Price         `json:"total" bson:"total"`
}

// ProductPostRequest type for the above middleware
type ProductPostRequest struct {
	ID       string `json:"id" bson:"id"`
	Name     string `bson:"name" json:"name"`
	Type     string `bson:"type" json:"type"`
	Category string `bson:"category" json:"category"`
	Image    string `json:"image" bson:"image"`
}

// ProductSoftPatchRequest type for the above middleware
type ProductSoftPatchRequest struct {
	Type     string `bson:"type" json:"type"`
	Name     string `bson:"name" json:"name"`
	Category string `bson:"category" json:"category"`
}

// ProductPatchRequest type for the above middleware
type ProductPatchRequest struct {
	Image    string `bson:"image" json:"image"`
	Type     string `bson:"type" json:"type"`
	Name     string `bson:"name" json:"name"`
	Category string `bson:"category" json:"category"`
}
