package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// CustomersRepoKey sdkfjn
var CustomersRepoKey key = "customer-repo"

// AttachCustomerCollection get the mogo collection
func AttachCustomerCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		collection, err := interfaces.NewCustomerCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}

		ctx := context.WithValue(r.Context(), CustomersRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TranferableCustomerKey is the shit
var TranferableCustomerKey key = "tranferable-cards-request-repo"

// ReadCustomerTranferableEmailFromBody skjdfn
func ReadCustomerTranferableEmailFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body types.TranferBody
		err := decoder.Decode(&body)
		if err != nil {
			log.Println("sdfjdsf")
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(body, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}

		customerRepo := r.Context().Value(CustomersRepoKey).(interfaces.CustomersRepo)
		user, err := customerRepo.GetByEmail(body.Email)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusPreconditionFailed), http.StatusPreconditionFailed)
			return
		}
		ctx := context.WithValue(r.Context(), TranferableCustomerKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CustomersQueryKey sdkfjn
var CustomersQueryKey key = "customer-query"

// ReadCustomerQueryFromBody skjdfn
func ReadCustomerQueryFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.CustomerQuery
		err := decoder.Decode(&query)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(query, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), CustomersQueryKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CustomersKey sdkfjn
var CustomersKey key = "customer"

// GetCustomerFromToken get the mogo collection
func GetCustomerFromToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var hey types.CustomerClaims
		_, err := jwt.ParseWithClaims(
			r.Header.Get("JWT_TOKEN"),
			&hey,
			shared.JWTAuth.Options.ValidationKeyGetter)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		customerCollection := r.Context().Value(CustomersRepoKey).(interfaces.CustomersRepo)
		customer, err := customerCollection.GetByID(hey.CustomerID)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		ctx := context.WithValue(r.Context(), CustomersKey, customer)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CustomerPatchKey sdkfjn
var CustomerPatchKey key = "customer-patch"

// ReadCustomerPatchFromBody skjdfn
func ReadCustomerPatchFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.CustomerPatch
		err := decoder.Decode(&query)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(query, "", "    ")
			log.Println("Customer Request received from body")
			log.Println(string(j))
		}
		// Check wheever que request is trying to add & remove
		// the same favorite club from an customer
		for _, add := range query.AddFavoriteClubs {
			for _, remove := range query.RemoveFavoriteClubs {
				if add == remove {
					http.Error(w, "bad.request", http.StatusBadRequest)
					return
				}
			}
		}
		ctx := context.WithValue(r.Context(), CustomerPatchKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CustomerCheckKey sdkfjn
var CustomerCheckKey key = "customer-check"

// ReadCustomerPatchFromBody skjdfn
func ReadCustomerCheckFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.CustomerCheck
		err := decoder.Decode(&query)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(query, "", "    ")
			log.Println("Customer Check Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), CustomerCheckKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ResetKey sdkfjn
var CustomerRequestKey key = "post-customer-reset"

// ReadCustomerPatchFromBody skjdfn
func ReadCustomerPostFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.CustomerPostRequest
		err := decoder.Decode(&query)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		j, _ := json.MarshalIndent(query, "", "    ")
		log.Println("Customer Reset Request received from body")
		log.Println(string(j))
		ctx := context.WithValue(r.Context(), CustomerRequestKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ResetKey sdkfjn
var ResetKey key = "customer-reset"

// ReadCustomerPatchFromBody skjdfn
func ReadResetFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.Reset
		err := decoder.Decode(&query)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		j, _ := json.MarshalIndent(query, "", "    ")
		log.Println("Customer Reset Request received from body")
		log.Println(string(j))
		ctx := context.WithValue(r.Context(), ResetKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
