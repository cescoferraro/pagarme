package interfaces

import (
	"fmt"
	"log"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BansRepo is a struc that hold a mongo collection
type BansRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// key is a context key
type key string

// NewBansCollection creates a new CardDAO
func NewBansCollection(store *infra.MongoStore) (BansRepo, error) {
	repo := BansRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("ban"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *BansRepo) GetByID(id string) (types.Ban, error) {
	result := types.Ban{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.Ban{}, err
		}
		return result, nil
	}
	return types.Ban{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *BansRepo) List() ([]types.Ban, error) {
	result := []types.Ban{}
	err := c.Collection.Find(bson.M{}).All(&result)
	return result, err
}

// ByTypeAndPayload Insert the user on a database
func (c *BansRepo) ByTypeAndPayload(ban types.Ban) (types.Ban, error) {
	log.Println(ban.Type, ban.Payload)
	result := types.Ban{}
	err := c.Collection.
		Pipe(c.makeComplete(bson.M{"type": ban.Type, "payload": ban.Payload}, bson.M{"creationDate": 1})).
		One(&result)
	return result, err
}

// ByTypeAndPayload Insert the user on a database
func (c *BansRepo) ByPayload(bans []types.Ban) ([]types.Ban, error) {
	payloads := []string{}
	for _, ban := range bans {
		if !shared.Contains(payloads, ban.Payload) {
			payloads = append(payloads, ban.Payload)
		}
	}
	log.Println(payloads)
	result := []types.Ban{}
	err := c.Collection.Pipe(c.makeComplete(
		bson.M{
			"payload": bson.M{"$in": payloads}},
		bson.M{"creationDate": 1},
	)).All(&result)
	return result, err
}

// IsCustomerBanned Insert the user on a database
func (c *BansRepo) IsCustomerBanned(customer string) (types.Ban, error) {
	result := types.Ban{}
	err := c.Collection.Pipe(c.makeComplete(
		bson.M{"type": "CUSTOMERID", "payload": customer},
		bson.M{"creationDate": 1},
	)).One(&result)
	return result, err
}

func (c *BansRepo) makeComplete(match bson.M, sort bson.M) []bson.M {
	all := []bson.M{
		bson.M{"$match": match},
		bson.M{"$sort": sort},
	}
	return all
}
