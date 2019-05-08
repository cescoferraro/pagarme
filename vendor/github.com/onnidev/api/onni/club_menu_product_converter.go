package onni

import (
	"context"
	"log"
	"time"

	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// ConvertProductsReqToMenuTickets TODO: NEEDS COMMENT INFO
func ConvertProductsReqToMenuProducts(ctx context.Context, products []types.SoftProductPostPartyProduct) ([]types.MenuProduct, error) {
	newProducts := []types.MenuProduct{}
	for _, product := range products {
		productxxx, err := SoftProductPartyProductSingleConverter(ctx, product)
		if err != nil {
			return newProducts, err
		}
		newProducts = append(newProducts, productxxx)
	}
	log.Println("final da conversão de productos")
	return newProducts, nil
}

// SoftProductPartyProductSingleConverter TODO: NEEDS COMMENT INFO
func SoftProductPartyProductSingleConverter(ctx context.Context, product types.SoftProductPostPartyProduct) (types.MenuProduct, error) {
	result := types.MenuProduct{}
	log.Println("=====================")
	log.Println(product.ExhibitionName)
	after := types.Timestamp(time.Time{})
	if product.DoNotSellAfter != nil {
		if !product.DoNotSellAfter.Time().IsZero() {
			after = *product.DoNotSellAfter
		}
	}
	before := types.Timestamp(time.Time{})
	if product.DoNotSellBefore != nil {
		if !product.DoNotSellBefore.Time().IsZero() {
			before = *product.DoNotSellBefore
		}
	}
	log.Println("foto do produto")
	log.Println("foto do produto")
	log.Println("foto do produto")
	log.Println("foto do produto")
	log.Println("foto do produto")
	log.Println("outside da function")
	image, err := ImageONNiCombo(ctx, product.Image)
	if err != nil {
		return result, err
	}
	log.Println("imagem criada to pianito")
	id := bson.NewObjectId()
	if bson.IsObjectIdHex(product.MenuProductID) {
		id = bson.ObjectIdHex(product.MenuProductID)
	}
	result = types.MenuProduct{
		ID:                 id,
		Price:              types.Price{Value: product.Price, CurrentIsoCode: "BRL"},
		ExhibitionName:     product.ExhibitionName,
		Category:           product.Category,
		DoNotSellMoreThan:  product.DoNotSellMoreThan,
		Featured:           product.Featured,
		DoNotSellBefore:    &before,
		DoNotSellAfter:     &after,
		Active:             product.Active,
		GeneralInformation: product.GeneralInformation,
		Image:              image,
		Combo:              product.Combo,
	}
	if product.ProductID != "" {
		if bson.IsObjectIdHex(product.ProductID) {
			id := bson.ObjectIdHex(product.ProductID)
			result.ProductID = &id
		}
	}
	return result, nil
}
