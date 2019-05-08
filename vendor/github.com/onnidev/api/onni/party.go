package onni

import (
	"context"
	"errors"
	"time"

	"github.com/bradfitz/slice"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
)

// PartyAndClub siksjdnf
func PartyAndClub(ctx context.Context, partyID string) (types.Party, types.Club, error) {
	party, err := Party(ctx, partyID)
	if err != nil {
		return types.Party{}, types.Club{}, err
	}
	club, err := Club(ctx, party.ClubID.Hex())
	if err != nil {
		return types.Party{}, types.Club{}, err
	}
	if club.Status == "INACTIVE" {
		err := errors.New("club.inactive")
		return party, club, err
	}
	return party, club, nil
}

// Party siksjdnf
func Party(ctx context.Context, partyID string) (types.Party, error) {
	var allParties types.Party
	partiesCollection, ok := ctx.Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	if !ok {
		err := errors.New("bug")
		return allParties, err
	}
	allParties, err := partiesCollection.GetByID(partyID)
	if err != nil {
		return allParties, err
	}
	return allParties, nil
}

// GetAppPartiesFromFilter siksjdnf
func GetAppPartiesFromFilter(ctx context.Context) ([]types.AppParty, error) {
	filters := ctx.Value(middlewares.PartyListFilterKey).(types.PartyFilter)
	partiesCollection := ctx.Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	var allParties []types.AppParty
	allParties, err := partiesCollection.AppFilteredList(filters)
	if err != nil {
		return allParties, err
	}
	var days []time.Time
	for _, party := range allParties {
		if !shared.ContainsTime(days, party.StartDate.Time()) {
			if party.Club.Status == "ACTIVE" {
				days = append(days, party.StartDate.Time())
			}
		}
	}
	slice.Sort(days[:], func(i, j int) bool {
		return (days[i].Before(days[j]))
	})
	result := []types.AppParty{}
	for _, time := range days {
		result = append(result, allAppPDate(filters, allParties, time)...)
	}
	return result, nil
}

// GetPartiesFromFilter siksjdnf
func GetPartiesFromFilter(ctx context.Context) ([]types.Party, error) {
	filters := ctx.Value(middlewares.PartyListFilterKey).(types.PartyFilter)
	partiesCollection := ctx.Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	var allParties []types.Party
	allParties, err := partiesCollection.FilteredList(filters)
	if err != nil {
		return allParties, err
	}
	var days []time.Time
	for _, party := range allParties {
		if !shared.ContainsTime(days, party.StartDate.Time()) {
			if party.Club.Status == "ACTIVE" {
				days = append(days, party.StartDate.Time())
			}
		}
	}
	slice.Sort(days[:], func(i, j int) bool {
		return (days[i].Before(days[j]))
	})
	result := []types.Party{}
	for _, time := range days {
		result = append(result, allPDate(filters, allParties, time)...)
	}
	return result, nil
}

func allPDate(filters types.PartyFilter, parties []types.Party, date time.Time) []types.Party {
	var result []types.Party
	for _, party := range parties {
		if shared.DateEqual(party.StartDate.Time(), date) {
			result = append(result, party)
		}
	}
	slice.Sort(result[:], func(i, j int) bool {
		return (result[i].Distance(filters.Long, filters.Lat) <
			result[j].Distance(filters.Long, filters.Lat))
	})
	return result
}
func allAppPDate(filters types.PartyFilter, parties []types.AppParty, date time.Time) []types.AppParty {
	result := []types.AppParty{}
	for _, party := range parties {
		if shared.DateEqual(party.StartDate.Time(), date) {
			result = append(result, party)
		}
	}
	slice.Sort(result[:], func(i, j int) bool {
		return (result[i].Distance(filters.Long, filters.Lat) <
			result[j].Distance(filters.Long, filters.Lat))
	})
	return result
}
