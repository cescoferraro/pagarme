package interfaces

import (
	"fmt"
	"log"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ClubLeadRepo is a struc that hold a mongo collection
type ClubLeadRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewClubLeadCollection creates a new ClubLeadDAO
func NewClubLeadCollection(store *infra.MongoStore) (ClubLeadRepo, error) {
	repo := ClubLeadRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("clubLead"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *ClubLeadRepo) GetByID(id string) (types.ClubLead, error) {
	result := types.ClubLead{}
	log.Println(id)
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.ClubLead{}, err
		}
		return result, nil
	}
	return types.ClubLead{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *ClubLeadRepo) List() ([]types.ClubLead, error) {
	result := []types.ClubLead{}
	err := c.Collection.Pipe([]bson.M{}).All(&result)
	return result, err
}

// Create Insert the card on a database
func (c *ClubLeadRepo) Create(ban types.ClubLead) error {
	err := c.Collection.Insert(ban)
	if err != nil {
		return err
	}
	return nil
}

func includeCustomerField(match, sort bson.M) []bson.M {
	return []bson.M{
		bson.M{"$match": match},
		bson.M{"$lookup": bson.M{
			"from":         "customer",
			"localField":   "customerId",
			"foreignField": "_id",
			"as":           "customer"}},
		bson.M{"$unwind": bson.M{
			"path": "$customer",
			"preserveNullAndEmptyArrays": false}},
		bson.M{"$sort": sort},
	}
}
