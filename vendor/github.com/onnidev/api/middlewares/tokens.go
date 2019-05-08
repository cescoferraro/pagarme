package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/shared"
	"github.com/spf13/viper"
)

// TokensRepoKey is a context key
var TokensRepoKey TokenKey = "token-repo"

// TokenKey is a context key
type TokenKey string

// AttachTokenCollection mongo wise
func AttachTokenCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		collection, err := interfaces.NewTokenCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}

		ctx := context.WithValue(r.Context(), TokensRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// HeaderScan skjdfn
func HeaderScan(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		securityToken := "mYX5a43As?V7LGhTbtJ_KHpE4;:xGl;P=QvM0iJd2oPH5V<FIgB[hy67>u_3@[pc"
		if viper.GetString("env") == "production" {
			securityToken = "8IO6q/yI14z;NnD0?H9|S60$n3M'#LWC;L2y|O47**,t&foARU]8fW14M2R^~C8"
		}
		xauthapptoken := r.Header.Get("X-AUTH-APPLICATION-TOKEN")
		xauthtoken := r.Header.Get("X-AUTH-TOKEN")
		xclientid := r.Header.Get("X-CLIENT-ID")
		log.Println("xauthapptoken", xauthapptoken)
		log.Println("xauthtoken", xauthtoken)
		log.Println("xclientid", xclientid)
		if xauthapptoken != securityToken {
			log.Println("xauthapptoken vazio")
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		if xauthtoken == "" {
			log.Println("xauthtoken vazio")
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		if xclientid == "" {
			log.Println("xclientid vazio")
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// CheckToken skjdfn
func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		TokenTokensRepo := r.Context().Value(TokensRepoKey).(interfaces.TokensRepo)
		log.Println("header", r.Header)
		token, err := TokenTokensRepo.GetByToken(r.Header.Get("X-AUTH-TOKEN"))
		if err != nil {
			log.Println("nao achei o token")
			log.Println(err.Error())
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		if token.UserID.Hex() != r.Header.Get("X-CLIENT-ID") {
			log.Println("nao token id nao match")
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
