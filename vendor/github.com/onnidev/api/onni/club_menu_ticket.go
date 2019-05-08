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

// ClubMenuTicket TODO: NEEDS COMMENT INFO
func ClubMenuTicket(ctx context.Context, id string) (types.ClubMenuTicket, error) {
	result := types.ClubMenuTicket{}
	repo, ok := ctx.Value(middlewares.ClubMenuTicketRepoKey).(interfaces.ClubMenuTicketRepo)
	if !ok {
		err := errors.New("assert")
		return result, err
	}
	result, err := repo.GetByID(id)
	if err != nil {

		return result, err
	}
	return result, nil
}

// MenuTicket TODO: NEEDS COMMENT INFO
func MenuTicket(ctx context.Context, partyID, id string) (types.MenuTicket, error) {
	repo, ok := ctx.Value(middlewares.ClubMenuTicketRepoKey).(interfaces.ClubMenuTicketRepo)
	if !ok {
		err := errors.New("assert")
		return types.MenuTicket{}, err
	}
	partyRepo, ok := ctx.Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	if !ok {
		err := errors.New("bug")
		return types.MenuTicket{}, err
	}
	party, err := partyRepo.GetByID(partyID)
	if err != nil {
		return types.MenuTicket{}, err
	}
	result, err := repo.GetByClub(party.ClubID.Hex())
	if err != nil {
		return types.MenuTicket{}, err
	}
	for _, menu := range result {
		log.Println("######", menu.Name)
		for _, ticket := range menu.Tickets {
			if ticket.ID.Hex() == id {
				return ticket, nil
			}
		}
	}
	return types.MenuTicket{}, errors.New("not found")
}

// CreateClubMenuTicket NEEDS COMMENT INFO
func CreateClubMenuTicket(ctx context.Context, req types.SoftPartyPostRequest) (types.ClubMenuTicket, error) {
	result := types.ClubMenuTicket{}
	repo, ok := ctx.Value(middlewares.ClubMenuTicketRepoKey).(interfaces.ClubMenuTicketRepo)
	if !ok {
		err := errors.New("assert")
		return result, err
	}
	log.Println("clulb", req.ClubID)
	now := types.Timestamp(time.Now())
	zero := types.Timestamp(time.Time{})
	menu := types.ClubMenuTicket{
		ID:           bson.NewObjectId(),
		CreationDate: &now,
		UpdateDate:   &zero,
		ClubID:       bson.ObjectIdHex(req.ClubID),
		Name:         req.ClubMenuTicket.Name,
		MenuDefault:  false,
		Status:       "ACTIVE",
		Tickets:      SoftTicketPartyProductConverter(req.Tickets),
	}
	err := repo.Collection.Insert(menu)
	if err != nil {
		return result, err
	}
	return menu, nil
}

// UpdateClubMenuTicket TODO: NEEDS COMMENT INFO
func UpdateClubMenuTicket(ctx context.Context, req types.SoftPartyPostRequest) (types.ClubMenuTicket, error) {
	menuid := bson.ObjectIdHex(*req.ClubMenuTicket.ID)
	result := types.ClubMenuTicket{}
	repo, ok := ctx.Value(middlewares.ClubMenuTicketRepoKey).(interfaces.ClubMenuTicketRepo)
	if !ok {
		err := errors.New("assert")
		return result, err
	}
	result, err := repo.GetByID(menuid.Hex())
	if err != nil {
		return result, err
	}
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate": &now,
				"tickets":    SoftTicketPartyProductConverter(req.Tickets),
				"name":       req.ClubMenuTicket.Name,
			}},
		ReturnNew: true,
	}
	patchedMenu := types.ClubMenuTicket{}
	_, err = repo.Collection.Find(bson.M{"_id": result.ID}).Apply(change, &patchedMenu)
	if err != nil {
		return patchedMenu, err
	}
	log.Printf("PATCHEDF MENU HAS %v TICKETS\n", len(patchedMenu.Tickets))
	return patchedMenu, nil
}

// EditMenuTicket TODO: NEEDS COMMENT INFO
func EditMenuTicket(ctx context.Context, ticket types.MenuTicket, party types.Party, newname string) error {
	result := types.ClubMenuTicket{}
	repo, ok := ctx.Value(middlewares.ClubMenuTicketRepoKey).(interfaces.ClubMenuTicketRepo)
	if !ok {
		err := errors.New("assert")
		return err
	}
	menuID := *party.ClubMenuTicketID
	result, err := repo.GetByID(menuID.Hex())
	if err != nil {
		return err
	}
	now := types.Timestamp(time.Now())
	tickets := []types.MenuTicket{}
	for _, tk := range result.Tickets {
		if tk.ID.Hex() == ticket.ID.Hex() {
			tickets = append(tickets, ticket)
			continue
		}
		tickets = append(tickets, tk)

	}
	log.Println("######### the new name is ", newname)
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate": &now,
				"tickets":    tickets,
				"name":       newname,
			}},
		ReturnNew: true,
	}
	patchedMenu := types.ClubMenuTicket{}
	_, err = repo.Collection.Find(bson.M{"_id": result.ID}).Apply(change, &patchedMenu)
	if err != nil {
		return err
	}
	return nil
}

// AddMenuTicket TODO: NEEDS COMMENT INFO
func AddMenuTicket(ctx context.Context, ticket types.SoftTicketPostPartyProduct, party types.Party, req types.SoftPartyPostRequest) error {
	log.Println("##### CREATING menuTicket")
	result := types.ClubMenuTicket{}
	repo, ok := ctx.Value(middlewares.ClubMenuTicketRepoKey).(interfaces.ClubMenuTicketRepo)
	if !ok {
		err := errors.New("assert")
		return err
	}
	menuID := *party.ClubMenuTicketID
	result, err := repo.GetByID(menuID.Hex())
	if err != nil {
		return err
	}
	now := types.Timestamp(time.Now())
	menuticket := SoftTicketPartyProductSingleConverter(ticket)
	set := bson.M{
		"updateDate": &now,
		"tickets":    append(result.Tickets, menuticket),
	}
	if req.ClubMenuProduct != nil {
		menu := *req.ClubMenuProduct
		set["name"] = menu.Name
	}
	change := mgo.Change{Update: bson.M{"$set": set}, ReturnNew: true}
	patchedMenu := types.ClubMenuTicket{}
	_, err = repo.Collection.Find(bson.M{"_id": result.ID}).Apply(change, &patchedMenu)
	if err != nil {
		return err
	}
	all := BrandNewSingle(menuticket, party.ID)
	log.Println(all)
	err = InsertPartyProducts(ctx, all)
	if err != nil {
		return err
	}
	log.Println("CESCO:", len(patchedMenu.Tickets))
	return nil
}
