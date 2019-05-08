package interfaces

import (
	"fmt"
	"log"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// InvitedCustomerRepo is a struc that hold a mongo collection
type InvitedCustomerRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewInvitedCustomerCollection creates a new CardDAO
func NewInvitedCustomerCollection(store *infra.MongoStore) (InvitedCustomerRepo, error) {
	repo := InvitedCustomerRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("invitedCustomer"),
	}
	for _, key := range []string{"mail"} {
		index := mgo.Index{
			Key:    []string{key},
			Unique: true,
		}
		if err := repo.Collection.EnsureIndex(index); err != nil {
			return repo, err
		}
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *InvitedCustomerRepo) GetByID(id string) (types.InvitedCustomer, error) {
	result := types.InvitedCustomer{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe(
			c.makeInvitedCustomerComplete(
				bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"creationDate": 1})).One(&result)
		if err != nil {
			return types.InvitedCustomer{}, err
		}
		return result, nil
	}
	return types.InvitedCustomer{}, fmt.Errorf("not a valid object id")
}

// GetByMail Insert the user on a database
func (c *InvitedCustomerRepo) GetByMail(mail string) (types.InvitedCustomer, error) {
	result := types.InvitedCustomer{}
	err := c.Collection.Pipe(
		c.makeInvitedCustomerComplete(
			bson.M{"mail": mail}, bson.M{"creationDate": 1})).One(&result)
	if err != nil {
		return types.InvitedCustomer{}, err
	}
	return result, nil
}

// List Insert the user on a database
func (c *InvitedCustomerRepo) List() ([]types.InvitedCustomer, error) {
	result := []types.InvitedCustomer{}
	fromDate := time.Now().Add(-time.Duration(100000) * time.Hour)
	err := c.Collection.Pipe(
		c.makeInvitedCustomerComplete(
			bson.M{"creationDate": bson.M{"$gt": fromDate}},
			bson.M{"creationDate": -1}),
	).All(&result)
	return result, err
}

func (c *InvitedCustomerRepo) makeInvitedCustomerComplete(match bson.M, sort bson.M) []bson.M {
	var response []bson.M
	if len(match) != 0 {
		response = []bson.M{bson.M{"$match": match}}
	}
	response = append(response, []bson.M{
		bson.M{"$sort": sort},
	}...)
	return response
}

// AddFacebookID Insert the user on a database
func (c *InvitedCustomerRepo) AddFacebookID(id string, fbid types.FacebookValidation, profile types.FacebookShit) (types.InvitedCustomer, error) {
	result := types.InvitedCustomer{}
	if bson.IsObjectIdHex(id) {
		updateDate := types.Timestamp(time.Now())
		log.Println("using")
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"updateDate":   &updateDate,
					"fbid":         &fbid.Data.AppID,
					"assignedMail": &profile.Email,
				}},
			ReturnNew: true,
		}
		_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &result)
		if err != nil {
			return types.InvitedCustomer{}, err
		}
		return result, nil
	}
	return types.InvitedCustomer{}, fmt.Errorf("not a valid object id")
}

// ChangeEmail Insert the user on a database
func (c *InvitedCustomerRepo) ChangeEmail(id, mail string) (types.InvitedCustomer, error) {
	result := types.InvitedCustomer{}
	if bson.IsObjectIdHex(id) {
		updateDate := types.Timestamp(time.Now())
		log.Println("using")
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"updateDate":   &updateDate,
					"assignedMail": &mail,
				}},
			ReturnNew: true,
		}
		_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &result)
		if err != nil {
			return types.InvitedCustomer{}, err
		}
		return result, nil
	}
	return types.InvitedCustomer{}, fmt.Errorf("not a valid object id")
}

// ChangeEmail Insert the user on a database
func (c *InvitedCustomerRepo) Done(id string) (types.InvitedCustomer, error) {
	result := types.InvitedCustomer{}
	if bson.IsObjectIdHex(id) {
		updateDate := types.Timestamp(time.Now())
		log.Println("using")
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"updateDate": &updateDate,
					"done":       true,
				}},
			ReturnNew: true,
		}
		_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &result)
		if err != nil {
			return types.InvitedCustomer{}, err
		}
		return result, nil
	}
	return types.InvitedCustomer{}, fmt.Errorf("not a valid object id")
}
