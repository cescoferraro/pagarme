package onni

import (
	"context"
	"errors"

	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// DeleteCard TODO: NEEDS COMMENT INFO
func DeleteCard(ctx context.Context, customer types.Customer, id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("not a objid")
	}
	repo, ok := ctx.Value(middlewares.CardsRepoKey).(interfaces.CardsRepo)
	if !ok {
		err := errors.New("context")
		return err
	}
	err := repo.DeleteByID(id)
	if err != nil {
		return err
	}
	cards, err := repo.GetByCustomerID(customer.ID.Hex())
	if err != nil {
		return err
	}
	for _, card := range cards {
		err := repo.ActivateByID(card.ID.Hex())
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
