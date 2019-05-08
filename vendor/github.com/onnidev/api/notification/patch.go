package notification

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Patch skdjfn
func Patch(w http.ResponseWriter, r *http.Request) {
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	notificationRepo := r.Context().Value(middlewares.NotificationRepoKey).(interfaces.NotificationRepo)
	notificationPatch := r.Context().Value(middlewares.NotificationPatchRequestKey).(types.NotificationPatchRequest)
	partiesRepo := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	clubsRepo := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	id := chi.URLParam(r, "notificationID")
	notification, err := notificationRepo.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if notification.Status == "DRAFT" {
		party, err := partiesRepo.GetByID(notificationPatch.PartyID)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		club, err := clubsRepo.GetByID(notificationPatch.ClubID)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		now := types.Timestamp(time.Now())
		change := mgo.Change{
			Update: bson.M{"$set": bson.M{
				"updateDate": &now,
				"updatedBy":  userClub.ID,
				"title":      orBlank(notification.Title, notificationPatch.Title),
				"text":       orBlank(notification.Text, notificationPatch.Text),

				"partyID":   &party.ID,
				"partyName": &party.Name,
				"clubID":    &club.ID,
				"clubName":  &club.Name,
			}},
			ReturnNew: true,
		}
		var patchedNotification types.Notification
		_, err = notificationRepo.Collection.Find(bson.M{"_id": notification.ID}).Apply(change, &patchedNotification)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, patchedNotification)
		return
	}
	err = errors.New("only draft notifications can be patched")
	shared.MakeONNiError(w, r, 400, err)

	return
}

func orBlank(og, sent string) string {
	if sent == "" {
		return og
	}
	return sent
}
