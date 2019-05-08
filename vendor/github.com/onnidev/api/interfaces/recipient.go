package interfaces

import (
	"fmt"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RecipientsRepo is a struc that hold a mongo collection
type RecipientsRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewRecipientCollection creates a new TokenDAO
func NewRecipientCollection(store *infra.MongoStore) (RecipientsRepo, error) {
	repo := RecipientsRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("pagarMeRecipient"),
	}
	return repo, nil
}

// GetByClubID Insert the user on a database
func (c *RecipientsRepo) GetByClubID(token string) ([]types.Recipient, error) {
	result := []types.Recipient{}
	err := c.Collection.Find(bson.M{"clubId": bson.ObjectIdHex(token)}).All(&result)
	return result, err
}

// Dull Insert the user on a database
func (c *RecipientsRepo) Dull(token string) error {
	result := types.Recipient{}
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"status": "INACTIVE",
			}},
		ReturnNew: true,
	}
	_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(token)}).Apply(change, &result)
	return err
}

// GetByToken Insert the user on a database
func (c *RecipientsRepo) GetByToken(token string) (types.Recipient, error) {
	result := types.Recipient{}
	err := c.Collection.Find(bson.M{"recipientId": token}).One(&result)
	return result, err
}

// GetByID Insert the user on a database
func (c *RecipientsRepo) GetByID(id string) (types.Recipient, error) {
	result := types.Recipient{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe(
			[]bson.M{
				bson.M{"$match": bson.M{
					"_id": bson.ObjectIdHex(id),
				}},
				bson.M{
					"$lookup": bson.M{
						"from":         "club",
						"localField":   "clubId",
						"foreignField": "_id",
						"as":           "club"}},
				bson.M{"$unwind": bson.M{
					"path": "$club",
					"preserveNullAndEmptyArrays": true}},
			}).One(&result)
		if err != nil {
			return types.Recipient{}, err
		}
		return result, nil
	}
	return types.Recipient{}, fmt.Errorf("not a valid object id")
}

// List lists a token collection
func (c *RecipientsRepo) List() ([]types.Recipient, error) {
	result := []types.Recipient{}
	err := c.Collection.Find(bson.M{}).All(&result)
	return result, err
}

// ListFull Insert the user on a database
func (c *RecipientsRepo) ListFull() ([]types.Recipient, error) {
	result := []types.Recipient{}
	err := c.Collection.Pipe(
		[]bson.M{
			bson.M{
				"$lookup": bson.M{
					"from":         "club",
					"localField":   "clubId",
					"foreignField": "_id",
					"as":           "club"}},
			bson.M{"$unwind": bson.M{
				"path": "$club",
				"preserveNullAndEmptyArrays": true}},
		}).All(&result)
	if err != nil {
		return []types.Recipient{}, err
	}
	return result, nil
}
