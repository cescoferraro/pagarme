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

// PromotionalCustomerRepoKey is the shit
var PromotionalCustomerRepoKey key = "promotionalCustomer-repo"

// AttachPromotionalCustomerCollection skjdfn
func AttachPromotionalCustomerCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}

		collection, err := interfaces.NewPromotionalCustomerCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)
			return
		}

		ctx := context.WithValue(r.Context(), PromotionalCustomerRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// PromotionalCustomerPostQueryKey sdkfjn
var PromotionalCustomerPostQueryKey key = "promotinal-customer-query"

// ReadCustomerQueryFromBody skjdfn
func ReadPromotionalCustomerQueryFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var query types.PromotionalCustomerPost
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
		ctx := context.WithValue(r.Context(), PromotionalCustomerPostQueryKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
