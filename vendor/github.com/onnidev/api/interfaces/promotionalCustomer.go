package interfaces

import (
	"errors"
	"fmt"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PromotionalCustomerRepo is a struc that hold a mongo collection
type PromotionalCustomerRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewPromotionalCustomerCollection creates a new CardDAO
func NewPromotionalCustomerCollection(store *infra.MongoStore) (PromotionalCustomerRepo, error) {
	repo := PromotionalCustomerRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("promotionalCustomer"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *PromotionalCustomerRepo) GetByID(id string) (types.PromotionalCustomer, error) {
	result := types.PromotionalCustomer{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe(
			c.makePromotionalCustomerComplete(
				bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"creationDate": 1})).One(&result)
		if err != nil {
			return types.PromotionalCustomer{}, err
		}
		return result, nil
	}
	return types.PromotionalCustomer{}, fmt.Errorf("not a valid object id")
}

// PromotionalCustomersFromPromotion Insert the user on a database
func (c *PromotionalCustomerRepo) PromotionalCustomersFromPromotion(promoID, customerID string) (types.PromotionalCustomer, error) {
	result := []types.PromotionalCustomer{}
	if bson.IsObjectIdHex(promoID) && bson.IsObjectIdHex(customerID) {
		err := c.Collection.Pipe(
			c.makePromotionalCustomerComplete(
				bson.M{"promotionId": bson.ObjectIdHex(promoID), "customerId": bson.ObjectIdHex(customerID)},
				bson.M{"creationDate": -1}),
		).All(&result)
		if len(result) == 0 {
			return types.PromotionalCustomer{}, errors.New("not found")
		}
		return result[0], err
	}
	return types.PromotionalCustomer{}, fmt.Errorf("not a valid object id")
}

// PromotionalCustomersAttachedToPromotion Insert the user on a database
func (c *PromotionalCustomerRepo) PromotionalCustomersAttachedToPromotion(id string) ([]types.PromotionalCustomer, error) {
	result := []types.PromotionalCustomer{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe(
			c.makePromotionalCustomerComplete(
				bson.M{"promotionId": bson.ObjectIdHex(id)},
				bson.M{"creationDate": -1}),
		).All(&result)
		return result, err
	}
	return []types.PromotionalCustomer{}, fmt.Errorf("not a valid object id")
}

// ByCustomer Insert the user on a database
func (c *PromotionalCustomerRepo) ByCustomer(id string) ([]types.PromotionalCustomer, error) {
	result := []types.PromotionalCustomer{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe(
			c.makePromotionalCustomerComplete(
				bson.M{"customerId": bson.ObjectIdHex(id)},
				bson.M{"creationDate": -1}),
		).All(&result)
		return result, err
	}
	return []types.PromotionalCustomer{}, fmt.Errorf("not a valid object id")
}

// SetInvitedNameOnPromotion TODO: NEEDS COMMENT INFO
func (c *PromotionalCustomerRepo) SetInvitedNameOnPromotion(id string, customer types.Customer) (types.PromotionalCustomer, error) {
	result := types.PromotionalCustomer{}
	if bson.IsObjectIdHex(id) {
		now := types.Timestamp(time.Now())
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"customerName": customer.Name(),
					"customerMail": customer.Mail,
					"updateDate":   &now,
				}},
			ReturnNew: true,
		}
		_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &result)
		if err != nil {
			return types.PromotionalCustomer{}, err
		}
		return result, nil
	}
	return types.PromotionalCustomer{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *PromotionalCustomerRepo) List() ([]types.PromotionalCustomer, error) {
	result := []types.PromotionalCustomer{}
	fromDate := time.Now().Add(-time.Duration(100000) * time.Hour)
	err := c.Collection.Pipe(
		c.makePromotionalCustomerComplete(
			bson.M{"creationDate": bson.M{"$gt": fromDate}},
			bson.M{"creationDate": -1}),
	).All(&result)
	return result, err
}

// GetAllFromCustomer Insert the user on a database
func (c *PromotionalCustomerRepo) GetAllFromCustomer(ids []bson.ObjectId, customerID string) ([]bson.ObjectId, error) {
	result := []types.PromotionalCustomer{}
	err := c.Collection.
		Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"promotionId": bson.M{"$in": ids},
				"customerId":  bson.ObjectIdHex(customerID),
			}},
			bson.M{"$sort": bson.M{"startDate": -1}},
		}).
		All(&result)
	if err != nil {
		return []bson.ObjectId{}, err
	}
	var aIDS []bson.ObjectId
	for _, device := range result {
		aIDS = append(aIDS, device.PromotionID)
	}
	return aIDS, nil
}

// GetFromCustomer Insert the user on a database
func (c *PromotionalCustomerRepo) MailHasThisPromotion(id bson.ObjectId, mail string) (bool, error) {
	result := []types.PromotionalCustomer{}
	err := c.Collection.
		Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"promotionId":  id,
				"customerMail": mail,
			}},
			bson.M{"$sort": bson.M{"startDate": -1}},
		}).
		All(&result)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, nil
	}
	return true, nil
}

// CustomerHasThisPromotion Insert the user on a database
func (c *PromotionalCustomerRepo) CustomerHasThisPromotion(id bson.ObjectId, customerID string) (bool, error) {
	result := []types.PromotionalCustomer{}
	err := c.Collection.
		Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"promotionId": id,
				"customerId":  bson.ObjectIdHex(customerID),
			}},
			bson.M{"$sort": bson.M{"startDate": -1}},
		}).
		All(&result)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, nil
	}
	return true, nil
}

func (c *PromotionalCustomerRepo) makePromotionalCustomerComplete(match bson.M, sort bson.M) []bson.M {
	var response []bson.M
	if len(match) != 0 {
		response = []bson.M{bson.M{"$match": match}}
	}
	response = append(response, []bson.M{
		bson.M{"$sort": sort},
	}...)
	return response
}
