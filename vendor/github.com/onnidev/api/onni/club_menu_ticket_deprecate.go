package onni

import (
	"context"
	"errors"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
)

// DeprecateUnusedTicket TODO: NEEDS COMMENT INFO
func DeprecateUnusedTicket(ctx context.Context, party types.Party, req types.SoftPartyPostRequest) error {
	partyPRepo, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	stillused := []string{}
	for _, ticket := range req.Tickets {
		if ticket.PartyProductID != "" {
			stillused = append(stillused, ticket.PartyProductID)
			for _, batch := range ticket.Batches {
				if batch.PartyProductID != "" {
					stillused = append(stillused, batch.PartyProductID)
				}
			}
		}
	}
	current, err := partyPRepo.GetByPartyIDAndType(party.ID.Hex(), "TICKET")
	if err != nil {
		return err
	}
	obsolete := []string{}
	for _, product := range current {
		if !shared.Contains(stillused, product.ID.Hex()) {
			// log.Println("product ", product.Name)
			obsolete = append(obsolete, product.ID.Hex())
		}
	}
	err = DeprecatePartyProductsByID(ctx, obsolete)
	if err != nil {
		return err
	}
	return nil
}

// DeprecatePartyProductAndUpdateClubMenuTicket TODO: NEEDS COMMENT INFO
func DeprecatePartyProductAndUpdateClubMenuTicket(ctx context.Context, party types.Party, req types.SoftPartyPostRequest) (types.ClubMenuTicket, error) {
	result := types.ClubMenuTicket{}
	err := DeprecatePartyProductsByType(ctx, party.ID, "TICKET")
	if err != nil {
		return result, err
	}
	menu, err := UpdateClubMenuTicket(ctx, req)
	if err != nil {
		return result, err
	}
	_, err = BrandNewTicketsFromMenu(ctx, menu, party.ID)
	if err != nil {
		return result, err
	}
	return menu, nil
}

// DeprecatePartyProductAndCreateClubMenuTicket TODO: NEEDS COMMENT INFO
func DeprecatePartyProductAndCreateClubMenuTicket(ctx context.Context, party types.Party, req types.SoftPartyPostRequest) (types.ClubMenuTicket, error) {
	result := types.ClubMenuTicket{}
	err := DeprecatePartyProductsByType(ctx, party.ID, "TICKET")
	if err != nil {
		return result, err
	}
	menu, err := CreateClubMenuTicket(ctx, req)
	if err != nil {
		return result, err
	}
	_, err = BrandNewTicketsFromMenu(ctx, menu, party.ID)
	if err != nil {
		return result, err
	}
	return menu, nil
}
