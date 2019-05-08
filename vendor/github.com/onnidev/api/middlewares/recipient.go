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

// RecipientCollectionKey is a context key
var RecipientCollectionKey key = "token-repo"

// AttachRecipientsCollection mongo wise
func AttachRecipientsCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		collection, err := interfaces.NewRecipientCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}

		ctx := context.WithValue(r.Context(), RecipientCollectionKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ReadAntecipationPostKey sdkjfn
var ReadAntecipationPostKey key = "read-antecipations-post-req"

// ReadAntecipationPostRequestFromBody skjdfn
func ReadAntecipationPostRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var productReq types.AntecipationPostRequest
		err := decoder.Decode(&productReq)
		if err != nil {

			shared.MakeONNiError(w, r, 400, err)

			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(productReq, "", "    ")
			log.Println("Antecipation Post Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ReadAntecipationPostKey, productReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RecipientTweaksPatchKey sdkfjn
var RecipientTweaksPatchKey key = "customer-patch"

// ReadRecipientTweaksPatchFromBody skjdfn
func ReadRecipientTweaksPatchFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.RecipientTweaksPatch
		err := decoder.Decode(&query)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(query, "", "    ")
			log.Println("Recipient Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), RecipientTweaksPatchKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RecipientPatchKey sdkfjn
var RecipientPatchKey key = "customer-patch"

// ReadRecipientPatchFromBody skjdfn
func ReadRecipientPatchFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.RecipientPatch
		err := decoder.Decode(&query)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(query, "", "    ")
			log.Println("Recipient Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), RecipientPatchKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ReadRecipientPostKey sdkjfn
var ReadRecipientPostKey key = "read-recipient-post-req"

// ReadRecipientPostRequestFromBody skjdfn
func ReadRecipientPostRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var productReq types.RecipientPost
		err := decoder.Decode(&productReq)
		if err != nil {

			shared.MakeONNiError(w, r, 400, err)

			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(productReq, "", "    ")
			log.Println("Antecipation Post Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ReadRecipientPostKey, productReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ReadRecipientWithDrawKey sdkjfn
var ReadRecipientWithDrawKey key = "read-recipient-withdraw-req"

// ReadRecipientWithDrawRequestFromBody skjdfn
func ReadRecipientWithDrawRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var productReq types.RecipientWithDraw
		err := decoder.Decode(&productReq)
		if err != nil {

			shared.MakeONNiError(w, r, 400, err)

			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(productReq, "", "    ")
			log.Println("Antecipation WithDraw Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ReadRecipientWithDrawKey, productReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
