package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// PartyRepoKey sdkfjn
var PartyRepoKey key = "party-repo"

// AttachPartiesCollection get the mogo collection
func AttachPartiesCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		collection, err := interfaces.NewPartiesCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)
			return
		}

		ctx := context.WithValue(r.Context(), PartyRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// PartyIDKey sdfkjn
var PartyIDKey key = "party-id"

// GetPartyByID get the mogo collection
func GetPartyByID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		partiesCollection := r.Context().Value(PartyRepoKey).(interfaces.PartiesRepo)
		id := chi.URLParam(r, "id")
		party, err := partiesCollection.GetByID(id)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		ctx := context.WithValue(r.Context(), PartyIDKey, party)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// PartyListFilterKey sdfkjn
var PartyListFilterKey key = "filter-party-key"

// ReadPartyListFilterRequestFromBody skjdfn
func ReadPartyListFilterRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var filters types.PartyFilter
		err := decoder.Decode(&filters)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(filters, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), PartyListFilterKey, filters)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// SoftPartyPostRequestKey sdfkjn
var SoftPartyPostRequestKey key = "create-persist-party-key"

// ReadSoftPartyPostRequestRequestFromBody skjdfn
func ReadSoftPartyPostRequestRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var filters types.SoftPartyPostRequest
		err := decoder.Decode(&filters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(filters, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), SoftPartyPostRequestKey, filters)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
