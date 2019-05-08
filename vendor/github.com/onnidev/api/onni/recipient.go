package onni

import (
	"context"
	"errors"
	"log"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// GetRecipientByClubID TODO: NEEDS COMMENT INFO
func GetRecipientByClubID(ctx context.Context, id string) (types.Recipient, error) {
	repo, ok := ctx.Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	if !ok {
		err := errors.New("bug")
		return types.Recipient{}, err
	}
	recipients, err := repo.GetByClubID(id)
	if err != nil {
		return types.Recipient{}, err
	}
	for _, recipient := range recipients {
		if recipient.Status == "ACTIVE" {
			return recipient, nil
		}
	}
	return types.Recipient{}, errors.New("not active recipient")
}

// GetRecipientByID TODO: NEEDS COMMENT INFO
func GetRecipientByID(ctx context.Context, id string) (types.Recipient, error) {
	repo, ok := ctx.Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	if !ok {
		err := errors.New("bug")
		return types.Recipient{}, err
	}
	recipient, err := repo.GetByID(id)
	if err != nil {
		return types.Recipient{}, err
	}
	return recipient, nil
}

// GetRecipientByToken TODO: NEEDS COMMENT INFO
func GetRecipientByToken(ctx context.Context, id string) (types.Recipient, error) {
	repo, ok := ctx.Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	if !ok {
		err := errors.New("bug")
		return types.Recipient{}, err
	}
	recipient, err := repo.GetByToken(id)
	if err != nil {
		return types.Recipient{}, err
	}
	return recipient, nil
}

// CreateONNiRecipient TODO: NEEDS COMMENT INFO
func CreateONNiRecipient(
	ctx context.Context,
	post types.RecipientPost,
	recipient types.PagarMeRecipient,
) (types.Recipient, error) {
	repo, ok := ctx.Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	if !ok {
		err := errors.New("bug")
		return types.Recipient{}, err
	}
	var id bson.ObjectId
	if post.ID == "" {
		id = bson.NewObjectId()
	} else {
		id = bson.ObjectIdHex(post.ID)
	}
	err := repo.Collection.Insert(types.Recipient{
		ID:            id,
		ClubID:        bson.ObjectIdHex(post.ClubID),
		Type:          "ALL",
		Status:        "ACTIVE",
		RecipientID:   recipient.ID,
		BankAccountID: recipient.BankAccount.ID,
		BankingInfo: types.BankingInfo{
			BankCode:        recipient.BankAccount.BankCode,
			BankBranch:      recipient.BankAccount.Agencia,
			BankBranchVC:    recipient.BankAccount.AgenciaDV,
			BankAccount:     recipient.BankAccount.Conta,
			PersonType:      post.PersonType,
			DocumentNumber:  post.DocumentNumber,
			BankAccountVC:   recipient.BankAccount.ContaDV,
			BankAccountName: recipient.BankAccount.LegalName,
		},
	})
	if err != nil {
		return types.Recipient{}, err
	}

	recipients, err := repo.GetByClubID(post.ClubID)
	if err != nil {
		return types.Recipient{}, err
	}

	for _, rec := range recipients {
		if rec.ID.Hex() != id.Hex() {
			if rec.Status == "ACTIVE" {
				err := repo.Dull(rec.ID.Hex())
				if err != nil {
					continue
				}
			}
		}
	}
	result := types.Recipient{}
	err = repo.Collection.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return types.Recipient{}, err
	}
	log.Println("before exisintini onni pagamre ekjsen")
	return result, nil
}

// CreateALLRecipients TODO: NEEDS COMMENT INFO
func CreateALLRecipients(ctx context.Context, lead types.ClubLead) (types.PagarMeRecipient, types.Recipient, error) {
	pagarmerecipient := types.PagarMeRecipient{}
	onniRecipient := types.Recipient{}
	api := pagarme.New(viper.GetString("PAGARME"))
	pagarmerecipient, _, err := api.RecipientCreate(ctx, lead.RecipientPost())
	if err != nil {
		return pagarmerecipient, onniRecipient, err
	}
	log.Println("before creaTING onni RECIPIENT[]")
	onniRecipient, err = CreateONNiRecipient(ctx, lead.RecipientPost(), pagarmerecipient)
	if err != nil {
		return pagarmerecipient, onniRecipient, err
	}
	return pagarmerecipient, onniRecipient, nil
}
