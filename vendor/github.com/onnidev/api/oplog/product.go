package oplog

import (
	"fmt"

	"github.com/onnidev/api/types"
	"github.com/rwynn/gtm"
	"gopkg.in/mgo.v2/bson"
)

func productOPLOG(op *gtm.Op) {
	OPLOGGLOBAL.Count = OPLOGGLOBAL.Count + 1
	if op.IsInsert() {
		insertProductOPLOG(op)
	}
	if op.IsUpdate() {
		updateProductOPLOG(op)
	}
}

func insertProductOPLOG(op *gtm.Op) {
	product := op.Data
	image := product["image"].(gtm.OpLogEntry)
	theproduct := types.Product{
		Image: types.Image{
			MimeType:     image["mimeType"].(string),
			FileID:       image["fileId"].(bson.ObjectId),
			CreationDate: image["creationDate"].(*types.Timestamp),
		},
		ID: product["_id"].(bson.ObjectId),
		// CreationDate: product["creationDate"].(*types.Timestamp),
		CreationDate: product["creationDate"].(*types.Timestamp),
		Name:         product["name"].(string),
		Type:         product["type"].(string),
		NameSort:     product["nameSort"].(string),
		Category:     product["category"].(string),
	}
	OPLOGGLOBAL.LastProduct = theproduct
	fmt.Println(product)
}

func updateProductOPLOG(op *gtm.Op) {
	product := op.Data
	image := product["image"].(map[string]interface{})
	theproduct := types.Product{
		Image: types.Image{
			MimeType:     image["mimeType"].(string),
			FileID:       image["fileId"].(bson.ObjectId),
			CreationDate: image["creationDate"].(*types.Timestamp),
		},
		ID:           product["_id"].(bson.ObjectId),
		CreationDate: product["creationDate"].(*types.Timestamp),
		Name:         product["name"].(string),
		Type:         product["type"].(string),
		NameSort:     product["nameSort"].(string),
		Category:     product["category"].(string),
	}
	OPLOGGLOBAL.LastProduct = theproduct
	fmt.Println(product)
}
