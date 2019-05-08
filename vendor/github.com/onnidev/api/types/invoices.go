package types

import (
	"gopkg.in/mgo.v2/bson"
)

// Invoice sdkjfn
type Invoice struct {
	ID             bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	CreationDate   *Timestamp     `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate     *Timestamp     `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	CustomerID     bson.ObjectId  `json:"customerId" bson:"customerId,omitempty"`
	PartyID        bson.ObjectId  `json:"partyId" bson:"partyId,omitempty"`
	ClubID         bson.ObjectId  `json:"clubId" bson:"clubId,omitempty"`
	Status         string         `bson:"status" json:"status"`
	OperationType  string         `json:"operationType" bson:"operationType"`
	Itens          []InvoiceItem  `json:"itens" bson:"itens"`
	TransactionID  *string        `json:"transactionId" bson:"transactionId"`
	ValueToOnni    float64        `json:"valueToOnni" bson:"valueToOnni"`
	ValueToClub    float64        `json:"valueToClub" bson:"valueToClub"`
	PercentToONNi  float64        `json:"percentToOnni" bson:"percentToOnni"`
	PercentToClub  float64        `json:"percentToClub" bson:"percentToClub"`
	Total          Price          `json:"total" bson:"total"`
	Log            Log            `json:"log" bson:"log"`
	InviteVendorID *bson.ObjectId `json:"inviteVendorId" bson:"inviteVendorId,omitempty"`
}

// Valid TODO: NEEDS COMMENT INFO
func (invoice Invoice) Valid() bool {
	if bson.IsObjectIdHex(invoice.ID.Hex()) {
		return true
	}
	return false
}

// Invoices TODO: NEEDS COMMENT INFO
type Invoices struct {
	Drinks  *Invoice `json:"drinks" bson:"drinks"`
	Tickets *Invoice `json:"tickets" bson:"tickets"`
}

// All TODO: NEEDS COMMENT INFO
func (invoices Invoices) All() []Invoice {
	result := []Invoice{}
	if invoices.Drinks != nil {
		invoice := *invoices.Drinks
		if bson.IsObjectIdHex(invoice.ID.Hex()) {
			result = append(result, *invoices.Drinks)
		}
	}
	if invoices.Tickets != nil {
		invoice := *invoices.Tickets
		if bson.IsObjectIdHex(invoice.ID.Hex()) {
			result = append(result, *invoices.Tickets)
		}
	}
	return result
}

// InvoiceItem TODO: NEEDS COMMENT INFO
type InvoiceItem struct {
	PartyProductID bson.ObjectId      `json:"partyProductId" bson:"partyProductId"`
	PromotionID    *bson.ObjectId     `json:"promotionId" bson:"promotionId"`
	UnitPrice      Price              `json:"unitPrice" bson:"unitPrice"`
	Quantity       int64              `json:"quantity" bson:"quantity"`
	Product        InvoiceItemProduct `json:"product" bson:"product"`
}

// InvoiceItemProduct TODO: NEEDS COMMENT INFO
type InvoiceItemProduct struct {
	Name  string `json:"name" bson:"name"`
	Type  string `json:"type" bson:"type"`
	Image *Image `json:"image" bson:"image"`
}

// Log dskfjn
type Log struct {
	AppVersion string  `json:"appVersion" bson:"appVersion"`
	DeviceID   string  `json:"deviceId" bson:"deviceId"`
	DeviceSO   string  `json:"deviceSO" bson:"deviceSO"`
	Latitude   float64 `json:"latitude" bson:"latitude"`
	Longitude  float64 `json:"longitude" bson:"longitude"`
}
