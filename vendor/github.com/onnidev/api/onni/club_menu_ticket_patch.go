package onni

import (
	"context"
	"log"

	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// PatchClubMenuTicket TODO: NEEDS COMMENT INFO
func PatchClubMenuTicket(ctx context.Context, req types.SoftPartyPostRequest, party types.Party) (*bson.ObjectId, error) {
	clubMenuTicketID := new(bson.ObjectId)
	log.Println("=======================================================")
	log.Println("TICKETSSS =============================================")
	log.Println("=======================================================")
	if req.ClubMenuTicket != nil {
		if req.ClubMenuTicket.ID != nil {
			if party.ClubMenuTicketID != nil {
				id := *party.ClubMenuTicketID
				if id.Hex() == *req.ClubMenuTicket.ID {
					log.Println("EDITANDO O MESMO MENU =================================")
					menu, err := UpdateCurrentClubMenuTicketAndPartyProducts(ctx, party, req)
					if err != nil {
						return clubMenuTicketID, err
					}
					clubMenuTicketID = &menu.ID
				}
			}
		}
	}
	if req.ClubMenuTicket != nil {
		if req.ClubMenuTicket.ID != nil {
			if party.ClubMenuTicketID != nil {
				id := *party.ClubMenuTicketID
				if id.Hex() != *req.ClubMenuTicket.ID {
					log.Println("TROCA DE MENU POR MENU EXISTENTE ======================")
					menu, err := DeprecatePartyProductAndUpdateClubMenuTicket(ctx, party, req)
					if err != nil {
						return clubMenuTicketID, err
					}
					clubMenuTicketID = &menu.ID
				}
			}
		}
	}
	if req.ClubMenuTicket != nil {
		if req.ClubMenuTicket.ID != nil {
			if party.ClubMenuTicketID == nil {
				log.Println("FESTA SEM TICKETS - TROCA DE MENU POR MENU EXISTENTE ======================")
				menu, err := DeprecatePartyProductAndUpdateClubMenuTicket(ctx, party, req)
				if err != nil {
					return clubMenuTicketID, err
				}
				clubMenuTicketID = &menu.ID
			}
		}
	}
	if req.ClubMenuTicket != nil {
		if req.ClubMenuTicket.ID == nil {
			log.Println("CRIANDO NOVO MENU PELO FRONT ==============================")
			menu, err := DeprecatePartyProductAndCreateClubMenuTicket(ctx, party, req)
			if err != nil {
				return clubMenuTicketID, err
			}
			clubMenuTicketID = &menu.ID
		}
	}
	if req.ClubMenuTicket == nil {
		log.Println("ZERANDO O MENU==============================")
		err := DeprecatePartyProductsByType(ctx, party.ID, "TICKET")
		if err != nil {
			return clubMenuTicketID, err
		}
	}
	if clubMenuTicketID != nil {
		id := *clubMenuTicketID
		log.Println("PatchClubMenuTicket return a menu", id, id.Hex())
	}
	return clubMenuTicketID, nil
}
