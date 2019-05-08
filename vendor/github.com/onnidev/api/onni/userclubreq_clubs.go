package onni

import (
	"context"
	"errors"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// Clubs TODO: NEEDS COMMENT INFO
func Clubs(ctx context.Context, req types.UserClubPostRequest) ([]bson.ObjectId, error) {
	clubs := []bson.ObjectId{}
	clubCollection, ok := ctx.Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	if !ok {
		err := errors.New("assert bug")
		return clubs, err
	}
	for _, id := range req.Clubs {
		_, err := clubCollection.GetByID(id)
		if err != nil {
			return clubs, err
		}
		clubs = append(clubs, bson.ObjectIdHex(id))
	}
	return clubs, nil
}

// UserClub TODO: NEEDS COMMENT INFO
func UserClub(ctx context.Context, id string) (types.UserClub, error) {
	userClubCollection, ok := ctx.Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	if !ok {
		err := errors.New("assert bug")
		return types.UserClub{}, err
	}
	userClub, err := userClubCollection.GetByID(id)
	if err != nil {
		return types.UserClub{}, err
	}
	return userClub, nil
}

// CreateUserClub TODO: NEEDS COMMENT INFO
func CreateUserClub(ctx context.Context, password string) (types.UserClub, error) {
	req, ok := ctx.Value(middlewares.ReadUserClubPostKey).(types.UserClubPostRequest)
	if !ok {
		err := errors.New("assert bug")
		return types.UserClub{}, err
	}
	horario := types.Timestamp(time.Now())
	id := bson.NewObjectId()
	if req.ID != "" {
		id = bson.ObjectIdHex(req.ID)
	}
	clubsids, err := Clubs(ctx, req)
	if err != nil {
		return types.UserClub{}, err
	}
	mail := shared.NormalizeEmail(req.Email)
	fav := new(bson.ObjectId)
	if len(req.Clubs) == 1 {
		id := bson.ObjectIdHex(req.Clubs[0])
		fav = &id
	}
	userClub := types.UserClub{
		ID:           id,
		CreationDate: &horario,
		Favorite:     fav,
		Name:         req.Name,
		Mail:         mail,
		Status:       "ACTIVE",
		Password:     shared.EncryptPassword2(password),
		Profile:      req.Profile,
		Tags:         req.Tags,
		Clubs:        clubsids,
	}
	return userClub, nil
}

// PersistUserClub sdkjfn
func PersistUserClub(ctx context.Context, userClub types.UserClub) error {
	userClubCollection, ok := ctx.Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	if !ok {
		err := errors.New("assert bug")
		return err
	}
	err := userClubCollection.Collection.Insert(userClub)
	if err != nil {
		return err
	}
	return nil
}
