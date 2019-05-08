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

// ClubLeadKey is a context key
var ClubLeadKey key = "lead-clubs-repo"

// AttachClubLeadCollection mongo wise
func AttachClubLeadCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		collection, err := interfaces.NewClubLeadCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		ctx := context.WithValue(r.Context(), ClubLeadKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ClubLeadRequestKey is the shit
var ClubLeadRequestKey key = "clublead-request"

// ReadCreateClubLeadRequestFromBody skjdfn
func ReadCreateClubLeadRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.ClubLeadPostRequest
		err := decoder.Decode(&card)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(card, "", "    ")
			log.Println("[BANNER] Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ClubLeadRequestKey, card)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ClubLeadPatchRequestKey is the shit
var ClubLeadPatchRequestKey key = "club-lead-patch-request"

// ReadPatchClubLeadRequestFromBody skjdfn
func ReadPatchClubLeadRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var banner types.ClubLeadPatch
		err := decoder.Decode(&banner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(banner, "", "    ")
			log.Println("[BANNER] Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ClubLeadPatchRequestKey, banner)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
