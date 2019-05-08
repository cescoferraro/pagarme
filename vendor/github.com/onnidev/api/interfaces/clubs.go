package interfaces

import (
	"errors"
	"fmt"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ClubsRepo is a struc that hold a mongo collection
type ClubsRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewClubsCollection creates a new ClubsDAO
func NewClubsCollection(store *infra.MongoStore) (ClubsRepo, error) {
	repo := ClubsRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("club"),
	}
	return repo, nil
}

// AppGetByID Insert the user on a database
func (c *ClubsRepo) AppGetByID(id string) (types.AppClub, error) {
	result := types.AppClub{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.AppClub{}, err
		}
		return result, nil
	}
	return types.AppClub{}, fmt.Errorf("not a valid object id")
}

// GetByID Insert the user on a database
func (c *ClubsRepo) GetByID(id string) (types.Club, error) {
	result := types.Club{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.Club{}, err
		}
		return result, nil
	}
	return types.Club{}, fmt.Errorf("not a valid object id")
}

// List lists a token collection
func (c *ClubsRepo) ListActive() ([]types.Club, error) {
	result := []types.Club{}
	err := c.Collection.Find(bson.M{"status": "ACTIVE"}).Sort("name").All(&result)
	return result, err
}

// List lists a token collection
func (c *ClubsRepo) List() ([]types.Club, error) {
	result := []types.Club{}
	err := c.Collection.Find(bson.M{}).Sort("name").All(&result)
	return result, err
}

// MineClubs lists a token collection
func (c *ClubsRepo) MineClubs(user types.UserClub) ([]types.Club, error) {
	clubs := []types.Club{}
	if user.Profile == "ONNI" {
		clubs, err := c.List()
		if err != nil {
			return clubs, err
		}
		return clubs, nil
	}
	clubs, err := c.Mine(user.Clubs)
	if err != nil {
		return clubs, err
	}
	if user.Profile != "ONNI" {
		if len(clubs) == 0 {
			err := errors.New("usuaio não possui clubs ativos")
			return clubs, err
		}
		return clubs, nil
	}
	return clubs, nil
}

// Mine lists a token collection
func (c *ClubsRepo) Mine(clubs []bson.ObjectId) ([]types.Club, error) {
	result := []types.Club{}
	for _, club := range clubs {
		inner := []types.Club{}
		err := c.Collection.FindId(club).All(&inner)
		if err != nil {
			return result, err
		}
		result = append(result, inner...)
	}
	return result, nil
}
