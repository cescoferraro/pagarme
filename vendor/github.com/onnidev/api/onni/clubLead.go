package onni

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PatchLead TODO: NEEDS COMMENT INFO
func PatchLead(ctx context.Context, id string, patch types.ClubLeadPatch) (types.ClubLead, error) {
	clubLead := types.ClubLead{}
	repo := ctx.Value(middlewares.ClubLeadKey).(interfaces.ClubLeadRepo)
	lead, err := repo.GetByID(id)
	if err != nil {
		return clubLead, err
	}
	log.Println(lead.Stage)
	stage, err := stageOrBlank(lead.Stage, patch.Stage)
	if err != nil {
		return clubLead, err
	}
	horario := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate":         &horario,
			"adminName":          orBlank(lead.AdminName, patch.AdminName),
			"adminPhone":         orBlank(lead.AdminPhone, patch.AdminPhone),
			"adminMail":          orBlank(lead.AdminMail, patch.AdminMail),
			"clubName":           orBlank(lead.ClubName, patch.ClubName),
			"clubDescription":    orBlank(lead.ClubDescription, patch.ClubDescription),
			"stage":              stage,
			"musicStyles":        orBlankStyles(lead.MusicStyles, patch.MusicStyles),
			"state":              orBlank(lead.State, patch.State),
			"street":             orBlank(lead.Street, patch.Street),
			"country":            orBlank(lead.Country, patch.Country),
			"city":               orBlank(lead.City, patch.City),
			"number":             orBlank(lead.Number, patch.Number),
			"unit":               orBlank(lead.Unit, patch.Unit),
			"clubType":           orBlank(lead.ClubType, patch.ClubType),
			"personType":         orBlank(lead.PersonType, patch.PersonType),
			"image":              orBlank(lead.Image, patch.Image),
			"backgroundImage":    orBlank(lead.BackgroundImage, patch.BackgroundImage),
			"fbURL":              orBlank(lead.FBURL, patch.FBURL),
			"bankCode":           orBlank(lead.BankCode, patch.BankCode),
			"bankLegalAddress":   orBlank(lead.BankLegalAddress, patch.BankLegalAddress),
			"bankBranch":         orBlank(lead.BankBranch, patch.BankBranch),
			"bankBranchVC":       orBlank(lead.BankBranchVC, patch.BankBranchVC),
			"bankAccount":        orBlank(lead.BankAccount, patch.BankAccount),
			"bankAccountVC":      orBlank(lead.BankAccountVC, patch.BankAccountVC),
			"bankAccountName":    orBlank(lead.BankAccountName, patch.BankAccountName),
			"bankAccountType":    orBlank(lead.BankAccountType, patch.BankAccountType),
			"bankDocumentNumber": orBlank(lead.BankDocumentNumber, patch.BankDocumentNumber),
		}},
		ReturnNew: true,
	}
	_, err = repo.Collection.Find(bson.M{"_id": lead.ID}).Apply(change, &clubLead)
	if err != nil {
		return clubLead, err
	}
	return clubLead, nil
}

func stageOrBlank(og, sent string) (string, error) {
	if sent == "" {
		return og, nil
	}
	asent, err := strconv.Atoi(sent)
	if err != nil {
		return "", err
	}
	aog, err := strconv.Atoi(og)
	if err != nil {
		return "", err
	}
	if asent > aog {
		return sent, nil
	}
	return og, nil
}

func orBlankStyles(og, sent []string) []string {
	if len(sent) != 0 {
		return sent
	}
	return og
}
func orBlank(og, sent string) string {
	if sent == "" {
		return og
	}
	return sent
}

func orBlankObjectID(og bson.ObjectId, sent string) bson.ObjectId {
	if sent == "" || !bson.ObjectIdHex(sent).Valid() {
		return og
	}
	return bson.ObjectIdHex(sent)
}
