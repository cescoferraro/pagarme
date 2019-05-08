package onni

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// CreatePartyProductsDrink TODO: NEEDS COMMENT INFO
func CreatePartyProductsDrink(ctx context.Context, products []types.SoftProductPostPartyProduct, menuid, partyID bson.ObjectId) error {
	productsCollection, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("assert")
		return err
	}
	horario := types.Timestamp(time.Now())
	empty := types.Timestamp(time.Time{})
	for _, product := range products {
		image, err := ImageONNi(ctx, product.Image)
		if err != nil {
			return err
		}
		instance := types.PartyProduct{
			ID:                bson.NewObjectId(),
			Name:              product.ExhibitionName,
			CreationDate:      &horario,
			PartyID:           partyID,
			MoneyAmount:       types.Price{Value: product.Price, CurrentIsoCode: "BRL"},
			Image:             &image,
			DoNotSellMoreThan: product.DoNotSellMoreThan,
			DoNotSellBefore:   &empty,
			DoNotSellAfter:    &empty,
			QuantityPurchased: 0.0,
			Category:          &product.Category,
			Type:              "DRINK",
			// Product
			QuantityFree:  0,
			MenuProductID: &menuid,

			AcceptsEvaluation: false,
			Featured:          false,
			Combo:             product.Combo,
			Status:            "INACTIVE",
		}
		if product.DoNotSellBefore != nil {
			instance.DoNotSellBefore = product.DoNotSellBefore
		}
		if product.DoNotSellAfter != nil {
			instance.DoNotSellAfter = product.DoNotSellAfter
		}
		if product.Active {
			instance.Status = "ACTIVE"
		}
		log.Println("adding a product [ ]")
		err = productsCollection.Collection.Insert(instance)
		if err != nil {
			return err
		}

	}
	return nil
}
