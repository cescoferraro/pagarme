package pushRegistry

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// Create is a comented function
func Create(w http.ResponseWriter, r *http.Request) {
	pushRegistryRequest := r.Context().
		Value(middlewares.PushRegistryRequestKey).(types.PushRegistryPostRequest)
	pushRegistryRepo := r.Context().
		Value(middlewares.PushRegistryRepoKey).(interfaces.PushRegistryRepo)
	customer := r.Context().
		Value(middlewares.CustomersKey).(types.Customer)
	now := types.Timestamp(time.Now())
	id := bson.NewObjectId()
	pushRegistry := types.PushRegistry{
		ID:           id,
		CreationDate: &now,
		Platform:     pushRegistryRequest.Platform,
		CustomerID:   customer.ID,
		DeviceUUID:   pushRegistryRequest.DeviceUUID,
		DeviceToken:  pushRegistryRequest.DeviceToken,
	}
	err := pushRegistryRepo.Create(pushRegistry)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(pushRegistry, "", "    ")
		log.Println("PushRegistry created on MongoDB")
		log.Println(string(j))
	}
	push, err := pushRegistryRepo.GetByID(id.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.JSON(w, r, push)
}
