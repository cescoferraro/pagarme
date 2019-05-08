package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// PromotionPatchRequestKey TODO: NEEDS COMMENT INFO
var PromotionPatchRequestKey key = "promotion-patch-party-key"

// ReadPromotionPatchRequestRequestFromBody skjdfn
func ReadPromotionPatchRequestRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var filters types.PromotionPatchRequest
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
		ctx := context.WithValue(r.Context(), PromotionPatchRequestKey, filters)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// PromotionPostRequestKey TODO: NEEDS COMMENT INFO
var PromotionPostRequestKey key = "promotion-filter-party-key"

// ReadPromotionPostRequestRequestFromBody skjdfn
func ReadPromotionPostRequestRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var filters types.PromotionPostRequest
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
		ctx := context.WithValue(r.Context(), PromotionPostRequestKey, filters)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
