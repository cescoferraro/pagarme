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

// InvitedCustomerRepoKey sdkfjn
var InvitedCustomerRepoKey key = "invited-repo"

// AttachInvitedCustomerCollection get the mogo collection
func AttachInvitedCustomerCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		collection, err := interfaces.NewInvitedCustomerCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}

		ctx := context.WithValue(r.Context(), InvitedCustomerRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// InvitedLinkCustomerPostKey sdkfjn
var InvitedLinkCustomerPostKey key = "invite-linkpost-request-key"

// ReadInvitedLinkCustomerPostRequestFromBody skjdfn
func ReadInvitedLinkCustomerPostRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.InvitedLinkCustomerPost
		err := decoder.Decode(&query)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(query, "", "    ")
			log.Println("ClubStage1 Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), InvitedLinkCustomerPostKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
