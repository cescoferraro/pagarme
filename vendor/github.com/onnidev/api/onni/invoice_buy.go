package onni

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/types"
)

// PatchVoucherBuyError TODO: NEEDS COMMENT INFO
func PatchVoucherBuyError(voucher types.Voucher) error {
	db, err := infra.Cloner()
	if err != nil {
		return err
	}
	defer db.Session.Close()
	repo, err := interfaces.NewVoucherCollection(db)
	if err != nil {
		return err
	}
	now := types.Timestamp(time.Now())
	update := bson.M{
		"updateDate": &now,
		"status":     "ERROR",
	}
	change := mgo.Change{
		Update:    bson.M{"$set": update},
		ReturnNew: true,
	}
	patchedVoucher := types.Voucher{}
	_, err = repo.Collection.Find(bson.M{"_id": voucher.ID}).Apply(change, &patchedVoucher)
	if err != nil {
		return err
	}
	return nil
}

// PatchVoucherBuySuccess TODO: NEEDS COMMENT INFO
func PatchVoucherBuySuccess(voucher types.Voucher, transactionID string) error {
	db, err := infra.Cloner()
	if err != nil {
		return err
	}
	defer db.Session.Close()
	repo, err := interfaces.NewVoucherCollection(db)
	if err != nil {
		return err
	}
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate":    &now,
				"status":        "AVAILABLE",
				"transactionId": &transactionID,
			}},
		ReturnNew: true,
	}
	patchedVoucher := types.Voucher{}
	_, err = repo.Collection.Find(bson.M{"_id": voucher.ID}).Apply(change, &patchedVoucher)
	if err != nil {
		return err
	}
	return nil
}

// PatchInvoiceBuyError TODO: NEEDS COMMENT INFO
func PatchInvoiceBuyError(invoice types.Invoice) error {
	db, err := infra.Cloner()
	if err != nil {
		return err
	}
	defer db.Session.Close()
	repo, err := interfaces.NewInvoicesCollection(db)
	if err != nil {
		return err
	}
	now := types.Timestamp(time.Now())
	update := bson.M{
		"updateDate": &now,
		"status":     "ERROR",
	}
	change := mgo.Change{Update: bson.M{"$set": update}, ReturnNew: true}
	patchedInvoice := types.Invoice{}
	_, err = repo.Collection.Find(bson.M{"_id": invoice.ID}).Apply(change, &patchedInvoice)
	if err != nil {
		return err
	}
	return nil
}

// PatchDrinksInvoiceBuySuccess TODO: NEEDS COMMENT INFO
func PatchDrinksInvoiceBuySuccess(club types.Club, party types.Party, invoice types.Invoice, products types.BuyPostList, transactionID string) error {
	db, err := infra.Cloner()
	if err != nil {
		return err
	}
	defer db.Session.Close()
	repo, err := interfaces.NewInvoicesCollection(db)
	if err != nil {
		return err
	}
	now := types.Timestamp(time.Now())
	percent := (100 - club.PercentDrink)
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate":    &now,
				"status":        "SUCCESS",
				"transactionId": &transactionID,
				"valueToOnni":   products.SumDrinksONNi(party, club),
				"valueToClub":   products.SumDrinksClub(party, club),
				"percentToOnni": &percent,
				"percentToClub": &club.PercentDrink,
			}},
		ReturnNew: true,
	}
	patchedInvoice := types.Invoice{}
	_, err = repo.Collection.Find(bson.M{"_id": invoice.ID}).Apply(change, &patchedInvoice)
	if err != nil {
		return err
	}
	return nil
}

// PatchTicketsInvoiceBuySuccess TODO: NEEDS COMMENT INFO
func PatchTicketsInvoiceBuySuccess(club types.Club, party types.Party, invoice types.Invoice, products types.BuyPostList, transactionID string) error {
	log.Println(">>>> patching ticket invoice")
	db, err := infra.Cloner()
	if err != nil {
		return err
	}
	defer db.Session.Close()
	repo, err := interfaces.NewInvoicesCollection(db)
	if err != nil {
		return err
	}
	now := types.Timestamp(time.Now())
	percent := (100 - club.PercentTicket)
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate":    &now,
				"status":        "SUCCESS",
				"transactionId": &transactionID,
				"valueToOnni":   products.SumTicketsONNi(party, club),
				"valueToClub":   products.SumTicketsClub(party, club),
				"percentToOnni": &percent,
				"percentToClub": &club.PercentDrink,
			}},
		ReturnNew: true,
	}
	patchedInvoice := types.Invoice{}
	_, err = repo.Collection.Find(bson.M{"_id": invoice.ID}).Apply(change, &patchedInvoice)
	if err != nil {
		return err
	}
	return nil
}

// InvoicesBuySuccess TODO: NEEDS COMMENT INFO
func InvoicesBuySuccess(club types.Club, party types.Party, invoices types.Invoices, products types.BuyPostList, transactionID string) error {
	log.Println("invoice suceess")
	if invoices.Drinks != nil {
		log.Println("invoice drinks suceess")
		invoice := *invoices.Drinks
		err := PatchDrinksInvoiceBuySuccess(club, party, invoice, products, transactionID)
		if err != nil {
			return err
		}
	}
	if invoices.Tickets != nil {
		log.Println("invoice tickets suceess")
		invoice := *invoices.Tickets
		err := PatchTicketsInvoiceBuySuccess(club, party, invoice, products, transactionID)
		if err != nil {
			return err
		}
	}
	return nil
}

// InvoicesBuyError TODO: NEEDS COMMENT INFO
func InvoicesBuyError(invoices []types.Invoice) error {
	for _, invoice := range invoices {
		err := PatchInvoiceBuyError(invoice)
		if err != nil {
			return err
		}
	}
	return nil
}
