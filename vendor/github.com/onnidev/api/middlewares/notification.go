package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// NotificationRepoKey is the shit
var NotificationRepoKey key = "bans-repo"

// AttachNotificationCollection skjdfn
func AttachNotificationCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		collection, err := interfaces.NewNotificationsCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}

		ctx := context.WithValue(r.Context(), NotificationRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NotificationPatchRequestKey is the shit
var NotificationPatchRequestKey key = "notification-request"

// ReadCreateNoitificationPatchRequestFromBody skjdfn
func ReadCreateNotificationPatchRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.NotificationPatchRequest
		err := decoder.Decode(&card)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(card, "", "    ")
			log.Println("[BANNER] Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), NotificationPatchRequestKey, card)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NotificationCreateRequestKey is the shit
var NotificationCreateRequestKey key = "notification-create-request"

// ReadCreateNotificationRequestFromBody skjdfn
func ReadCreateNotificationRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var notification types.NotificationPostRequest
		err := decoder.Decode(&notification)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(notification, "", "    ")
			log.Println("[BANNER] Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), NotificationCreateRequestKey, notification)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
