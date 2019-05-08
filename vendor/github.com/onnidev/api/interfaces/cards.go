package interfaces

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CardsRepo is a struc that hold a mongo collection
type CardsRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewCardsCollection creates a new CardDAO
func NewCardsCollection(store *infra.MongoStore) (CardsRepo, error) {
	repo := CardsRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("customerCard"),
	}
	return repo, nil
}

// GetAllByCustomerID Insert the user on a database
func (c *CardsRepo) GetAllByCustomerID(id string) ([]types.Card, error) {
	result := []types.Card{}
	err := c.Collection.
		Find(bson.M{
			"customerId": bson.ObjectIdHex(id),
		}).
		Sort("-creationDate").All(&result)
	if err != nil {
		return []types.Card{}, err
	}
	return result, nil
}

// GetUniquiByCustomerID Insert the user on a database
func (c *CardsRepo) GetUniquiByCustomerID(id string) ([]types.Card, error) {
	result := []types.Card{}
	err := c.Collection.
		Find(bson.M{
			"customerId": bson.ObjectIdHex(id),
		}).
		Sort("-creationDate").All(&result)
	if err != nil {
		return []types.Card{}, err
	}
	end := []types.Card{}
	ends := []string{}
	for _, card := range result {
		if !shared.Contains(ends, card.Last4) {
			end = append(end, card)
			ends = append(ends, card.Last4)
		}
	}
	return end, nil
}

// GetByCustomerID Insert the user on a database
func (c *CardsRepo) GetByCustomerID(id string) ([]types.Card, error) {
	result := []types.Card{}
	err := c.Collection.
		Find(bson.M{
			"customerId": bson.ObjectIdHex(id),
		}).
		Sort("-creationDate").All(&result)
	if err != nil {
		return []types.Card{}, err
	}
	end := []types.Card{}
	for _, card := range result {
		if card.Deprecated != nil {
			dep := *card.Deprecated
			if dep == "true" {
				continue
			}
		}
		end = append(end, card)
	}
	return end, nil
}

// GetByCustomerDefaultCard Insert the user on a database
func (c *CardsRepo) GetByCustomerDefaultCard(id string) (types.Card, error) {
	result := []types.Card{}
	err := c.Collection.Find(bson.M{"customerId": bson.ObjectIdHex(id)}).
		Sort("-creationDate").All(&result)
	if err != nil {
		return types.Card{}, err
	}
	for _, card := range result {
		if card.Deprecated == nil {
			if card.Default {
				return card, nil
			}
		}
	}
	return types.Card{}, errors.New("add.customer.card.error.customer.not.found")
}

// GetByID Insert the user on a database
func (c *CardsRepo) GetByID(id string) (types.Card, error) {
	result := types.Card{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.Card{}, err
		}
		return result, nil
	}
	return types.Card{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *CardsRepo) List() ([]types.Card, error) {
	result := []types.Card{}
	err := c.Collection.Find(bson.M{}).Limit(10).All(&result)
	return result, err
}

// GetByCardToken Insert the card on a database
func (c *CardsRepo) GetByCardToken(token string) (types.Card, error) {
	var card types.Card
	err := c.Collection.Find(bson.M{"cardToken": token}).One(&card)
	if err != nil {
		return card, err
	}
	return card, nil
}

// Exists Insert the card on a database
func (c *CardsRepo) Exists(card types.Card) (bool, error) {
	count, err := c.Collection.Find(bson.M{"cardToken": card.CardToken}).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// Create Insert the card on a database
func (c *CardsRepo) Create(card types.Card) (types.Card, error) {
	err := c.Collection.Insert(card)
	if err != nil {
		return card, err
	}
	return card, err
}

// ActivateByID is commented
func (c *CardsRepo) ActivateByID(id string) error {
	horario := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate": &horario,
			"default":    true,
		}},
		ReturnNew: true,
	}
	var patchedCard types.Card
	_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &patchedCard)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByID is commented
func (c *CardsRepo) DeleteByID(id string) error {
	horario := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate": &horario,
			"deprecated": "true",
		}},
		ReturnNew: true,
	}
	var patchedCard types.Card
	_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &patchedCard)
	if err != nil {
		return err
	}
	return nil
}

// Patch is commented
func (c *CardsRepo) Patch(id string, card types.Card) error {
	coolmap := structs.Map(&card)
	setBSON := bson.M{}
	log.Println(card)
	log.Println(coolmap)
	forbideen := []string{
		"id",
		"customerid",
		"creationdate",
		"last4",
		"brand",
		"updatedate",
	}
	for key, value := range coolmap {
		log.Println(key)
		if !contains(forbideen, strings.ToLower(key)) {
			setBSON[strings.ToLower(key)] = value
		}
	}
	changes := bson.M{"$set": setBSON}
	err := c.Collection.UpdateId(bson.ObjectIdHex(id), changes)
	return err
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
