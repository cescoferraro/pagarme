package interfaces

import (
	"fmt"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AntiTheftRepo is a struc that hold a mongo collection
type AntiTheftRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewAntiTheftCollection creates a new CardDAO
func NewAntiTheftCollection(store *infra.MongoStore) (AntiTheftRepo, error) {
	repo := AntiTheftRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("firewall"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *AntiTheftRepo) GetByID(id string) (types.AntiTheftResult, error) {
	result := types.AntiTheftResult{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.AntiTheftResult{}, err
		}
		return result, nil
	}
	return types.AntiTheftResult{}, fmt.Errorf("not a valid object id")
}

// PendingAntithefts Insert the user on a database
func (c *AntiTheftRepo) PendingAntithefts() ([]types.AntiTheftResult, error) {
	result := []types.AntiTheftResult{}
	err := c.Collection.Pipe(c.makeComplete(bson.M{}, bson.M{"creationDate": -1}, 200)).All(&result)
	if err != nil {
		return []types.AntiTheftResult{}, err
	}
	return result, nil
}

func (c *AntiTheftRepo) makeComplete(match bson.M, sort bson.M, limit int) []bson.M {
	all := []bson.M{
		bson.M{"$match": match},
		bson.M{
			"$lookup": bson.M{
				"from":         "club",
				"localField":   "clubId",
				"foreignField": "_id",
				"as":           "club"}},
		bson.M{"$unwind": bson.M{
			"path": "$club",
			"preserveNullAndEmptyArrays": true}},
		bson.M{
			"$lookup": bson.M{
				"from":         "party",
				"localField":   "partyId",
				"foreignField": "_id",
				"as":           "party"}},
		bson.M{"$unwind": bson.M{
			"path": "$party",
			"preserveNullAndEmptyArrays": true}},
		bson.M{
			"$lookup": bson.M{
				"from":         "customer",
				"localField":   "customerId",
				"foreignField": "_id",
				"as":           "customer"},
		},
		bson.M{"$unwind": bson.M{
			"path": "$customer",
			"preserveNullAndEmptyArrays": true}},
		bson.M{"$sort": sort},
		bson.M{"$limit": limit},
	}
	return all
}
