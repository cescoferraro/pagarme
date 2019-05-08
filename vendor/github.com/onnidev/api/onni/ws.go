package onni

import (
	"context"
	"errors"
	"log"

	"github.com/bradfitz/slice"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// HTTPPartyReport TODO: NEEDS COMMENT INFO
func HTTPPartyReport(ctx context.Context, userClubID string) ([]types.WebSocketReport, error) {
	var reports []types.WebSocketReport
	voucherRepo, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("voucherRepo")
		return reports, err
	}
	partyCollection, ok := ctx.Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	if !ok {
		err := errors.New("PartiesRepo")
		return reports, err
	}
	clubsRepo, ok := ctx.Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	if !ok {
		err := errors.New("ClubsRepo")
		return reports, err
	}

	userClubRepo, ok := ctx.Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	if !ok {
		err := errors.New("UserClubRepo")
		return reports, err
	}
	return WSGetPartyReport(userClubID, voucherRepo, partyCollection, clubsRepo, userClubRepo, false)
}

// WSGetPartyReport TODO: NEEDS COMMENT INFO
func WSGetPartyReport(userClubID string,
	voucherRepo interfaces.VouchersRepo,
	partyCollection interfaces.PartiesRepo,
	clubsRepo interfaces.ClubsRepo,
	userClubRepo interfaces.UserClubRepo,
	fetchVouchers bool) ([]types.WebSocketReport, error) {
	var reports []types.WebSocketReport
	clubUser, err := userClubRepo.GetByID(userClubID)
	if err != nil {
		return reports, err
	}
	var clubs []types.Club
	var allParties []types.Party
	if clubUser.Profile == "ONNI" {
		allParties, err = partyCollection.List()
		if err != nil {
			log.Println(err.Error())
			return reports, err
		}
	} else {
		clubs, err = clubsRepo.Mine(clubUser.Clubs)
		if err != nil {
			return reports, err
		}
		allParties, err = partyCollection.GetByClubs(clubs)
		if err != nil {
			log.Println(err.Error())
			return reports, err
		}
	}
	log.Println(len(allParties))
	for _, party := range allParties {
		if fetchVouchers {
			vouchers, err := voucherRepo.GetByParty(party.ID.Hex())
			if err != nil {
				return reports, nil
			}
			log.Println(len(vouchers))
			list := types.VouchersList(vouchers)
			reports = append(reports, types.WebSocketReport{
				PartyID:        party.ID.Hex(),
				PartyName:      party.Name,
				PartyAddress:   party.Address,
				PartyStartDate: party.StartDate,
				PartyEndDate:   party.EndDate,
				ClubID:         party.ClubID.Hex(),
				ClubName:       party.Club.Name,
				Used:           list.FilterByStatus("USED").Size(),
				Available:      list.FilterByStatus("AVAILABLE").Size(),
			})
			continue
		}

		reports = append(reports, types.WebSocketReport{
			PartyID:        party.ID.Hex(),
			PartyName:      party.Name,
			PartyAddress:   party.Address,
			PartyStartDate: party.StartDate,
			PartyEndDate:   party.EndDate,
			ClubID:         party.ClubID.Hex(),
			ClubName:       party.Club.Name,
			Used:           0,
			Available:      0,
		})
	}

	slice.Sort(reports[:], func(i, j int) bool {
		return (reports[i].PartyStartDate.Time().Before(reports[j].PartyStartDate.Time()))
	})
	if len(reports) == 0 {
		return []types.WebSocketReport{}, nil
	}
	return reports, nil
}
