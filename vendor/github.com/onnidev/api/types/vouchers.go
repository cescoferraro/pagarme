package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Voucher is a mongo document
type Voucher struct {
	ID           bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreationDate *Timestamp    `json:"creationDate,omitempty" bson:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `json:"updateDate,omitempty" bson:"updateDate,omitempty"`
	StartDate    *Timestamp    `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate      *Timestamp    `json:"endDate,omitempty" bson:"endDate,omitempty"`

	CustomerID     bson.ObjectId  `json:"customerId" bson:"customerId"`
	PartyID        bson.ObjectId  `json:"partyId" bson:"partyId"`
	ClubID         bson.ObjectId  `json:"clubId" bson:"clubId"`
	PartyProductID bson.ObjectId  `json:"partyProductId" bson:"partyProductId"`
	InvoiceID      *bson.ObjectId `json:"invoiceId,omitempty" bson:"invoiceId,omitempty"`
	PromotionID    *bson.ObjectId `json:"promotionId,omitempty" bson:"promotionId,omitempty"`
	ClubName       string         `json:"clubName" bson:"clubName"`
	PartyName      string         `json:"partyName" bson:"partyName"`
	CustomerName   string         `json:"customerName" bson:"customerName"`

	Status        string         `json:"status" bson:"status"`
	Price         Price          `json:"price" bson:"price"`
	Product       VoucherProduct `json:"product" bson:"product"`
	Type          string         `json:"type" bson:"type"`
	TransactionID *string        `json:"transactionId,omitempty" bson:"transactionId,omitempty"`

	VoucherUseDate         *Timestamp     `json:"voucherUseDate,omitempty" bson:"voucherUseDate,omitempty"`
	VoucherUseUserClubID   *bson.ObjectId `json:"voucherUseUserClubId,omitempty" bson:"voucherUseUserClubId,omitempty"`
	VoucherUseUserClubName *string        `json:"voucherUseUserClubName,omitempty" bson:"voucherUseUserClubName,omitempty"`
	ResponsableUserClubID  *bson.ObjectId `json:"responsibleUserClubId,omitempty" bson:"responsibleUserClubId,omitempty"`
	TransferedFrom         *bson.ObjectId `json:"transferedFrom,omitempty" bson:"transferedFrom,omitempty"`
}

// BuyPartyProductsItemRequest TODO: NEEDS COMMENT INFO
func (voucher Voucher) BuyPartyProductsItemRequest() []BuyPartyProductsItemRequest {
	result := BuyPartyProductsItemRequest{
		PartyProductID: voucher.PartyProductID.Hex(),
		Quantity:       1,
	}
	if voucher.PromotionID != nil {
		id := *voucher.PromotionID
		result.PromotionID = id.Hex()
	}
	return []BuyPartyProductsItemRequest{result}
}

// VoucherPostRequest type for the above middleware
type VoucherPostRequest struct {
	Type           string   `bson:"type" json:"type"`
	Quantity       int      `bson:"quantity" json:"quantity"`
	PartyID        string   `json:"partyId" bson:"partyId"`
	PartyProductID string   `json:"partyProductId" bson:"partyProductId"`
	Emails         []string `bson:"mails" json:"mails"`
}

// VoucherUseConstrain sjkdfn
type VoucherUseConstrain struct {
	Drink  bool `bson:"drink" json:"drink"`
	Ticket bool `json:"ticket" bson:"ticket"`
}

// VoucherSoftValidateReq sjkdfn
type VoucherSoftValidateReq struct {
	VoucherID string `json:"voucherId" bson:"voucherId"`
	ClubID    string `json:"clubId" bson:"clubId"`
}

// VoucherSoftReadReq sjkdfn
type VoucherSoftReadReq struct {
	VoucherID  string `json:"voucherId" bson:"voucherId"`
	ClubID     string `json:"clubId" bson:"clubId"`
	UserClubID string `json:"userClubId" bson:"userClubId"`
}

// ToBeTransfed Insert the user on a database
func (c *CompleteVoucher) ToBeTransfed(customer Customer) CompleteVoucher {
	horario := Timestamp(time.Now())
	return CompleteVoucher{
		ID:                     bson.NewObjectId(),
		CreationDate:           &horario,
		UpdateDate:             &horario,
		CustomerID:             customer.ID,
		CustomerName:           customer.FirstName + " " + customer.LastName,
		Status:                 c.Status,
		PartyID:                c.PartyID,
		Customer:               c.Customer,
		ClubID:                 c.ClubID,
		Responsable:            c.Responsable,
		Price:                  c.Price,
		InvoiceID:              c.InvoiceID,
		PartyProductID:         c.PartyProductID,
		Product:                c.Product,
		ClubName:               c.ClubName,
		PartyName:              c.PartyName,
		StartDate:              c.StartDate,
		EndDate:                c.EndDate,
		VoucherUseDate:         c.VoucherUseDate,
		VoucherUseUserClubID:   c.VoucherUseUserClubID,
		ResponsableUserClubID:  c.ResponsableUserClubID,
		VoucherUseUserClubName: c.VoucherUseUserClubName,
		Type:           "TRANSFERED",
		TransferedFrom: c.ID,
	}
}

