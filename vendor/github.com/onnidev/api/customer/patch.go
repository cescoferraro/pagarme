package customer

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Patch skdjfn
func Patch(w http.ResponseWriter, r *http.Request) {
	customer, err := PatchCustomer(r.Context())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.JSON(w, r, customer)
}

// FixUpPatch skdjfn
func FixUpPatch(w http.ResponseWriter, r *http.Request) {
	customer, err := PatchCustomer(r.Context())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	token, err := customer.GenerateToken()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.JSON(w, r, token)
}

// PatchCustomer TODO: NEEDS COMMENT INFO
func PatchCustomer(ctx context.Context) (types.Customer, error) {
	log.Println("########## endpoint thing")
	customerRepo := ctx.Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	customerPatch := ctx.Value(middlewares.CustomerPatchKey).(types.CustomerPatch)
	customer := ctx.Value(middlewares.CustomersKey).(types.Customer)
	// var a []bson.ObjectId
	// for _, favoriteClub := range customer.FavoriteClubs {
	// 	for _, remove := range customerPatch.RemoveFavoriteClubs {
	// 		if favoriteClub.Hex() != remove {
	// 			if !shared.ContainsObjectID(a, favoriteClub) {
	// 				a = append(a, favoriteClub)
	// 			}
	// 		}
	// 	}
	// }
	a := customer.FavoriteClubs
	for i, favoriteClub := range a {
		for _, remove := range customerPatch.RemoveFavoriteClubs {
			if favoriteClub.Hex() == remove {
				a = append(a[:i], a[i+1:]...)
			}
		}
	}
	clubsRepo := ctx.Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	for _, add := range customerPatch.AddFavoriteClubs {
		_, err := clubsRepo.GetByID(add)
		if err != nil {
			return types.Customer{}, err
		}
		a = append(a, bson.ObjectIdHex(add))
	}
	a = removeDuplicatesUnordered(a)
	updateDate := types.Timestamp(time.Now())
	if customerPatch.Mail != "" {
		log.Println("eustou tentando alterar username")
	}
	exists, err := customerRepo.ExistsByKey("mail", shared.NormalizeEmail(customerPatch.Mail))
	if err != nil {
		return types.Customer{}, err
	}
	if exists && customer.Mail != customerPatch.Mail {
		err := errors.New("mail already exists")
		return types.Customer{}, err
	}
	if customerPatch.UserName != "" {
		exists, err = customerRepo.ExistsByKey("username", customerPatch.UserName)
		if err != nil {
			return types.Customer{}, err
		}
	}
	if exists && customer.UserName != customerPatch.UserName {
		err := errors.New("username already exists")
		return types.Customer{}, err
	}
	var newDoc *string
	if customerPatch.DocumentNumber != "" {
		newDoc = &customerPatch.DocumentNumber

	} else {
		if customer.DocumentNumber != nil {
			newDoc = customer.DocumentNumber
		} else {
			newDoc = nil
		}
	}
	changes := bson.M{"$set": bson.M{
		"updateDate":     &updateDate,
		"firstName":      orBlank(customer.FirstName, customerPatch.FirstName),
		"lastName":       orBlank(customer.LastName, customerPatch.LastName),
		"mail":           orBlank(customer.Mail, customerPatch.Mail),
		"password":       orBlankPassword(customer.Password, shared.EncryptPassword2(customerPatch.Password)),
		"phone":          orBlank(customer.Phone, strings.Replace(customerPatch.Phone, " ", "", -1)),
		"username":       orBlank(customer.UserName, customerPatch.UserName),
		"documentNumber": newDoc,
		"birthDate":      orBlanktime(customer.BirthDate, customerPatch.BirthDate),
		"favoriteClubs":  a,
	}}
	change := mgo.Change{Update: changes, ReturnNew: true}
	var patchedCustomer types.Customer
	_, err = customerRepo.Collection.Find(bson.M{"_id": customer.ID}).Apply(change, &patchedCustomer)
	if err != nil {
		return types.Customer{}, err
	}

	return onni.RemoveInactiveClubs(ctx, patchedCustomer), nil
}
func orBlank(og, sent string) string {
	if sent == "" {
		return og
	}
	return sent
}
func orBlankPassword(og, sent string) string {
	if sent == "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=" {
		return og
	}
	return sent
}

func orBlanktime(og, sent *types.Timestamp) *types.Timestamp {
	if sent == nil {
		return og
	}
	return sent
}

func removeDuplicatesUnordered(elements []bson.ObjectId) []bson.ObjectId {
	encountered := map[bson.ObjectId]bool{}
	for v := range elements {
		encountered[elements[v]] = true
	}
	result := []bson.ObjectId{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}
