package interfaces

import (
	"fmt"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MusicStylesRepo is a struc that hold a mongo collection
type MusicStylesRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewMusicStylesCollection creates a new CardDAO
func NewMusicStylesCollection(store *infra.MongoStore) (MusicStylesRepo, error) {
	repo := MusicStylesRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("musicStyle"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *MusicStylesRepo) GetByID(id string) (types.Style, error) {
	result := types.Style{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.Style{}, err
		}
		return result, nil
	}
	return types.Style{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *MusicStylesRepo) ByNames(names []string) ([]types.Style, error) {
	result := []types.Style{}
	err := c.Collection.
		Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"name": bson.M{"$in": names},
			}},
			bson.M{"$sort": bson.M{"creationDate": -1}},
		}).
		All(&result)
	return result, err
}

// List Insert the user on a database
func (c *MusicStylesRepo) List() ([]types.Style, error) {
	result := []types.Style{}
	err := c.Collection.Find(bson.M{}).All(&result)
	return result, err
}
