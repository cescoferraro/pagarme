package onni

import (
	"context"
	"errors"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// PersistClub sdkjfn
func PersistClub(ctx context.Context, club types.Club) error {
	repo, ok := ctx.Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	err := repo.Collection.Insert(club)
	if err != nil {
		return err
	}
	return nil
}

// Club siksjdnf
func Club(ctx context.Context, id string) (types.Club, error) {
	var club types.Club
	clubsCollection, ok := ctx.Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	if !ok {
		err := errors.New("bug")
		return club, err
	}
	club, err := clubsCollection.GetByID(id)
	if err != nil {
		return club, err
	}
	return club, nil
}

// UserClubClubs siksjdnf
func UserClubClubs(ctx context.Context, user types.UserClub) ([]types.Club, error) {
	var clubs []types.Club
	clubsCollection, ok := ctx.Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	if !ok {
		err := errors.New("bug")
		return clubs, err
	}
	clubs, err := clubsCollection.MineClubs(user)
	if err != nil {
		return clubs, err
	}
	return clubs, nil
}

// UserClubClubsIDS siksjdnf
func UserClubClubsIDS(ctx context.Context, user types.UserClub) ([]types.Club, error) {
	// var sclubs []string
	clubs, err := UserClubClubs(ctx, user)
	if err != nil {
		return clubs, err
	}
	// for _, club := range clubs {
	// 	sclubs = append(sclubs, club.ID.Hex())
	// }
	return clubs, nil
}
