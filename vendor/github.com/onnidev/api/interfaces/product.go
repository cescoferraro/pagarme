package interfaces

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/structs"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ProductsRepo is a struc that hold a mongo collection
type ProductsRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewProductsCollection creates a new CardDAO
func NewProductsCollection(store *infra.MongoStore) (ProductsRepo, error) {
	repo := ProductsRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("product"),
	}
	return repo, nil
}

// GetByID Insert the user on a database
func (c *ProductsRepo) GetByID(id string) (types.Product, error) {
	result := types.Product{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.Product{}, err
		}
		return result, nil
	}
	return types.Product{}, fmt.Errorf("not a valid object id")
}

// List Insert the user on a database
func (c *ProductsRepo) List() ([]types.Product, error) {
	result := []types.Product{}
	err := c.Collection.Find(bson.M{"deprecated": false}).Sort("-creationDate").All(&result)
	return result, err
}

// DeleteByID is commented
func (c *ProductsRepo) DeleteByID(id string) error {
	err := c.Collection.RemoveId(bson.ObjectIdHex(id))
	return err
}

// Patch is commented
func (c *ProductsRepo) PatchWithImage(id string, file *mgo.GridFile, req types.ProductPatchRequest) error {
	setBSON := bson.M{}
	coolmap := structs.Map(&req)
	forbideen := []string{}
	horario := types.Timestamp(file.UploadDate())
	for key, value := range coolmap {
		if !contains(forbideen, key) && value != "" {
			if !(key == "Image") {
				setBSON[strings.ToLower(key)] = value
			} else {
				setBSON[strings.ToLower(key)] = types.Image{
					FileID:       bson.ObjectIdHex(value.(string)),
					MimeType:     "IMAGE_PNG",
					CreationDate: &horario,
				}
			}
		}
	}
	if req.Type == "TICKET" {
		setBSON["category"] = ""
	}
	changes := bson.M{"$set": setBSON}
	log.Println(changes)
	err := c.Collection.UpdateId(bson.ObjectIdHex(id), changes)
	return err
}

// Patch is commented
func (c *ProductsRepo) Patch(id string, req types.ProductPatchRequest) error {
	setBSON := bson.M{}
	coolmap := structs.Map(&req)
	forbideen := []string{"Image"}
	for key, value := range coolmap {
		if !contains(forbideen, key) && value != "" {
			log.Println(key, value)
			setBSON[strings.ToLower(key)] = value
		}
	}
	if req.Type == "TICKET" {
		setBSON["category"] = ""
	}
	changes := bson.M{"$set": setBSON}
	log.Println(changes)
	err := c.Collection.UpdateId(bson.ObjectIdHex(id), changes)
	return err
}
