package onni

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// UserClub siksjdnf
func UserClubRepo(ctx context.Context, id string) (types.UserClub, error) {
	var userclub types.UserClub
	collection, ok := ctx.Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	if !ok {
		err := errors.New("bug")
		return userclub, err
	}
	userclub, err := collection.GetByID(id)
	if err != nil {
		return userclub, err
	}
	return userclub, nil
}

// SoftLoginUserClub TODO: NEEDS COMMENT INFO
func SoftLoginUserClub(ctx context.Context) (types.UserClub, error) {
	var user types.UserClub
	collection, ok := ctx.Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	if !ok {
		err := errors.New("bug")
		return user, err
	}
	loginRequest, ok := ctx.Value(middlewares.SoftUserClubLoginReq).(types.SoftLoginRequest)
	if !ok {
		err := errors.New("bug")
		return user, err
	}
	email := strings.ToLower(loginRequest.Email)
	user, err := collection.Login(email, loginRequest.Password)
	if err != nil {
		if err.Error() == "not found" {
			err := errors.New("user.club.login.error.login.failed")
			return user, err
		}
		return user, err
	}
	if user.Profile == "ATTENDANT" {
		err := errors.New("attendente não loga no backoffice")
		return user, err
	}
	return user, nil

}

// LoginUserClub TODO: NEEDS COMMENT INFO
func LoginUserClub(ctx context.Context) (types.UserClub, error) {
	var user types.UserClub
	collection, ok := ctx.Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	if !ok {
		err := errors.New("bug")
		return user, err
	}
	loginRequest, ok := ctx.Value(middlewares.UserClubLoginReq).(types.LoginRequest)
	if !ok {
		err := errors.New("bug")
		return user, err
	}
	log.Println("got ")
	email := strings.ToLower(loginRequest.Email)
	user, err := collection.Login(email, loginRequest.Password)
	if err != nil {
		if err.Error() == "not found" {
			err := errors.New("user.club.login.error.login.failed")
			return user, err
		}
		return user, err
	}
	if user.Status == "INACTIVE" {
		err := errors.New("user.club.login.error.status.inactive")
		return user, err
	}
	if user.Profile == "ATTENDANT" {
		err := errors.New("attendente não loga no backoffice")
		return user, err
	}
	return user, nil

}

// SoftResponse sjk sdfkjnsdf
func SoftResponse(ctx context.Context, user types.UserClub, clubs []types.Club) (types.SoftUserClubLoginResponse, error) {
	token, err := user.GenerateToken()
	if err != nil {
		return types.SoftUserClubLoginResponse{}, err
	}

	id := ""
	club := types.Club{}
	if user.Favorite != nil {
		fav := *user.Favorite
		id = fav.Hex()
		club, err = THISClub(id, clubs)
		if err != nil {
			return types.SoftUserClubLoginResponse{}, err
		}
	}
	if id == "" {
		rand.Seed(time.Now().UTC().UnixNano())
		id = clubs[rand.Intn(len(clubs))].ID.Hex()
	}
	var sclubs []string
	for _, club := range clubs {
		sclubs = append(sclubs, club.ID.Hex())
	}
	percent, err := Club(ctx, id)
	if err != nil {
		return types.SoftUserClubLoginResponse{}, nil
	}
	return types.SoftUserClubLoginResponse{
		ID:                user.ID,
		Token:             token,
		Name:              user.Name,
		Mail:              user.Mail,
		Tags:              club.Tags,
		Profile:           user.Profile,
		Clubs:             sclubs,
		ClubPercentTicket: percent.PercentTicket,
		ClubPercentDrink:  percent.PercentDrink,
		ClubID:            id,
	}, nil
}

// THISClub TODO: NEEDS COMMENT INFO
func THISClub(id string, clubs []types.Club) (types.Club, error) {
	for _, club := range clubs {
		if club.ID.Hex() == id {
			return club, nil
		}
	}
	err := errors.New("club not found onni boy")
	return types.Club{}, err
}

// Response sjk sdfkjnsdf
func Response(ctx context.Context, user types.UserClub, clubs []types.Club) (types.UserClubLoginResponse, error) {
	token, err := user.GenerateToken()
	if err != nil {
		return types.UserClubLoginResponse{}, nil
	}
	id := ""
	if user.Favorite != nil {
		fav := *user.Favorite
		id = fav.Hex()
	}
	if id == "" {
		rand.Seed(time.Now().UTC().UnixNano())
		id = clubs[rand.Intn(len(clubs))].ID.Hex()
	}
	return types.UserClubLoginResponse{
		ID:      user.ID,
		Token:   token,
		Name:    user.Name,
		Mail:    user.Mail,
		Profile: user.Profile,
		Clubs:   clubs,
		ClubID:  id,
	}, nil
}
