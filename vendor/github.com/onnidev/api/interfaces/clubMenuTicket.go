package interfaces

import (
	"fmt"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ClubMenuTicketRepo is a struc that hold a mongo collection
type ClubMenuTicketRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewClubMenuTicketCollection creates a new CardDAO
func NewClubMenuTicketCollection(store *infra.MongoStore) (ClubMenuTicketRepo, error) {
	repo := ClubMenuTicketRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("clubMenuTicket"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *ClubMenuTicketRepo) GetByID(id string) (types.ClubMenuTicket, error) {
	result := types.ClubMenuTicket{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.ClubMenuTicket{}, err
		}
		return result, nil
	}
	return types.ClubMenuTicket{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *ClubMenuTicketRepo) List() ([]types.ClubMenuTicket, error) {
	result := []types.ClubMenuTicket{}
	err := c.Collection.Find(bson.M{}).All(&result)
	return result, err
}

// GetByClub TODO: NEEDS COMMENT INFO
func (c *ClubMenuTicketRepo) GetByClub(id string) ([]types.ClubMenuTicket, error) {
	result := []types.ClubMenuTicket{}
	err := c.Collection.Find(bson.M{
		"clubId": bson.ObjectIdHex(id),
	}).All(&result)
	return result, err
}

// GetActivesByClub TODO: NEEDS COMMENT INFO
func (c *ClubMenuTicketRepo) GetActivesByClub(id string) ([]types.ClubMenuTicket, error) {
	result := []types.ClubMenuTicket{}
	err := c.Collection.Find(bson.M{
		"clubId": bson.ObjectIdHex(id),
		"status": "ACTIVE",
	}).All(&result)
	return result, err
}
