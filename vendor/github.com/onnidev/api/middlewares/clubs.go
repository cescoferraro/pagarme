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

// CollectionKey is a context key
var CollectionKey key = "clubs-repo"

// ByIDKey is important
var ByIDKey key = "club"

// AttachClubsCollection mongo wise
func AttachClubsCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		collection, err := interfaces.NewClubsCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}

		ctx := context.WithValue(r.Context(), CollectionKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetClubFromPartyMiddleware get the mogo collection
func GetClubFromPartyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clubsCollection := r.Context().Value(CollectionKey).(interfaces.ClubsRepo)
		party := r.Context().Value(PartyIDKey).(types.Party)
		club, err := clubsCollection.GetByID(party.ClubID.Hex())
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		ctx := context.WithValue(r.Context(), ByIDKey, club)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ClubPatchKey sdkfjn
var ClubPatchKey key = "club-patch"

// ReadClubPatchFromBody skjdfn
func ReadClubPatchFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.ClubPatch
		err := decoder.Decode(&query)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(query, "", "    ")
			log.Println("Club Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ClubPatchKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Stage1Key sdkfjn
var Stage1Key key = "stage1-key"

// ReadClubStage1FromBody skjdfn
func ReadClubStage1FromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.ClubStage1
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
		ctx := context.WithValue(r.Context(), Stage1Key, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ClubPostRequestKey sdkfjn
var ClubPostRequestKey key = "clulb-post-request-key"

// ReadClubPostRequestFromBody skjdfn
func ReadClubPostRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.ClubPostRequest
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
		ctx := context.WithValue(r.Context(), ClubPostRequestKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
