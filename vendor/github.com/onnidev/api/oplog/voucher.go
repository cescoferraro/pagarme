package oplog

import (
	"fmt"

	"github.com/onnidev/api/types"
	"github.com/rwynn/gtm"
	"gopkg.in/mgo.v2/bson"
)

func voucherOPLOG(op *gtm.Op) {
	OPLOGGLOBAL.Count = OPLOGGLOBAL.Count + 1
	voucher := types.Voucher{
		ID:             op.Data["_id"].(bson.ObjectId),
		CreationDate:   op.Data["creationDate"].(*types.Timestamp),
		StartDate:      op.Data["startDate"].(*types.Timestamp),
		EndDate:        op.Data["endDate"].(*types.Timestamp),
		PartyID:        op.Data["partyId"].(bson.ObjectId),
		CustomerID:     op.Data["customerId"].(bson.ObjectId),
		ClubID:         op.Data["clubId"].(bson.ObjectId),
		PartyProductID: op.Data["partyProductId"].(bson.ObjectId),
		PartyName:      op.Data["partyName"].(string),
		Status:         op.Data["status"].(string),
		Type:           op.Data["type"].(string),
		ClubName:       op.Data["clubName"].(string),
		CustomerName:   op.Data["customerName"].(string),
	}
	if op.IsInsert() {
		insertVoucherOP(op, &voucher)
	}
	if op.IsUpdate() {
		updateVoucherOP(op, &voucher)
	}
	omnipotemtyOPLOG(op, &voucher)
}

func updateVoucherOP(op *gtm.Op, voucher *types.Voucher) {
	product := op.Data["product"].(map[string]interface{})
	price := op.Data["price"].(map[string]interface{})
	image := product["image"].(map[string]interface{})
	theprice := types.Price{
		Value:          price["value"].(float64),
		CurrentIsoCode: price["currencyIsoCode"].(string),
	}
	theproduct := types.VoucherProduct{
		Image: types.Image{
			MimeType:     image["mimeType"].(string),
			FileID:       image["fileId"].(bson.ObjectId),
			CreationDate: image["creationDate"].(*types.Timestamp),
		},
		Name: product["name"].(string),
		Type: product["type"].(string),
	}
	voucher.Product = theproduct
	voucher.Price = theprice
}

func omnipotemtyOPLOG(op *gtm.Op, voucher *types.Voucher) {
	upDate, ok := op.Data["updateDate"].(*types.Timestamp)
	if ok {
		voucher.UpdateDate = upDate
	}
	useDate, ok := op.Data["voucherUseDate"].(*types.Timestamp)
	if ok {
		voucher.VoucherUseDate = useDate
	}
	useUserClubID, ok := op.Data["voucherUseUserClubId"].(bson.ObjectId)
	if ok {
		voucher.VoucherUseUserClubID = &useUserClubID
	}
	useUserClubName, ok := op.Data["voucherUseUserClubName"].(string)
	if ok {
		voucher.VoucherUseUserClubName = &useUserClubName
	}
	responsableUserClubID, ok := op.Data["responsibleUserClubId"].(bson.ObjectId)
	if ok {
		voucher.ResponsableUserClubID = &responsableUserClubID
	}
	OPLOGGLOBAL.LastVoucher = *voucher
	fmt.Println(voucher)
}

func insertVoucherOP(op *gtm.Op, voucher *types.Voucher) {
	product := op.Data["product"].(gtm.OpLogEntry)
	price := op.Data["price"].(gtm.OpLogEntry)
	image := product["image"].(gtm.OpLogEntry)
	theprice := types.Price{
		Value:          price["value"].(float64),
		CurrentIsoCode: price["currencyIsoCode"].(string),
	}
	theproduct := types.VoucherProduct{
		Image: types.Image{
			MimeType:     image["mimeType"].(string),
			FileID:       image["fileId"].(bson.ObjectId),
			CreationDate: image["creationDate"].(*types.Timestamp),
		},
		Name: product["name"].(string),
		Type: product["type"].(string),
	}
	voucher.Product = theproduct
	voucher.Price = theprice
}
