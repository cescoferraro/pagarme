package interfaces

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserClubRepo is a struc that hold a mongo collection
type UserClubRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewUserClubCollection creates a new UserDAO
func NewUserClubCollection(store *infra.MongoStore) (UserClubRepo, error) {
	repo := UserClubRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("userClub"),
	}
	return repo, nil
}

// ExistsByKey TODO: NEEDS COMMENT INFO
func (c *UserClubRepo) ExistsByKey(key string, value interface{}) (bool, error) {
	count, err := c.Collection.Find(bson.M{key: value}).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// Login Insert the user on a database
func (c *UserClubRepo) Login(email, password string) (types.UserClub, error) {
	var result types.UserClub
	err := c.Collection.Pipe(
		[]bson.M{
			bson.M{"$match": bson.M{
				"password": encryptPassword2(password),
				"mail":     email,
			}}}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FindClubAdmins Insert the user on a database
func (c *UserClubRepo) FindClubAdmins(clubID string) ([]types.UserClub, error) {
	var result []types.UserClub
	err := c.Collection.Pipe(
		[]bson.M{
			bson.M{"$match": bson.M{
				"profile": "ADMIN",
				"status":  "ACTIVE",
				"clubs":   bson.ObjectIdHex(clubID),
			}}}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FacebookExtraUserRegex Insert the user on a database
func (c *UserClubRepo) FacebookExtraUserRegex(email string) (types.UserClub, error) {
	result := types.UserClub{}
	half := strings.Split(email, "@")
	things := half[0] + "\\+([1-999]+)@" + half[1]
	err := c.Collection.
		Find(bson.M{
			"profile": bson.RegEx{
				Pattern: "(ATTENDANT|ADMIN)",
				Options: "ix",
			},
			"$or": []bson.M{
				bson.M{"mail": bson.M{"$eq": email}},
				bson.M{
					"mail": bson.RegEx{
						Pattern: things,
						Options: "ix",
					},
				},
			},
		}).One(&result)
	if err != nil {
		return types.UserClub{}, err
	}
	return result, nil
}

// GetByEmail Insert the user on a database
func (c *UserClubRepo) GetByEmail(email string) (types.UserClub, error) {
	result := types.UserClub{}
	err := c.Collection.Find(bson.M{"mail": email}).One(&result)
	if err != nil {
		return types.UserClub{}, err
	}
	return result, nil
}

// GetByID Insert the user on a database
func (c *UserClubRepo) GetByID(id string) (types.UserClub, error) {
	result := types.UserClub{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.UserClub{}, err
		}
		return result, nil
	}
	return types.UserClub{}, fmt.Errorf("not a valid object id")
}

// GetByIDS TODO: NEEDS COMMENT INFO
func (c *UserClubRepo) GetByIDS(ids []bson.ObjectId) ([]types.UserClub, error) {
	result := []types.UserClub{}
	err := c.Collection.
		Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"_id": bson.M{"$in": ids},
			}},
			bson.M{"$sort": bson.M{"startDate": -1}},
		}).
		All(&result)
	if err != nil {
		return []types.UserClub{}, err
	}
	return result, nil
}

// ListByClub Insert the user on a database
func (c *UserClubRepo) ListByClub(id string) ([]types.UserClub, error) {
	result := []types.UserClub{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe(
			[]bson.M{
				bson.M{"$match": bson.M{
					"clubs": bson.M{"$in": []bson.ObjectId{bson.ObjectIdHex(id)}},
				}},
				bson.M{"$sort": bson.M{"creationDate": -1}},
			}).All(&result)
		if err != nil {
			return []types.UserClub{}, err
		}
		return result, nil
	}
	return result, errors.New("dsfkj")
}

// ListPromotersByClub Insert the user on a database
func (c *UserClubRepo) ListPromotersByClub(id string) ([]types.UserClub, error) {
	result := []types.UserClub{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.Pipe(
			[]bson.M{
				bson.M{"$match": bson.M{
					"clubs":   bson.M{"$in": []bson.ObjectId{bson.ObjectIdHex(id)}},
					"profile": "PROMOTER",
				}},
				bson.M{"$sort": bson.M{"creationDate": -1}},
			}).All(&result)
		if err != nil {
			return []types.UserClub{}, err
		}
		return result, nil
	}
	return result, errors.New("dsfkj")
}

// ListOnni Insert the user on a database
func (c *UserClubRepo) ListOnni() ([]types.UserClub, error) {
	result := []types.UserClub{}
	err := c.Collection.Pipe(
		[]bson.M{
			bson.M{"$match": bson.M{
				"profile": "ONNI",
			}},
			bson.M{"$sort": bson.M{"creationDate": -1}},
		}).All(&result)
	if err != nil {
		return []types.UserClub{}, err
	}
	return result, nil
}

// List Insert the user on a database
func (c *UserClubRepo) List() ([]types.UserClub, error) {
	result := []types.UserClub{}
	err := c.Collection.Find(bson.M{}).All(&result)
	return result, err
}

func encryptPassword2(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}
