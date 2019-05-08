package interfaces

import (
	"fmt"
	"strings"

	"github.com/fatih/structs"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CustomersRepo is a struc that hold a mongo collection
type CustomersRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewCustomerCollection creates a new UserDAO
func NewCustomerCollection(store *infra.MongoStore) (CustomersRepo, error) {
	repo := CustomersRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("customer"),
	}
	return repo, nil
}

// GetByEmail Insert the user on a database
func (c *CustomersRepo) GetByEmail(email string) (types.Customer, error) {
	result := types.Customer{}
	err := c.Collection.Find(bson.M{"mail": email}).One(&result)
	if err != nil {
		return types.Customer{}, err
	}
	return result, nil
}

// FacebookIDExists Insert the user on a database
func (c *CustomersRepo) FacebookIDExists(id string) (int, error) {
	return c.Collection.Find(bson.M{"facebookId": id}).Count()
}

// GetByFacebookID Insert the user on a database
func (c *CustomersRepo) GetByFacebookID(id string) (types.Customer, error) {
	result := types.Customer{}
	err := c.Collection.Find(bson.M{"facebookId": id}).One(&result)
	if err != nil {
		return types.Customer{}, err
	}
	return result, nil
}

// ExistsByKey TODO: NEEDS COMMENT INFO
func (c *CustomersRepo) ExistsByKey(key string, value interface{}) (bool, error) {
	count, err := c.Collection.Find(bson.M{key: value}).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// GetByID Insert the user on a database
func (c *CustomersRepo) GetByID(id string) (types.Customer, error) {
	result := types.Customer{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.Customer{}, err
		}
		return result, nil
	}
	return types.Customer{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *CustomersRepo) List() ([]types.Customer, error) {
	result := []types.Customer{}
	err := c.Collection.Find(bson.M{}).Limit(100).All(&result)
	return result, err
}

// Query Insert the user on a database
func (c *CustomersRepo) Query(query types.CustomerQuery) ([]types.Customer, error) {
	result := []types.Customer{}
	err := c.Collection.Find(
		bson.M{
			"mail": bson.RegEx{Pattern: query.Query, Options: ""},
		}).All(&result)
	if err != nil {
		return []types.Customer{}, err
	}
	return result, nil
}

// Login Insert the user on a database
func (c *CustomersRepo) Login(login types.LoginRequest) (types.Customer, error) {
	var result types.Customer
	err := c.Collection.
		Find(bson.M{
			"password": bson.M{
				"$eq": encryptPassword2(login.Password),
			},
			"$or": []bson.M{
				bson.M{"mail": bson.M{"$eq": shared.NormalizeEmail(login.Email)}},
				bson.M{"username": bson.M{"$eq": login.UserName}},
			},
		}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Patch whatever
func (c *CustomersRepo) Patch(id bson.ObjectId, customerPatch types.CustomerPatch) error {
	coolmap := structs.Map(&customerPatch)
	setBSON := bson.M{}
	forbideen := []string{
		"id",
		"creationdate",
		"updatedate",
	}
	for key, value := range coolmap {
		if !contains(forbideen, strings.ToLower(key)) {
			setBSON[strings.ToLower(key)] = value
		}
	}
	changes := bson.M{"$set": setBSON}
	err := c.Collection.UpdateId(id.Hex(), changes)
	return err
}
