package interfaces

import (
	"fmt"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BannerRepo is a struc that hold a mongo collection
type BannerRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewBannerCollection creates a new CardDAO
func NewBannerCollection(store *infra.MongoStore) (BannerRepo, error) {
	repo := BannerRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("banner"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *BannerRepo) GetByID(id string) (types.Banner, error) {
	result := types.Banner{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe(
			c.makeBannerComplete(
				bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"creationDate": 1})).One(&result)
		if err != nil {
			return types.Banner{}, err
		}
		return result, nil
	}
	return types.Banner{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *BannerRepo) List() ([]types.Banner, error) {
	result := []types.Banner{}
	fromDate := time.Now().Add(-time.Duration(100000) * time.Hour)
	err := c.Collection.Pipe(
		c.makeBannerComplete(
			bson.M{"creationDate": bson.M{"$gt": fromDate}},
			bson.M{"creationDate": -1}),
	).All(&result)
	return result, err
}

// Create Insert the card on a database
func (c *BannerRepo) Create(ban types.Banner) error {
	err := c.Collection.Insert(ban)
	if err != nil {
		return err
	}
	return nil
}

// GetPublishedBanners Insert the user on a database
func (c *BannerRepo) AllBanners() ([]types.Banner, error) {
	result := []types.Banner{}
	err := c.Collection.Pipe(
		c.makeBannerComplete(
			bson.M{},
			bson.M{"creationDate": 1}),
	).All(&result)
	return result, err
}

// GetPublishedBanners Insert the user on a database
func (c *BannerRepo) GetPublishedBanners() ([]types.Banner, error) {
	result := []types.Banner{}
	err := c.Collection.Pipe(
		c.makeBannerComplete(
			bson.M{"status": "PUBLISHED"},
			bson.M{"creationDate": 1}),
	).All(&result)
	return result, err
}

func (c *BannerRepo) makeBannerComplete(match bson.M, sort bson.M) []bson.M {
	var response []bson.M
	if len(match) != 0 {
		response = []bson.M{bson.M{"$match": match}}
	}
	response = append(response, []bson.M{
		bson.M{
			"$lookup": bson.M{
				"from":         "club",
				"localField":   "action",
				"foreignField": "_id",
				"as":           "club"}},
		bson.M{"$unwind": bson.M{
			"path": "$club",
			"preserveNullAndEmptyArrays": true,
		}},
		bson.M{
			"$lookup": bson.M{
				"from":         "party",
				"localField":   "action",
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
