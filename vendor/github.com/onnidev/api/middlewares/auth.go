package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// PagarMeReAuthRequestKey is the shit
var PagarMeReAuthRequestKey key = "pg-jwt-refresh"

// ReadCreatePagarMeReAuthRequestFromBody skjdfn
func ReadCreatePagarMeReAuthRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var tokenReq types.PagarMeReAuth
		err := decoder.Decode(&tokenReq)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(tokenReq, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), PagarMeReAuthRequestKey, tokenReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// JWTRefreshRequestKey is the shit
var JWTRefreshRequestKey key = "jwt-refresh"

// ReadCreateJWTRefreshRequestFromBody skjdfn
func ReadCreateJWTRefreshRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var tokenReq types.JWTRefresh
		err := decoder.Decode(&tokenReq)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(tokenReq, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), JWTRefreshRequestKey, tokenReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
