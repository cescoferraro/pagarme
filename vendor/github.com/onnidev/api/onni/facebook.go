package onni

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	fb "github.com/huandu/facebook"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// FindCustomerByFacebookID sdfjkn
func FindCustomerByFacebookID(ctx context.Context, id string) (types.Customer, error) {
	customersCollection := ctx.Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	customer, err := customersCollection.GetByFacebookID(id)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

// ExistCustomerWithFacebookID sdfjkn
func ExistCustomerWithFacebookID(ctx context.Context, id string) (bool, error) {
	customersCollection := ctx.Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	_, err := customersCollection.GetByFacebookID(id)
	if err != nil {

		if err.Error() == "not found" {
			log.Println("dentro do onni pagckage deu not found")
			return false, nil
		}
		log.Println("dentro do onni pagckage deu erro")
		return false, err
	}
	log.Println("dentro do onni pagckage found")
	return true, nil
}

// SetCustomerFacebookID sdfjkn
func SetCustomerFacebookID(ctx context.Context, id, facebookID string) (types.Customer, error) {
	customersCollection := ctx.Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	var result types.Customer
	if bson.IsObjectIdHex(id) {
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{"facebookId": facebookID}},
			ReturnNew: true,
		}
		_, err := customersCollection.Collection.
			Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &result)
		if err != nil {
			return types.Customer{}, err
		}
		return result, nil
	}
	return types.Customer{}, fmt.Errorf("not a valid object id")
}

// FindCustomerByMail sdfjkn
func FindCustomerByMail(ctx context.Context, id string) (types.Customer, error) {
	customersCollection := ctx.Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	customer, err := customersCollection.GetByEmail(id)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

// FacebookAppValidate skdjfnsd
func FacebookAppValidate(token string) (types.FacebookValidation, error) {
	return facebookValidate(token, viper.GetString("FBAPPTOKEN"))
}

// FacebookStaffAppValidate skdjfnsd
func FacebookStaffAppValidate(token string) (types.FacebookValidation, error) {
	return facebookValidate(token, viper.GetString("FBSTAFFAPPTOKEN"))
}

func facebookValidate(token, fbtoken string) (types.FacebookValidation, error) {
	var validation types.FacebookValidation
	res, err := fb.Get("/debug_token", fb.Params{
		"input_token":  token,
		"access_token": fbtoken,
	})
	if err != nil {
		return validation, err
	}
	err = res.Decode(&validation)
	if err != nil {
		return validation, err
	}
	log.Println("Token Expires at: ", time.Unix(int64(validation.Data.ExpiresAt), 0))
	if !validation.Data.IsValid {
		err := errors.New("token is not valid")
		return validation, err
	}
	return validation, nil
}

// GetFacebookProfile sjkfdnskdnfj
func GetFacebookProfile(token string) (types.FacebookShit, error) {
	var validation types.FacebookShit
	res, err := fb.Get("/me", fb.Params{
		"fields":       "id,first_name,last_name,email,birthday,photos",
		"access_token": token,
	})
	log.Println(res)
	if err != nil {
		return validation, err
	}
	err = res.Decode(&validation)
	if err != nil {
		return validation, err
	}
	return validation, nil
}
