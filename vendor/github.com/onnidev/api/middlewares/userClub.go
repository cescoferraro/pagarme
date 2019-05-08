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

// UserClubRepoKey sdkfjn
var UserClubRepoKey key = "userclub-repo"

// AttachUserClubCollection get the mogo collection
func AttachUserClubCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		collection, err := interfaces.NewUserClubCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}

		ctx := context.WithValue(r.Context(), UserClubRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserClubReq sdfkjn
var UserClubReq key = "userClub-request"

// ReadChangePasswordRequestFromBody skjdfn
func ReadChangePasswordRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.ChangePasswordRequest
		err := decoder.Decode(&card)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(card, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), UserClubReq, card)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// SoftUserClubLoginReq sdfkjn
var SoftUserClubLoginReq key = "Softlogin-userClub"

// ReadSoftLoginRequestFromBody skjdfn
func ReadSoftLoginRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.SoftLoginRequest
		err := decoder.Decode(&card)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(card, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), SoftUserClubLoginReq, card)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserClubLoginReq sdfkjn
var UserClubLoginReq key = "login-userClub"

// ReadLoginRequestFromBody skjdfn
func ReadLoginRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.LoginRequest
		err := decoder.Decode(&card)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(card, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), UserClubLoginReq, card)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserClubOauth2MailKey sdfkjn
var UserClubOauth2MailKey key = "UserClubOauth2MailKey"

// ReadOauth2LoginEmail skjdfn
func ReadOauth2LoginEmail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var login types.Oauth2LoginRequest
		err := decoder.Decode(&login)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(login, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), UserClubOauth2MailKey, login)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserClubPatchRequestKey sdkjfn
var UserClubPatchRequestKey key = "read-userClub-post-req"

// ReadUserClubPatchRequestFromBody skjdfn
func ReadUserClubPatchRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var userClubReq types.UserClubPatch
		err := decoder.Decode(&userClubReq)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(userClubReq, "", "    ")
			log.Println("UserClub Patch Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), UserClubPatchRequestKey, userClubReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ReadUserClubPostKey sdkjfn
var ReadUserClubPostKey key = "read-userClub-post-req"

// ReadUserClubPostRequestFromBody skjdfn
func ReadUserClubPostRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var userClubReq types.UserClubPostRequest
		err := decoder.Decode(&userClubReq)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(userClubReq, "", "    ")
			log.Println("UserClub Post Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ReadUserClubPostKey, userClubReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserClubKey sdfkjn
var UserClubKey key = "userClub"

// GetUserClubFromToken get the mogo collection
func GetUserClubFromToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var hey types.MyClaimsType
		_, err := jwt.ParseWithClaims(
			r.Header.Get("JWT_TOKEN"),
			&hey,
			shared.JWTAuth.Options.ValidationKeyGetter)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		userClubCollection := r.Context().Value(UserClubRepoKey).(interfaces.UserClubRepo)
		clubUser, err := userClubCollection.GetByID(hey.ClientID)

		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		ctx := context.WithValue(r.Context(), UserClubKey, clubUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
