package onni

import (
	"context"
	"log"

	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// PatchClubMenuProduct TODO: NEEDS COMMENT INFO
func PatchClubMenuProduct(ctx context.Context, req types.SoftPartyPostRequest, party types.Party) (*bson.ObjectId, error) {
	clubMenuProductID := new(bson.ObjectId)
	log.Println("=======================================================")
	log.Println("PRODUCTS ==============================================")
	log.Println("=======================================================")
	if req.ClubMenuProduct != nil {
		if req.ClubMenuProduct.ID != nil {
			if party.ClubMenuProductID != nil {
				id := *party.ClubMenuProductID
				if id.Hex() == *req.ClubMenuProduct.ID {
					log.Println("EDITANDO O MESMO MENU =================================")
					menu, err := UpdateCurrentClubMenuProductAndPartyProducts(ctx, party, req)
					if err != nil {
						return clubMenuProductID, err
					}
					clubMenuProductID = &menu.ID
				}
			}
		}
	}
	if req.ClubMenuProduct != nil {
		if req.ClubMenuProduct.ID != nil {
			if party.ClubMenuProductID != nil {
				id := *party.ClubMenuProductID
				if id.Hex() != *req.ClubMenuProduct.ID {
					log.Println("TRODA DE MENU POR MENU EXISTENTE ======================")
					// ..asd
					menu, err := DeprecatePartyProductAndUpdateClubMenuProduct(ctx, party, req)
					if err != nil {
						return clubMenuProductID, err
					}
					clubMenuProductID = &menu.ID
				}
			}
		}
	}
	if req.ClubMenuProduct != nil {
		if req.ClubMenuProduct.ID != nil {
			if party.ClubMenuProductID == nil {
				log.Println("FESTA SEM PRODUCTS - TROCA DE MENU POR MENU EXISTENTE ======================")
				menu, err := DeprecatePartyProductAndUpdateClubMenuProduct(ctx, party, req)
				if err != nil {
					return clubMenuProductID, err
				}
				clubMenuProductID = &menu.ID
			}
		}
	}
	if req.ClubMenuProduct != nil {
		if req.ClubMenuProduct.ID == nil {
			log.Println("CRIANDO NOVO MENU PELO FRONT ==============================")
			menu, err := DeprecatePartyProductAndCreateClubMenuProduct(ctx, party, req)
			if err != nil {
				return clubMenuProductID, err
			}
			clubMenuProductID = &menu.ID
		}
	}
	if req.ClubMenuProduct == nil {
		log.Println("ZERANDO O MENU==============================")
		err := DeprecatePartyProductsByType(ctx, party.ID, "DRINK")
		if err != nil {
			return clubMenuProductID, err
		}
	}

	if clubMenuProductID != nil {
		id := *clubMenuProductID
		log.Println("PatchClubMenuProduct return a menu", id, id.Hex())
	}
	return clubMenuProductID, nil
}
