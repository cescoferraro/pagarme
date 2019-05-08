package onni

import (
	"context"
	"errors"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// GetPagarmeRecipientBalance TODO: NEEDS COMMENT INFO
func GetPagarmeRecipientBalance(ctx context.Context,
	recipientID string) (types.PagarMeTransactionBalance, error) {
	var balance types.PagarMeTransactionBalance
	key := middlewares.RecipientCollectionKey
	recipientRepo, ok := ctx.Value(key).(interfaces.RecipientsRepo)
	if !ok {
		return balance, errors.New("context missing")
	}
	recipient, err := recipientRepo.GetByToken(recipientID)
	if err != nil {
		return balance, err
	}
	api := pagarme.New(viper.GetString("PAGARME"))
	balance, err = api.RecipientBalance(ctx, recipient.RecipientID)
	if err != nil {
		return balance, err
	}
	return balance, nil
}

// GetFinanceQueryFromBody TODO: NEEDS COMMENT INFO
func GetFinanceQueryFromBody(ctx context.Context) (types.FinanceQuery, error) {
	filter, ok := ctx.Value(middlewares.FinanceQueryReq).(types.FinanceQuery)
	if !ok {
		return filter, errors.New("missing context")
	}
	return filter, nil
}
