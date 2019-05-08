package interfaces

import (
	"fmt"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PushRegistryRepo is a struc that hold a mongo collection
type PushRegistryRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewPushRegistryCollection creates a new CardDAO
func NewPushRegistryCollection(store *infra.MongoStore) (PushRegistryRepo, error) {
	repo := PushRegistryRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("pushRegistry"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *PushRegistryRepo) GetByID(id string) (types.PushRegistry, error) {
	result := types.PushRegistry{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.PushRegistry{}, err
		}
		return result, nil
	}
	return types.PushRegistry{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *PushRegistryRepo) List() ([]types.PushRegistry, error) {
	result := []types.PushRegistry{}
	err := c.Collection.Find(bson.M{}).All(&result)
	return result, err
}

// Create Insert the card on a database
func (c *PushRegistryRepo) Create(ban types.PushRegistry) error {
	err := c.Collection.Insert(ban)
	if err != nil {
		return err
	}
	return nil
}

// GetCustomersTokens Insert the user on a database
func (c *PushRegistryRepo) GetCustomersUUID(id string) ([]string, error) {
	result := []types.PushRegistry{}
	err := c.Collection.
		Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"customerId": bson.ObjectIdHex(id),
			}},
			bson.M{"$sort": bson.M{"startDate": -1}},
		}).
		All(&result)
	if err != nil {
		return []string{}, err
	}
	var aIDS []string
	for _, device := range result {
		if !shared.Contains(aIDS, device.DeviceUUID) {
			aIDS = append(aIDS, device.DeviceUUID)
		}
	}
	return aIDS, nil
}

// GetMobileTokenFromCustomers Insert the user on a database
func (c *PushRegistryRepo) GetMobileTokenFromCustomers(ids []bson.ObjectId, platform string) ([]string, error) {
	result := []types.PushRegistry{}
	err := c.Collection.
		Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"customerId": bson.M{"$in": ids},
				"platform":   platform,
			}},
			bson.M{"$sort": bson.M{"startDate": -1}},
		}).
		All(&result)
	if err != nil {
		return []string{}, err
	}
	var aIDS []string
	for _, device := range result {
		if !shared.Contains(aIDS, device.DeviceToken) {
			aIDS = append(aIDS, device.DeviceToken)
		}
	}
	return aIDS, nil
}