// ToBeTransfed Insert the user on a database
func (c *Voucher) ToBeTransfed(customer Customer) Voucher {
	horario := Timestamp(time.Now())
	return Voucher{
		ID:           bson.NewObjectId(),
		CreationDate: &horario,
		UpdateDate:   &horario,
		CustomerID:   customer.ID,
		CustomerName: customer.FirstName + " " + customer.LastName,
		Status:       c.Status,
		PartyID:      c.PartyID,
		ClubID:       c.ClubID,
		Price:        c.Price,
		InvoiceID:    c.InvoiceID,

		PartyProductID:         c.PartyProductID,
		Product:                c.Product,
		ClubName:               c.ClubName,
		PartyName:              c.PartyName,
		StartDate:              c.StartDate,
		EndDate:                c.EndDate,
		VoucherUseDate:         c.VoucherUseDate,
		VoucherUseUserClubID:   c.VoucherUseUserClubID,
		ResponsableUserClubID:  c.ResponsableUserClubID,
		VoucherUseUserClubName: c.VoucherUseUserClubName,
		Type:           "TRANSFERED",
		TransferedFrom: &c.ID,
	}
}

// AppCompleteVoucher is a mongo document
type AppCompleteVoucher struct {
	ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreationDate   *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate     *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	StartDate      *Timestamp    `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate        *Timestamp    `bson:"endDate,omitempty" json:"endDate,omitempty"`
	CustomerID     bson.ObjectId `json:"customerId" bson:"customerId"`
	PartyID        bson.ObjectId `json:"partyId" bson:"partyId"`
	ClubID         bson.ObjectId `json:"clubId" bson:"clubId"`
	InvoiceID      bson.ObjectId `json:"invoiceId" bson:"invoiceId"`
	PartyProductID bson.ObjectId `json:"partyProductId" bson:"partyProductId"`
	ClubName       string        `json:"clubName" bson:"clubName"`
	PartyName      string        `json:"partyName" bson:"partyName"`
	CustomerName   string        `json:"customerName" bson:"customerName"`

	VoucherUseDate         *Timestamp     `bson:"voucherUseDate,omitempty" json:"voucherUseDate,omitempty"`
	VoucherUseUserClubID   *bson.ObjectId `json:"voucherUseUserClubId" bson:"voucherUseUserClubId"`
	ResponsableUserClubID  *bson.ObjectId `json:"responsibleUserClubId" bson:"responsibleUserClubId"`
	VoucherUseUserClubName string         `json:"voucherUseUserClubName" bson:"voucherUseUserClubName"`

	Status  string  `json:"status" bson:"status"`
	Price   Price   `json:"price" bson:"price"`
	Product Product `json:"product" bson:"product"`
	Type    string  `json:"type" bson:"type"`

	Customer     *Customer     `json:"customer" bson:"customer"`
	Responsable  *UserClub     `json:"responsable" bson:"responsable"`
	PartyProduct *PartyProduct `json:"partyProduct,omitempty" bson:"partyProduct,omitempty" `
	Party        *AppParty     `json:"party,omitempty" bson:"party,omitempty" `
	Club         *AppClub      `json:"club,omitempty" bson:"club,omitempty" `

	TransferedFrom bson.ObjectId `json:"transferedFrom" bson:"transferedFrom"`
}

// creationDate : 1529808466924
// currency : "BRL"
// customerId : "59077368cc922d1471b53a9f"
// customerMail : "fbruce@gmail.com"
// customerName : "Felipe Bruce"
// price : 77
// product : {name: "Ingresso Masculino - Promo", type: "TICKET"}
// name : "Ingresso Masculino - Promo"
// type : "TICKET"
// responsibleUserClubId : null
// responsibleUserClubName : null
// status : "AVAILABLE"
// transferedFrom : null
// type : "NORMAL"
// voucherId : "5b2f0652f2ae177c753bb23d"
// voucherUseDate : null
// voucherUseUserClubId : null
// voucherUseUserClubName : null

// SoftVoucher is a mongo document
type SoftVoucher struct {
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	Currency     string        `json:"currency" bson:"currency"`
	CustomerID   bson.ObjectId `json:"customerId" bson:"customerId"`
	CustomerName string        `json:"customerName" bson:"customerName"`
	CustomerMail string        `json:"customerMail" bson:"customerMail"`
	Price        float64       `json:"price" bson:"price"`

	Product VoucherProduct `json:"product" bson:"product"`

	Status    string        `json:"status" bson:"status"`
	Name      string        `json:"name" bson:"name"`
	Type      string        `json:"type" bson:"type"`
	VoucherID bson.ObjectId `bson:"voucherId" json:"voucherId"`

	TransferedFrom          *bson.ObjectId `json:"transferedFrom" bson:"transferedFrom"`
	ResponsableUserClubID   *bson.ObjectId `json:"responsibleUserClubId" bson:"responsibleUserClubId"`
	ResponsableUserClubName *string        `json:"responsibleUserClubName" bbbson:""`
	VoucherUseDate          *Timestamp     `bson:"voucherUseDate,omitempty" json:"voucherUseDate,omitempty"`
	VoucherUseUserClubID    *bson.ObjectId `json:"voucherUseUserClubId" bson:"voucherUseUserClubId"`
	VoucherUseUserClubName  string         `json:"voucherUseUserClubName" bson:"voucherUseUserClubName"`
}

// CompleteVoucher is a mongo document
type CompleteVoucher struct {
	ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreationDate   *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate     *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	StartDate      *Timestamp    `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate        *Timestamp    `bson:"endDate,omitempty" json:"endDate,omitempty"`
	CustomerID     bson.ObjectId `json:"customerId" bson:"customerId"`
	PartyID        bson.ObjectId `json:"partyId" bson:"partyId"`
	ClubID         bson.ObjectId `json:"clubId" bson:"clubId"`
	InvoiceID      bson.ObjectId `json:"invoiceId" bson:"invoiceId"`
	PartyProductID bson.ObjectId `json:"partyProductId" bson:"partyProductId"`
	ClubName       string        `json:"clubName" bson:"clubName"`
	PartyName      string        `json:"partyName" bson:"partyName"`
	CustomerName   string        `json:"customerName" bson:"customerName"`

	VoucherUseDate         *Timestamp     `bson:"voucherUseDate,omitempty" json:"voucherUseDate,omitempty"`
	VoucherUseUserClubID   *bson.ObjectId `json:"voucherUseUserClubId" bson:"voucherUseUserClubId"`
	ResponsableUserClubID  *bson.ObjectId `json:"responsibleUserClubId" bson:"responsibleUserClubId"`
	VoucherUseUserClubName string         `json:"voucherUseUserClubName" bson:"voucherUseUserClubName"`

	Status  string         `json:"status" bson:"status"`
	Price   Price          `json:"price" bson:"price"`
	Product VoucherProduct `json:"product" bson:"product"`
	Type    string         `json:"type" bson:"type"`

	Customer     *Customer     `json:"customer" bson:"customer"`
	Responsable  *UserClub     `json:"responsable" bson:"responsable"`
	PartyProduct *PartyProduct `json:"partyProduct,omitempty" bson:"partyProduct,omitempty" `
	Party        *Party        `json:"party,omitempty" bson:"party,omitempty" `
	Club         *Club         `json:"club,omitempty" bson:"club,omitempty" `

	TransferedFrom bson.ObjectId `json:"transferedFrom" bson:"transferedFrom"`
}

