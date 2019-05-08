package interfaces

import (
	"errors"
	"fmt"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PartyProductsRepo is a struc that hold a mongo collection
type PartyProductsRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewPartyProductsCollection creates a new CardDAO
func NewPartyProductsCollection(store *infra.MongoStore) (PartyProductsRepo, error) {
	repo := PartyProductsRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("partyProduct"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *PartyProductsRepo) GetByID(id string) (types.PartyProduct, error) {
	result := types.PartyProduct{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.PartyProduct{}, err
		}
		return result, nil
	}
	return types.PartyProduct{}, fmt.Errorf("not a valid object id")
}

// GetPromotion Insert the user on a database
func (c *PartyProductsRepo) GetPromotion(id string) (types.PartyProduct, types.Promotion, error) {
	endresult := types.Promotion{}
	products := []types.PartyProduct{}
	product := types.PartyProduct{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe([]bson.M{}).All(&products)
		if err != nil {
			return product, endresult, err
		}
		for _, product := range products {
			if product.PromotionalPrices != nil {
				for _, promotion := range *product.PromotionalPrices {
					if promotion.ID.Hex() == id {
						return product, promotion, nil
					}
				}
			}
		}
		return product, endresult, errors.New("not found")
	}
	return product, endresult, errors.New("not a valid object id")
}

// GetTicketsByPartyID Insert the user on a database
func (c *PartyProductsRepo) GetTicketsByPartyID(id string) ([]types.PartyProduct, error) {
	result := []types.PartyProduct{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"partyId": bson.ObjectIdHex(id),
				"type":    "TICKET",
			}},
		}).All(&result)
		if err != nil {
			return []types.PartyProduct{}, err
		}
		return result, nil
	}
	return []types.PartyProduct{}, errors.New("not a valid object id")
}

// GetDrinksByPartyID Insert the user on a database
func (c *PartyProductsRepo) GetDrinksByPartyID(id string) ([]types.PartyProduct, error) {
	result := []types.PartyProduct{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"partyId": bson.ObjectIdHex(id),
				"type":    "DRINK",
			}},
		}).All(&result)
		if err != nil {
			return []types.PartyProduct{}, err
		}
		return result, nil
	}
	return []types.PartyProduct{}, errors.New("not a valid object id")
}

// GetByPartyID Insert the user on a database
func (c *PartyProductsRepo) GetByPartyID(id string) ([]types.PartyProduct, error) {
	result := []types.PartyProduct{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"partyId": bson.ObjectIdHex(id),
			}},
		}).All(&result)
		if err != nil {
			return []types.PartyProduct{}, err
		}
		return result, nil
	}
	return []types.PartyProduct{}, errors.New("not a valid object id")
}

// GetByPartyID Insert the user on a database
func (c *PartyProductsRepo) GetByPartyIDAndType(id, kind string) ([]types.PartyProduct, error) {
	result := []types.PartyProduct{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"partyId": bson.ObjectIdHex(id),
				"type":    kind,
			}},
		}).All(&result)
		if err != nil {
			return []types.PartyProduct{}, err
		}
		return result, nil
	}
	return []types.PartyProduct{}, errors.New("not a valid object id")
}

// CountByPartyIDandType Insert the user on a database
func (c *PartyProductsRepo) CountByPartyIDandType(id, typs string) (int, error) {
	result := []types.PartyProduct{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"partyId": bson.ObjectIdHex(id),
				"type":    typs,
			}},
		}).All(&result)
		if err != nil {
			return 0, err
		}
		return len(result), nil
	}
	return 0, errors.New("not a valid object id")
}
