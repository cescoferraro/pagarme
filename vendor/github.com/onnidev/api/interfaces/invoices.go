package interfaces

import (
	"fmt"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// InvoicesRepo is a struc that hold a mongo collection
type InvoicesRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewInvoicesCollection creates a new CardDAO
func NewInvoicesCollection(store *infra.MongoStore) (InvoicesRepo, error) {
	repo := InvoicesRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("invoice"),
	}
	return repo, nil
}

// GetByUserID Insert the user on a database
func (c *InvoicesRepo) GetByUserID(id string) ([]types.Invoice, error) {
	result := []types.Invoice{}
	err := c.Collection.Find(bson.M{"customerId": bson.ObjectIdHex(id)}).All(&result)
	return result, err
}

// GetUniqueDeviceIDFromCustomer Insert the user on a database
func (c *InvoicesRepo) GetUniqueDeviceIDFromCustomer(useriD string) ([]string, error) {
	result := []types.Invoice{}
	all := []string{}
	err := c.Collection.Pipe(
		[]bson.M{{"$match": bson.M{
			"customerId": bson.ObjectIdHex(useriD),
		}}}).All(&result)
	if err != nil {
		return all, err
	}
	for _, invoice := range result {
		if invoice.Log.DeviceSO != "WEB" {
			deviceID := invoice.Log.DeviceID
			if deviceID != "" {
				if !shared.Contains(all, deviceID) {
					all = append(all, deviceID)
				}
			}
		}
	}
	return all, nil
}

// ByCustomer Insert the user on a database
func (c *InvoicesRepo) ByCustomer(useriD string) ([]types.Invoice, error) {
	result := []types.Invoice{}
	err := c.Collection.Pipe(
		[]bson.M{{"$match": bson.M{
			"customerId": bson.ObjectIdHex(useriD),
		}}}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FutureByCustomer Insert the user on a database
func (c *InvoicesRepo) FutureByCustomer(useriD string) ([]types.Invoice, error) {
	result := []types.Invoice{}
	err := c.Collection.Pipe(
		[]bson.M{{"$match": bson.M{
			"customerId": bson.ObjectIdHex(useriD),
			"startDate":  bson.M{"$gt": time.Now()},
		}}}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetByID Insert the user on a database
func (c *InvoicesRepo) GetByID(id string) (types.Invoice, error) {
	result := types.Invoice{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.Invoice{}, err
		}
		return result, nil
	}
	return types.Invoice{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *InvoicesRepo) List() ([]types.Invoice, error) {
	result := []types.Invoice{}
	err := c.Collection.Find(bson.M{}).Limit(10).All(&result)
	return result, err
}

// GetByPartyID Insert the user on a database
func (c *InvoicesRepo) GetByPartyID(partyID string) ([]types.Invoice, error) {
	result := []types.Invoice{}
	if bson.IsObjectIdHex(partyID) {
		err := c.Collection.Pipe([]bson.M{{"$match": bson.M{"partyId": bson.ObjectIdHex(partyID)}}}).All(&result)
		if err != nil {
			return []types.Invoice{}, err
		}
		return result, nil
	}
	return []types.Invoice{}, fmt.Errorf("not a valid object partyID")
}
