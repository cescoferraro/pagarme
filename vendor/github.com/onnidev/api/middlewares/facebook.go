package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// FacebookLoginRequestKey is the shit
var FacebookLoginRequestKey key = "facebook-login-request"

// ReadFacebookLoginRequest skjdfn
func ReadFacebookLoginRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.FacebookLoginRequest
		err := decoder.Decode(&card)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(card, "", "    ")
			log.Println("[FACEBOOK] Login Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), FacebookLoginRequestKey, card)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// FacebookSignUpRequestKey is the shit
var FacebookSignUpRequestKey key = "facebook-signup-request"

// ReadFacebookLoginRequest skjdfn
func ReadFacebookSignUpRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.FacebookSignUpRequest
		err := decoder.Decode(&card)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(card, "", "    ")
			log.Println("[FACEBOOK] Sign up request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), FacebookSignUpRequestKey, card)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
