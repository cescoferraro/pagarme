package interfaces

import (
	"fmt"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NotificationRepo is a struc that hold a mongo collection
type NotificationRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewNotificationsCollection creates a new CardDAO
func NewNotificationsCollection(store *infra.MongoStore) (NotificationRepo, error) {
	repo := NotificationRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("notification"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *NotificationRepo) GetByID(id string) (types.Notification, error) {
	result := types.Notification{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.Notification{}, err
		}
		return result, nil
	}
	return types.Notification{}, fmt.Errorf("not a valid object id")
}

// GetByParty Insert the user on a database
func (c *NotificationRepo) GetByParty(id string) ([]types.Notification, error) {
	result := []types.Notification{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.
			Pipe(includePartyClubFieldOnNotification(
				bson.M{"partyId": bson.ObjectIdHex(id)},
				bson.M{"startDate": -1},
			)).
			All(&result)
		if err != nil {
			return []types.Notification{}, err
		}
		return result, nil
	}
	return []types.Notification{}, fmt.Errorf("not a valid object id")
}

// GetByCustomer Insert the user on a database
func (c *NotificationRepo) GetByCustomer(id string) ([]types.Notification, error) {
	result := []types.Notification{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.
			Pipe(includePartyClubFieldOnNotification(
				bson.M{"customers": bson.ObjectIdHex(id)},
				bson.M{"startDate": -1},
			)).
			All(&result)
		if err != nil {
			return []types.Notification{}, err
		}
		return result, nil
	}
	return []types.Notification{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *NotificationRepo) List() ([]types.Notification, error) {
	result := []types.Notification{}
	fromDate := time.Now().Add(-time.Duration(100000) * time.Hour)
	err := c.Collection.Pipe(
		c.makeBannerComplete(
			bson.M{"creationDate": bson.M{"$gt": fromDate}},
			bson.M{"creationDate": -1}),
	).All(&result)
	return result, err
}

// Create Insert the card on a database
func (c *NotificationRepo) Create(ban types.Notification) error {
	err := c.Collection.Insert(ban)
	if err != nil {
		return err
	}
	return nil
}

func (c *NotificationRepo) makeBannerComplete(match bson.M, sort bson.M) []bson.M {
	var response []bson.M
	if len(match) != 0 {
		response = []bson.M{bson.M{"$match": match}}
	}
	response = append(response, []bson.M{
		bson.M{
			"$lookup": bson.M{
				"from":         "club",
				"localField":   "clubId",
				"foreignField": "_id",
				"as":           "club"}},
		bson.M{"$unwind": bson.M{
			"path": "$club",
			"preserveNullAndEmptyArrays": true,
		}},
		bson.M{
			"$lookup": bson.M{
				"from":         "party",
				"localField":   "partyId",
				"foreignField": "_id",
				"as":           "party",
			}},
		bson.M{"$unwind": bson.M{
			"path": "$party",
			"preserveNullAndEmptyArrays": true,
		}},
		bson.M{"$sort": sort},
	}...)
	return response
}

func includePartyClubFieldOnNotification(match, sort bson.M) []bson.M {
	return []bson.M{
		bson.M{"$match": match},
		bson.M{"$lookup": bson.M{
			"from":         "party",
			"localField":   "partyId",
			"foreignField": "_id",
			"as":           "party"}},
		bson.M{"$unwind": bson.M{
			"path": "$party",
			"preserveNullAndEmptyArrays": true}},
		bson.M{"$lookup": bson.M{
			"from":         "club",
			"localField":   "clubId",
			"foreignField": "_id",
			"as":           "club"}},
		bson.M{"$unwind": bson.M{
			"path": "$club",
			"preserveNullAndEmptyArrays": true}},
		bson.M{"$sort": sort},
	}
}
