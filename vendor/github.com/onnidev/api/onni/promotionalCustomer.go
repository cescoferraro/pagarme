package onni

import (
	"context"
	"errors"
	"log"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// PersistPromotionalCustomer TODO: NEEDS COMMENT INFO
func PersistPromotionalCustomer(ctx context.Context, party types.Party, partyProduct types.PartyProduct, prom types.PromotionalCustomer, email string) error {
	repo, ok := ctx.Value(middlewares.PromotionalCustomerRepoKey).(interfaces.PromotionalCustomerRepo)
	if !ok {
		return errors.New("2bug assert")
	}
	err := repo.Collection.Insert(prom)
	if err != nil {
		return err
	}
	log.Println("calling function to sending email to ", email)
	go PromoNewMail(ctx, party, partyProduct, prom, email)
	return nil
}

// PromotionalCustomer TODO: NEEDS COMMENT INFO
func PromotionalCustomer(ctx context.Context, promotionID string) (types.PromotionalCustomer, error) {
	repo, ok := ctx.Value(middlewares.PromotionalCustomerRepoKey).(interfaces.PromotionalCustomerRepo)
	if !ok {
		return types.PromotionalCustomer{}, errors.New("2bug assert")
	}
	promo, err := repo.GetByID(promotionID)
	if err != nil {
		return types.PromotionalCustomer{}, err
	}
	return promo, nil
}
