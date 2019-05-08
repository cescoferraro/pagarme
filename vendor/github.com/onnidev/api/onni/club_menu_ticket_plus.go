package onni

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UpdateCurrentClubMenuTicketAndPartyProducts TODO: NEEDS COMMENT INFO
func UpdateCurrentClubMenuTicketAndPartyProducts(ctx context.Context, party types.Party, req types.SoftPartyPostRequest) (types.ClubMenuTicket, error) {
	menuid := bson.ObjectIdHex(*req.ClubMenuTicket.ID)
	result := types.ClubMenuTicket{}
	repo, ok := ctx.Value(middlewares.ClubMenuTicketRepoKey).(interfaces.ClubMenuTicketRepo)
	if !ok {
		err := errors.New("assert")
		return result, err
	}
	menu, err := repo.GetByID(menuid.Hex())
	if err != nil {
		return result, err
	}
	err = DeprecateUnusedTicket(ctx, party, req)
	if err != nil {
		return result, err
	}

	for _, ticket := range req.Tickets {
		log.Println("====================================")
		if bson.IsObjectIdHex(ticket.MenuTicketID) {
			// sdfkjn
			err := DealWithMenuTicket(ctx, ticket, party, menu, req.ClubMenuTicket.Name)
			if err != nil {
				return result, err
			}
			continue
		}
		if !bson.IsObjectIdHex(ticket.MenuTicketID) {
			err := AddMenuTicket(ctx, ticket, party, req)
			if err != nil {
				return result, err
			}
		}
	}
	return menu, nil

}

// DealWithMenuTicket TODO: NEEDS COMMENT INFO
func DealWithMenuTicket(ctx context.Context, ticket types.SoftTicketPostPartyProduct, party types.Party, menu types.ClubMenuTicket, newname string) error {
	log.Println("########################################")
	log.Println("##### EDITING an existing menuTicket")
	batches := []types.MenuTicket{}
	pbatches := []types.ObjectPartyProduct{}
	for _, batch := range ticket.Batches {
		after := types.Timestamp(time.Time{})
		before := types.Timestamp(time.Time{})
		log.Println("###### a batch")
		log.Println(batch.Active)
		if bson.IsObjectIdHex(batch.PartyProductID) {
			pbatches = append(pbatches, types.ObjectPartyProduct{
				PartyProductID: bson.ObjectIdHex(batch.PartyProductID),
			})
		}
		instance := types.MenuTicket{}
		id := bson.NewObjectId()
		if len(batch.MenuTicketID) != 0 {
			id = bson.ObjectIdHex(batch.MenuTicketID)
		}
		log.Println(batch.Active)
		log.Println(batch.Active)
		if batch.DoNotSellBefore != nil {
			before = *batch.DoNotSellBefore
		}
		if batch.DoNotSellAfter != nil {
			after = *batch.DoNotSellAfter
		}
		instance = types.MenuTicket{
			ID:                     id,
			Price:                  types.Price{Value: batch.Price, CurrentIsoCode: "BRL"},
			ExhibitionName:         batch.ExhibitionName,
			DoNotSellMoreThan:      batch.DoNotSellMoreThan,
			DoNotSellBefore:        &before,
			DoNotSellAfter:         &after,
			Active:                 batch.Active,
			GeneralInformation:     ticket.AdditionalInformations.General,
			FreeInformation:        ticket.AdditionalInformations.Free,
			AnniversaryInformation: ticket.AdditionalInformations.Anniversary,
			Batches:                []types.MenuTicket{},
		}
		if bson.IsObjectIdHex(batch.PartyProductID) {
			err := PatchTicketPartyProduct(ctx, batch.PartyProductID, instance, []types.ObjectPartyProduct{})
			if err != nil {
				return err
			}
		} else {
			log.Println("###### adding a batch")
			id, err := AddBatchToPartyProduct(ctx, instance, party.ID, bson.ObjectIdHex(ticket.PartyProductID))
			if err != nil {
				return err
			}
			pbatches = append(pbatches, types.ObjectPartyProduct{PartyProductID: bson.ObjectIdHex(id)})
		}
		batches = append(batches, instance)

	}
	after := types.Timestamp(time.Time{})
	before := types.Timestamp(time.Time{})
	if ticket.DoNotSellBefore != nil {
		before = *ticket.DoNotSellBefore
	}
	if ticket.DoNotSellAfter != nil {
		after = *ticket.DoNotSellAfter
	}
	master := types.MenuTicket{
		ID:                     bson.ObjectIdHex(ticket.MenuTicketID),
		Price:                  types.Price{Value: ticket.Price, CurrentIsoCode: "BRL"},
		ExhibitionName:         ticket.ExhibitionName,
		DoNotSellMoreThan:      ticket.DoNotSellMoreThan,
		DoNotSellBefore:        &before,
		DoNotSellAfter:         &after,
		Active:                 ticket.Active,
		GeneralInformation:     ticket.AdditionalInformations.General,
		FreeInformation:        ticket.AdditionalInformations.Free,
		AnniversaryInformation: ticket.AdditionalInformations.Anniversary,
		Batches:                batches,
	}
	err := EditMenuTicket(ctx, master, party, newname)
	if err != nil {
		return err
	}
	if bson.IsObjectIdHex(ticket.PartyProductID) {
		err = PatchTicketPartyProduct(ctx, ticket.PartyProductID, master, pbatches)
		if err != nil {
			return err
		}
		return nil
	}
	all := BrandNewSingle(master, party.ID)
	log.Println(all)
	err = InsertPartyProducts(ctx, all)
	if err != nil {
		return err
	}
	return nil
}

