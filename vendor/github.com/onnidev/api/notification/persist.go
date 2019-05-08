package notification

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

// PersistToDB is a comented function
func PersistToDB(w http.ResponseWriter, r *http.Request) {
	notificationRequest := r.Context().Value(middlewares.NotificationCreateRequestKey).(types.NotificationPostRequest)
	partiesRepo := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	clubsRepo := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	party, err := partiesRepo.GetByID(notificationRequest.PartyID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	club, err := clubsRepo.GetByID(party.Club.ID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	now := types.Timestamp(time.Now())

	if !bson.ObjectIdHex(notificationRequest.ID).Valid() {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	notification := types.Notification{
		ID:            bson.ObjectIdHex(notificationRequest.ID),
		Type:          "TEXT",
		Status:        "DRAFT",
		Text:          notificationRequest.Text,
		Title:         notificationRequest.Title,
		CreationDate:  &now,
		PartyID:       &party.ID,
		PartyName:     &party.Name,
		ClubID:        &club.ID,
		ClubName:      &club.Name,
		CreatedBy:     userClub.ID,
		CreatedByName: userClub.Name,
	}
	notificationRepo := r.Context().Value(middlewares.NotificationRepoKey).(interfaces.NotificationRepo)
	err = notificationRepo.Create(notification)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(notification, "", "    ")
		log.Println("Notification created on MongoDB")
		log.Println(string(j))
	}
	render.JSON(w, r, notificationRequest)
}
