package onni

import (
	"context"
	"errors"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// TicketStatus TODO: NEEDS COMMENT INFO
func TicketStatus(ctx context.Context, products []types.CartPartyProduct, partyID string) (string, error) {
	return typeStatus(ctx, products, partyID, "TICKET")
}

// DrinkStatus TODO: NEEDS COMMENT INFO
func DrinkStatus(ctx context.Context, result map[string][]types.CartPartyProduct, partyID string) (string, error) {
	products := []types.CartPartyProduct{}
	for key, value := range result {
		if key != "_empty" {
			products = append(products, value...)
		}

	}
	return typeStatus(ctx, products, partyID, "DRINK")
}

func typeStatus(ctx context.Context, products []types.CartPartyProduct, partyID string, ourtype string) (string, error) {
	existType := func(products []types.CartPartyProduct, ourtype string) bool {
		for _, product := range products {
			if product.Type == ourtype {
				return true
			}
		}
		return false
	}
	if len(products) > 0 && existType(products, ourtype) {
		return "AVAILABLE", nil
	}
	partyProductCollection, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("dumb bug")
		return "", err
	}
	count, err := partyProductCollection.CountByPartyIDandType(partyID, ourtype)
	if err != nil {
		return "", err
	}
	if count == 0 {
		return "N_A", nil
	}
	return "ZERO", nil
}