// Price sdfks
type Price struct {
	Value          float64 `json:"value" bson:"value"`
	CurrentIsoCode string  `json:"currencyIsoCode" bson:"currencyIsoCode"`
}

// Day sdfkjn
func (voucher Voucher) Day() time.Time {
	t := voucher.CreationDate.Time()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// Day sdfkjn
func (c CompleteVoucher) Day() time.Time {
	t := c.CreationDate.Time()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// VouchersResponse is a list of cards of a giver user
// swagger:response vouchersType
type VouchersResponse struct {
	// in: body
	Body []Voucher `json:"body"`
}

// VoucherResponse is the representation of a Voucher document on MongoDB
// swagger:response voucherType
type VoucherResponse struct {
	// in: body
	Body Voucher `json:"body"`
}

// VoucherPathParamID are the data you need to send in order to
// swagger:parameters  transferVoucher
type VoucherPathParamID struct {
	// in: path
	VoucherID string `json:"voucherId"`
}

// VoucherHistoryResume TODO: NEEDS COMMENT INFO
type VoucherHistoryResume struct {
	ID                     bson.ObjectId  `json:"id"`
	CreationDate           *Timestamp     `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate             *Timestamp     `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	StartDate              *Timestamp     `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate                *Timestamp     `bson:"endDate,omitempty" json:"endDate,omitempty"`
	VoucherUseDate         *Timestamp     `bson:"voucherUseDate,omitempty" json:"voucherUseDate,omitempty"`
	VoucherUseUserClubID   *bson.ObjectId `json:"voucherUseUserClubId" bson:"voucherUseUserClubId"`
	ResponsableUserClubID  *bson.ObjectId `json:"responsibleUserClubId" bson:"responsibleUserClubId"`
	VoucherUseUserClubName string         `json:"voucherUseUserClubName" bson:"voucherUseUserClubName"`
	Status                 string         `json:"status" bson:"status"`
	Price                  Price          `json:"price" bson:"price"`
	Product                VoucherProduct `json:"product" bson:"product"`
	Type                   string         `json:"type" bson:"type"`
	CustomerImage          string         `json:"customerImage" bson:"customerImage"`
	ProductImage           string         `json:"productImage" bson:"productImage"`
	PartyImage             string         `json:"partyImage" bson:"partyImage"`
	ClubImage              string         `json:"clubImage" bson:"clubImage"`
	ClubName               string         `json:"clubName" bson:"clubName"`
	PartyName              string         `json:"partyName" bson:"partyName"`
	ProductName            string         `json:"productName" bson:"productName"`
	CustomerName           string         `json:"customerName" bson:"customerName"`
}
