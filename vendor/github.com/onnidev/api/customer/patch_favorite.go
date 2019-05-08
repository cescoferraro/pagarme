package customer

import (
	"net/http"
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
func PatchFavorite(w http.ResponseWriter, r *http.Request) {
	customerRepo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	customerPatch := r.Context().Value(middlewares.CustomerPatchKey).(types.CustomerPatch)
	customer := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	a := customer.FavoriteClubs
	for i, favoriteClub := range a {
		for _, remove := range customerPatch.RemoveFavoriteClubs {
			if favoriteClub.Hex() == remove {
				a = append(a[:i], a[i+1:]...)
			}
		}
	}
	clubsRepo := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	for _, add := range customerPatch.AddFavoriteClubs {
		_, err := clubsRepo.GetByID(add)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		a = append(a, bson.ObjectIdHex(add))
	}
	a = removeDuplicatesUnordered(a)
	updateDate := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate":    &updateDate,
			"favoriteClubs": a,
		}},
		ReturnNew: true,
	}
	var patchedCustomer types.Customer
	_, err := customerRepo.Collection.Find(bson.M{"_id": customer.ID}).Apply(change, &patchedCustomer)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, onni.RemoveInactiveClubs(r.Context(), patchedCustomer))
}
