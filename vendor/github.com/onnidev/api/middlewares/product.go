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

// ProductsRepoKey is the shit
var ProductsRepoKey key = "cards-repo"

// AttachProductsCollection skjdfn
func AttachProductsCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		collection, err := interfaces.NewProductsCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}

		ctx := context.WithValue(r.Context(), ProductsRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ReadProductPostKey sdkjfn
var ReadProductPostKey key = "read-product-post-req"

// ReadProductPostRequestFromBody skjdfn
func ReadProductPostRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var productReq types.ProductPostRequest
		err := decoder.Decode(&productReq)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(productReq, "", "    ")
			log.Println("Product Post Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ReadProductPostKey, productReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ReadProductPatchKey sdkjfn
var ReadProductPatchKey key = "read-product-patch-req"

// ReadProductPatchRequestFromBody skjdfn
func ReadProductPatchRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var productReq types.ProductPatchRequest
		err := decoder.Decode(&productReq)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(productReq, "", "    ")
			log.Println("Product Patch Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ReadProductPatchKey, productReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ReadProductSoftPatchKey sdkjfn
var ReadProductSoftPatchKey key = "sdkjnsd-0read-product-patch-req"

// ReadProductSoftPatchRequestFromBody skjdfn
func ReadProductSoftPatchRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var productReq types.ProductSoftPatchRequest
		err := decoder.Decode(&productReq)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(productReq, "", "    ")
			log.Println("Product Patch Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ReadProductSoftPatchKey, productReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
