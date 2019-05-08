package interfaces

import (
	"fmt"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ClubMenuProductRepo is a struc that hold a mongo collection
type ClubMenuProductRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewClubMenuProductCollection creates a new CardDAO
func NewClubMenuProductCollection(store *infra.MongoStore) (ClubMenuProductRepo, error) {
	repo := ClubMenuProductRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("clubMenuProduct"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *ClubMenuProductRepo) GetByID(id string) (types.ClubMenuProduct, error) {
	result := types.ClubMenuProduct{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.ClubMenuProduct{}, err
		}
		return result, nil
	}
	return types.ClubMenuProduct{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *ClubMenuProductRepo) List() ([]types.ClubMenuProduct, error) {
	result := []types.ClubMenuProduct{}
	err := c.Collection.Find(bson.M{}).All(&result)
	return result, err
}

// GetByClubActive TODO: NEEDS COMMENT INFO
func (c *ClubMenuProductRepo) GetByClubActive(id string) ([]types.ClubMenuProduct, error) {
	result := []types.ClubMenuProduct{}
	err := c.Collection.Find(bson.M{
		"clubId": bson.ObjectIdHex(id),
		"status": "ACTIVE",
	}).All(&result)
	return result, err
}

// GetByClub TODO: NEEDS COMMENT INFO
func (c *ClubMenuProductRepo) GetByClub(id string) ([]types.ClubMenuProduct, error) {
	result := []types.ClubMenuProduct{}
	err := c.Collection.Find(bson.M{
		"clubId": bson.ObjectIdHex(id),
	}).All(&result)
	return result, err
}
