package recipient

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Patch skdjfn
func Patch(w http.ResponseWriter, r *http.Request) {
	recipientRepo := r.Context().Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	recipientPatch := r.Context().Value(middlewares.RecipientPatchKey).(types.RecipientPatch)
	j, _ := json.MarshalIndent(recipientPatch, "", "    ")
	log.Println("Recieved from mongo")
	log.Println(string(j))
	id := chi.URLParam(r, "recipientID")
	log.Println(id)
	recipient, err := onni.GetRecipientByID(r.Context(), id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	recipients, err := recipientRepo.GetByClubID(recipient.ClubID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if recipientPatch.Status == "INACTIVE" {
		if recipient.Status == "ACTIVE" {
			err := errors.New("Make another recipient ACTIVE instead")
			shared.MakeONNiError(w, r, 400, err)
			return
		}

	}
	log.Println(recipientPatch.Status)
	if recipientPatch.Status == "ACTIVE" {
		for _, rec := range recipients {
			if rec.Status == "ACTIVE" {
				err := recipientRepo.Dull(rec.ID.Hex())
				if err != nil {
					shared.MakeONNiError(w, r, 400, err)
					return
				}
			}

		}

	}
	updateDate := types.Timestamp(time.Now())
	conta := types.BankingInfo{
		BankCode:        recipientPatch.BankCode,
		BankBranch:      recipientPatch.BankBranch,
		BankBranchVC:    recipient.BankingInfo.BankBranchVC,
		BankAccount:     recipientPatch.BankAccount,
		BankAccountVC:   recipientPatch.BankAccountVC,
		PersonType:      recipient.BankingInfo.PersonType,
		DocumentNumber:  recipient.BankingInfo.DocumentNumber,
		BankAccountName: recipientPatch.BankAccountName,
	}
	if recipientPatch.BankBranchVC != "" {
		conta.BankBranchVC = &recipientPatch.BankBranchVC
	} else {
		conta.BankBranchVC = recipient.BankingInfo.BankBranchVC
	}
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate":  &updateDate,
			"type":        orBlank(recipient.Type, recipientPatch.Type),
			"status":      orBlank(recipient.Status, recipientPatch.Status),
			"bankingInfo": conta,
		}},
		ReturnNew: true,
	}
	var patchedRecipient types.Recipient
	_, err = recipientRepo.Collection.
		Find(bson.M{"_id": bson.ObjectIdHex(id)}).
		Apply(change, &patchedRecipient)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	j, _ = json.MarshalIndent(patchedRecipient, "", "    ")
	log.Println("Recieved from mongo")
	log.Println(string(j))
	api := pagarme.New(viper.GetString("PAGARME"))

	editedpagarmerecipient, code, err := api.RecipientEditBankingInfo(r.Context(), patchedRecipient)
	if err != nil {
		shared.MakeONNiError(w, r, code, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, editedpagarmerecipient)
}

func orBlank(og, sent string) string {
	if sent == "" {
		return og
	}
	return sent
}
