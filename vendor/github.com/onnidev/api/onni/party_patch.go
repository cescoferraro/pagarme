package onni

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PatchParty TODO: NEEDS COMMENT INFO
func PatchParty(ctx context.Context, party types.Party, req types.SoftPartyPostRequest, musicStyles []types.Style, clubMenuProductID *bson.ObjectId, clubMenuTicketID *bson.ObjectId) (types.Party, error) {
	result := types.Party{}
	repo, ok := ctx.Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	if !ok {
		err := errors.New("assert bug")
		return result, err
	}
	now := types.Timestamp(time.Now())
	set := bson.M{
		"updateDate":        &now,
		"name":              shared.OrBlank(party.Name, req.Name),
		"startDate":         &req.StartDate,
		"endDate":           &req.EndDate,
		"description":       &req.Description,
		"musicStyles":       musicStyles,
		"address":           req.Address,
		"assumeServiceFee":  req.AssumeServiceFee,
		"clubMenuProductId": nil,
		"clubMenuTicketId":  nil,
	}
	if clubMenuProductID != nil {
		id := *clubMenuProductID
		if id.Valid() {
			log.Println("setting a ticket menu on the party ===========")
			set["clubMenuProductId"] = clubMenuProductID
		}
	}
	if clubMenuTicketID != nil {
		id := *clubMenuTicketID
		if id.Valid() {
			log.Println("setting a clubmenuTicket the party ===========")
			set["clubMenuTicketId"] = clubMenuTicketID
		}
	}
	log.Println("===================================")
	log.Println("before patching party")
	log.Println(party.ClubMenuTicketID)
	log.Println(party.ClubMenuProductID)
	log.Println("===================================")
	change := mgo.Change{Update: bson.M{"$set": set}, ReturnNew: true}
	var patchedParty types.Party
	_, err := repo.Collection.Find(bson.M{"_id": party.ID}).Apply(change, &patchedParty)
	if err != nil {
		return result, err
	}
	log.Println("===================================")
	log.Println("after patching party")
	log.Println(patchedParty.ClubMenuTicketID)
	log.Println(patchedParty.ClubMenuProductID)
	log.Println("===================================")
	return patchedParty, err

}
