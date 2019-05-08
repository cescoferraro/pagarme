package interfaces

import (
	"fmt"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PartiesRepo is a struc that hold a mongo collection
type PartiesRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewPartiesCollection creates a new UserDAO
func NewPartiesCollection(store *infra.MongoStore) (PartiesRepo, error) {
	repo := PartiesRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("party"),
	}
	return repo, nil
}

// GetByAppClubID Insert the user on a database
func (c *PartiesRepo) GetByAppClubID(id string, app bool) ([]types.AppParty, error) {
	result := []types.AppParty{}
	match := bson.M{"clubId": bson.ObjectIdHex(id)}
	if app {
		match = bson.M{
			"clubId": bson.ObjectIdHex(id),
			"status": "ACTIVE",
		}
	}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.
			Pipe(includeClubField(
				match,
				bson.M{"startDate": -1},
			)).
			All(&result)
		if err != nil {
			return []types.AppParty{}, err
		}
		return result, nil
	}
	return []types.AppParty{}, fmt.Errorf("not a valid object id")
}

// GetByClubID Insert the user on a database
func (c *PartiesRepo) GetByClubID(id string, app bool) ([]types.Party, error) {
	result := []types.Party{}
	match := bson.M{"clubId": bson.ObjectIdHex(id)}
	if app {
		match = bson.M{
			"clubId": bson.ObjectIdHex(id),
			"status": "ACTIVE",
		}
	}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.
			Pipe(includeClubField(
				match,
				bson.M{"startDate": -1},
			)).
			All(&result)
		if err != nil {
			return []types.Party{}, err
		}
		return result, nil
	}
	return []types.Party{}, fmt.Errorf("not a valid object id")
}

// GetByClubIDSite Insert the user on a database
func (c *PartiesRepo) GetByClubIDSite(id string) ([]types.Party, error) {
	result := []types.Party{}
	match := bson.M{
		"clubId":  bson.ObjectIdHex(id),
		"endDate": bson.M{"$gt": time.Now()},
		"status":  "ACTIVE",
	}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.
			Pipe(includeClubField(
				match,
				bson.M{"startDate": -1},
			)).
			All(&result)
		if err != nil {
			return []types.Party{}, err
		}
		return result, nil
	}
	return []types.Party{}, fmt.Errorf("not a valid object id")
}

// GetByClubs TODO: NEEDS COMMENT INFO
func (c *PartiesRepo) GetByClubs(clubs []types.Club) ([]types.Party, error) {
	var clubIDS []bson.ObjectId
	for _, club := range clubs {
		clubIDS = append(clubIDS, club.ID)
	}
	return c.GetByClubIDS(clubIDS)
}

// GetByClubIDS Insert the user on a database
func (c *PartiesRepo) GetByClubIDS(ids []bson.ObjectId) ([]types.Party, error) {
	result := []types.Party{}
	match := bson.M{
		"clubId":  bson.M{"$in": ids},
		"endDate": bson.M{"$gt": time.Now()},
		"status":  "ACTIVE",
	}
	err := c.Collection.
		Pipe(includeClubField(
			match,
			bson.M{"startDate": -1},
		)).
		All(&result)
	if err != nil {
		return []types.Party{}, err
	}
	return result, nil
}

// GetByAppID Insert the user on a database
func (c *PartiesRepo) GetByAppID(id string) (types.AppParty, error) {
	result := types.AppParty{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.
			Pipe(includeClubField(
				bson.M{"_id": bson.ObjectIdHex(id)},
				bson.M{"creationDate": 1},
			)).
			One(&result)
		if err != nil {
			return types.AppParty{}, err
		}
		return result, nil
	}
	return types.AppParty{}, fmt.Errorf("not a valid object id")
}

// GetByID Insert the user on a database
func (c *PartiesRepo) GetByID(id string) (types.Party, error) {
	result := types.Party{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.
			Pipe(includeClubField(
				bson.M{"_id": bson.ObjectIdHex(id)},
				bson.M{"creationDate": 1},
			)).
			One(&result)
		if err != nil {
			return types.Party{}, err
		}
		return result, nil
	}
	return types.Party{}, fmt.Errorf("not a valid object id")
}

// Latest Insert the user on a database
func (c *PartiesRepo) Latest() ([]types.Party, error) {
	result := []types.Party{}
	// now := time.Now()
	fromDate := time.Now()
	err := c.Collection.
		Pipe(includeClubField(
			bson.M{"$gt": fromDate},
			bson.M{},
		)).
		All(&result)
	return result, err
}

// List Insert the user on a database
func (c *PartiesRepo) List() ([]types.Party, error) {
	result := []types.Party{}
	err := c.Collection.
		Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"status": bson.M{"$eq": "ACTIVE"},
			}},
			bson.M{"$lookup": bson.M{
				"from":         "club",
				"localField":   "clubId",
				"foreignField": "_id",
				"as":           "club"}},
			bson.M{"$unwind": bson.M{
				"path": "$club",
				"preserveNullAndEmptyArrays": true}},
		}).
		All(&result)
	return result, err
}

// AppFilteredList Insert the user on a database
func (c *PartiesRepo) AppFilteredList(filter types.PartyFilter) ([]types.AppParty, error) {
	result := []types.AppParty{}
	err := c.Collection.
		Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"endDate": bson.M{"$gt": time.Now()},
				"status":  bson.M{"$eq": "ACTIVE"}}},
			bson.M{"$lookup": bson.M{
				"from":         "club",
				"localField":   "clubId",
				"foreignField": "_id",
				"as":           "club"}},
			bson.M{"$unwind": bson.M{
				"path": "$club",
				"preserveNullAndEmptyArrays": true}},
		}).All(&result)
	return result, err
}

// FilteredList Insert the user on a database
func (c *PartiesRepo) FilteredList(filter types.PartyFilter) ([]types.Party, error) {
	result := []types.Party{}

	err := c.Collection.
		Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"endDate": bson.M{"$gt": time.Now()},
				"status":  bson.M{"$eq": "ACTIVE"}}},
			bson.M{"$lookup": bson.M{
				"from":         "club",
				"localField":   "clubId",
				"foreignField": "_id",
				"as":           "club"}},
			bson.M{"$unwind": bson.M{
				"path": "$club",
				"preserveNullAndEmptyArrays": true}},
		}).All(&result)
	return result, err
}

// CloseParty Insert the user on a database
func (c *PartiesRepo) CloseParty(id string) (types.Party, error) {
	var result types.Party
	if bson.IsObjectIdHex(id) {
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{"status": "CLOSED"}},
			ReturnNew: true,
		}
		_, err := c.Collection.
			Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &result)
		if err != nil {
			return types.Party{}, err
		}
		return result, nil
	}
	return types.Party{}, fmt.Errorf("not a valid object id")
}

func includeClubField(match, sort bson.M) []bson.M {
	return []bson.M{
		bson.M{"$match": match},
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
