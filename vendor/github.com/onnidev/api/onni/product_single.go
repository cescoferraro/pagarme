package onni

import (
	"context"
	"errors"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// CreateProduct TODO: NEEDS COMMENT INFO
func CreateProduct(ctx context.Context, product types.Product) error {
	productsCollection, ok := ctx.Value(middlewares.ProductsRepoKey).(interfaces.ProductsRepo)
	if !ok {
		err := errors.New("bug assert")
		return err
	}
	return productsCollection.Collection.Insert(product)
}

// GetProduct TODO: NEEDS COMMENT INFO
func GetProduct(ctx context.Context, id string) (types.Product, error) {
	result := types.Product{}
	productsCollection, ok := ctx.Value(middlewares.ProductsRepoKey).(interfaces.ProductsRepo)
	if !ok {
		err := errors.New("bug assert")
		return result, err
	}
	product, err := productsCollection.GetByID(id)
	if err != nil {
		return product, err
	}
	return product, nil
}
