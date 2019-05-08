package interfaces

import (
	"fmt"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TokensRepo is a struc that hold a mongo collection
type TokensRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewTokenCollection creates a new TokenDAO
func NewTokenCollection(store *infra.MongoStore) (TokensRepo, error) {
	repo := TokensRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("securityToken"),
	}
	return repo, nil
}

// Create Insert the user on a database
func (c *TokensRepo) Create(id bson.ObjectId) (types.Token, error) {
	horario := types.Timestamp(time.Now())
	token := types.Token{
		ID:           bson.NewObjectId(),
		CreationDate: &horario,
		UserID:       id,
		Token:        uuid.NewV4().String(),
	}
	err := c.Collection.Insert(token)
	if err != nil {
		return types.Token{}, err
	}
	return token, nil
}

// GetByToken Insert the user on a database
func (c *TokensRepo) GetByToken(token string) (types.Token, error) {
	result := types.Token{}
	err := c.Collection.Find(bson.M{"token": token}).One(&result)
	if err != nil {
		return types.Token{}, err
	}
	return result, nil
}

// GetByID Insert the user on a database
func (c *TokensRepo) GetByID(id string) (types.Token, error) {
	result := types.Token{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.Token{}, err
		}
		return result, nil
	}
	return types.Token{}, fmt.Errorf("not a valid object id")
}

// List lists a token collection
func (c *TokensRepo) List() ([]types.Token, error) {
	result := []types.Token{}
	err := c.Collection.Find(bson.M{}).Limit(10).All(&result)
	return result, err
}
