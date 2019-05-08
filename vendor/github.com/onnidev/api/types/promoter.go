package types

import "gopkg.in/mgo.v2/bson"

// Promoter TODO: NEEDS COMMENT INFO
type Promoter struct {
	PromoterID bson.ObjectId `json:"promoterId" bson:"promoterId"`
	Deleted    bool          `json:"deleted" bson:"deleted"`
}

// SoftPromoter TODO: NEEDS COMMENT INFO
type SoftPromoter struct {
	PromoterID   string `json:"promoterId" bson:"promoterId"`
	PromoterName string `json:"promoterName" bson:"promoterName"`
	Selected     bool   `json:"selected" bson:"selected"`
}
