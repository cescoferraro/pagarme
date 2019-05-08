package onni

import (
	"context"
	"errors"
	"log"
	"math"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// BuyPagarme TODO: NEEDS COMMENT INFO
func BuyPagarme(ctx context.Context, req types.PagarMeTransactionRequest) (types.PagarMeTransactionResponse, error) {
	result := types.PagarMeTransactionResponse{}
	ctx, cancel := context.WithTimeout(ctx, 400*time.Second)
	defer cancel()
	token := viper.GetString("PAGARME")
	pagarmerecipient, err := pagarme.New(token).TransactionCreate(ctx, req)
	if err != nil {
		return result, err
	}
	if pagarmerecipient.Status != "paid" {
		err := errors.New("not paid")
		log.Println("nao tinha grana mano ", err.Error())
		return pagarmerecipient, err
	}
	return pagarmerecipient, nil
}

// CreateBuyRequest TODO: NEEDS COMMENT INFO
func CreateBuyRequest(ctx context.Context,
	club types.Club,
	party types.Party,
	customer types.Customer,
	products types.BuyPostList,
) (types.PagarMeTransactionRequest, error) {
	result := types.PagarMeTransactionRequest{}
	if customer.DocumentNumber == nil {
		err := errors.New("assserrt errro")
		return result, err
	}
	cpf := ""
	if customer.DocumentNumber != nil {
		id := *customer.DocumentNumber
		if id == "" {
			err := errors.New("assserrt errro")
			return result, err
		}
		cpf = id
	}
	cardsRepo, ok := ctx.Value(middlewares.CardsRepoKey).(interfaces.CardsRepo)
	if !ok {
		err := errors.New("assserrt errro")
		return result, err
	}
	card, err := cardsRepo.GetByCustomerDefaultCard(customer.ID.Hex())
	if err != nil {
		return result, err
	}
	token := viper.GetString("PAGARME")
	name := club.Name
	if len(name) > 12 {
		name = name[:11]
	}
	recipient, err := GetRecipientByClubID(ctx, club.ID.Hex())
	if err != nil {
		return result, err
	}
	return types.PagarMeTransactionRequest{
		Amount:         products.SumBuy(party, club),
		PaymentMethod:  "credit_card",
		APIKey:         token,
		Async:          "false",
		CardID:         card.CardToken,
		SoftDescriptor: name,
		Capture:        "true",
		PagarMeCustomer: types.PagarMeCustomer{
			Type:           "individual",
			Country:        "BR",
			Name:           customer.Name(),
			Email:          customer.Mail,
			DocumentNumber: cpf,
			Documents:      []types.Document{{Type: "cpf"}},
		},
		SplitRules: []types.SplitRule{
			{
				Amount:              products.SumONNiString(party, club),
				RecipientID:         onniRecipient(),
				Liable:              "true",
				ChargeProcessingFee: "true",
				ChargeRemainderFee:  "true",
			},
			{
				Amount:              products.SumClubString(party, club),
				RecipientID:         recipient.RecipientID,
				Liable:              liabilityCheck(ctx, club),
				ChargeProcessingFee: "false",
				ChargeRemainderFee:  "false",
			},
		},
		MetaData: types.MetaData{
			ClubID:     club.ID.Hex(),
			ClubName:   club.Name,
			CustomerID: customer.ID.Hex(),
			PartyID:    party.ID.Hex(),
			PartyName:  party.Name,
		},
	}, nil
}

func liabilityCheck(ctx context.Context, club types.Club) string {
	if club.Liability != nil {
		if *club.Liability == "ONNI" {
			return "false"
		}
	}
	return "true"
}
func floor(value float64) float64 {
	return math.Floor(value*100) / 100
}
func onniRecipient() string {
	if viper.GetString("env") == "production" {
		return "re_ciw98fu3s07vrcv5xm0r77hnp"
	}
	if viper.GetString("env") == "homolog" {
		return "re_cjip2e4if01vzmb6d86hqh9ju"
	}
	return "re_ciw98fu3s07vrcv5xm0r77hnp"
}