// AddBatchToPartyProduct TODO: NEEDS COMMENT INFO
func AddBatchToPartyProduct(ctx context.Context, instance types.MenuTicket, partyID bson.ObjectId, masterID bson.ObjectId) (string, error) {
	horario := types.Timestamp(time.Now())
	partyP := types.PartyProduct{
		ID:                bson.NewObjectId(),
		Name:              instance.ExhibitionName,
		CreationDate:      &horario,
		PartyID:           partyID,
		MoneyAmount:       instance.Price,
		DoNotSellMoreThan: instance.DoNotSellMoreThan,
		DoNotSellBefore:   instance.DoNotSellBefore,
		DoNotSellAfter:    instance.DoNotSellAfter,
		QuantityPurchased: 0.0,
		Type:              "TICKET",
		QuantityFree:      0,
		MenuTicketID:      &instance.ID,
		OwnerBatchID:      &masterID,
		AcceptsEvaluation: false,
		Featured:          false,
		Status:            PartyProductStatus(instance.Active),
		// Product
	}
	repo, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return "", err
	}
	err := repo.Collection.Insert(partyP)
	if err != nil {
		return "", err
	}
	log.Println(partyP)
	return partyP.ID.Hex(), nil
}

// PatchTicketPartyProduct TODO: NEEDS COMMENT INFO
func PatchTicketPartyProduct(ctx context.Context, id string, instance types.MenuTicket, pbatches []types.ObjectPartyProduct) error {
	repo, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	partyP, err := repo.GetByID(id)
	if err != nil {
		return err
	}
	log.Println(partyP.Batches)
	// log.Println(pbatches)
	now := types.Timestamp(time.Now())
	after := types.Timestamp(time.Time{})
	before := types.Timestamp(time.Time{})
	if instance.DoNotSellBefore != nil {
		before = *instance.DoNotSellBefore
	}
	if instance.DoNotSellAfter != nil {
		after = *instance.DoNotSellAfter
	}
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate":             &now,
				"moneyAmount":            instance.Price,
				"exhibitionName":         instance.ExhibitionName,
				"doNotSellMoreThan":      instance.DoNotSellMoreThan,
				"doNotSellBefore":        &before,
				"doNotSellAfter":         &after,
				"active":                 instance.Active,
				"status":                 PartyProductStatus(instance.Active),
				"generalInformation":     instance.GeneralInformation,
				"freeInformation":        instance.FreeInformation,
				"anniversaryInformation": instance.AnniversaryInformation,
				"batches":                pbatches,
			}},
		ReturnNew: true,
	}
	patchedProduct := types.PartyProduct{}
	_, err = repo.Collection.Find(bson.M{"_id": partyP.ID}).Apply(change, &patchedProduct)
	if err != nil {
		return err
	}
	return nil
}
