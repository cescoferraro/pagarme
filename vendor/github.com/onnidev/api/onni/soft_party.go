package onni

import (
	"context"
	"errors"
	"log"

	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// SoftParty TODO: NEEDS COMMENT INFO
func SoftParty(ctx context.Context, party types.Party) (types.SoftParty, error) {
	log.Println("&&&&&&&&&&&&&&&&&&&&&&&&&")
	log.Println("&&&&&&&&&&&&&&&&&&&&&&&&&")
	log.Println("entrando no endpoint")
	log.Println("&&&&&&&&&&&&&&&&&&&&&&&&&")
	log.Println("&&&&&&&&&&&&&&&&&&&&&&&&&")
	log.Println("&&&&&&&&&&&&&&&&&&&&&&&&&")
	softparty := types.SoftParty{}
	softticket, err := TicketsFromParty(ctx, party.ID.Hex())
	if err != nil {
		return softparty, err
	}
	log.Println("=========================")
	log.Println("=========================")
	log.Println("=========================")
	log.Println("=========================")
	log.Println("=========================")
	log.Println("=========================")
	log.Println("=========================")
	log.Println("=========================")
	softproducts, err := DrinksFromParty(ctx, party.ID.Hex())
	if err != nil {
		return softparty, err
	}
	log.Println("-------------------------")
	log.Println("-------------------------")
	log.Println("-------------------------")
	log.Println("-------------------------")
	log.Println("-------------------------")
	log.Println("-------------------------")
	musicstyles := []string{}
	for _, style := range party.MusicStyles {
		musicstyles = append(musicstyles, style.Name)
	}
	softparty = types.SoftParty{
		ID:               party.ID.Hex(),
		StartDate:        party.StartDate,
		EndDate:          party.EndDate,
		Name:             party.Name,
		ClubID:           party.ClubID,
		Description:      party.Description,
		Status:           party.Status,
		Address:          party.Address,
		MainAttraction:   party.MainAttraction,
		BackgroundImage:  party.BackgroundImage,
		AssumeServiceFee: party.AssumeServiceFee,
		OtherAttractions: party.OtherAttractions,
		MusicStyles:      musicstyles,
		Tickets:          softticket,
		Products:         softproducts,
	}
	if party.ClubMenuTicketID != nil {
		id := *party.ClubMenuTicketID
		if bson.IsObjectIdHex(id.Hex()) {
			log.Println(id.Hex(), "****^")
			all, err := ClubMenuTicket(ctx, id.Hex())
			if err != nil {
				return softparty, errors.New("TICKET:" + err.Error())
			}
			idd := all.ID.Hex()
			idname := types.IDName{ID: &idd, Name: all.Name}
			softparty.ClubMenuTicket = &idname
		}
	}
	log.Println("((((((((()))))))))")
	log.Println(party.ClubMenuProductID)
	if party.ClubMenuProductID != nil {
		id := *party.ClubMenuProductID
		if id.Hex() != "" {
			log.Println("****^")
			log.Println(id.Hex(), "****^")
			all, err := ClubMenuProduct(ctx, id.Hex())
			if err != nil {
				return softparty, errors.New("PRODUCT:" + err.Error())
			}
			idd := all.ID.Hex()
			idname := types.IDName{ID: &idd, Name: all.Name}
			softparty.ClubMenuProduct = &idname
		}
	}
	return softparty, nil

}
