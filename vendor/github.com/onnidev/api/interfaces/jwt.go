package interfaces

import (
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
)

// IsFromAppKey sdkjfn
var IsFromAppKey key = "isapp-repo"

// IsRequestFromApp sdkjfn
func IsRequestFromApp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var hey types.CustomerClaims
		_, err := jwt.ParseWithClaims(
			r.Header.Get("JWT_TOKEN"),
			&hey,
			shared.JWTAuth.Options.ValidationKeyGetter)
		if err != nil {
			return
		}
		isApp := true
		if hey.CustomerID == "" {
			isApp = false
		}
		ctx := context.WithValue(r.Context(), IsFromAppKey, isApp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// IsRequestFromApp sdkjfn
func Anonymous(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := shared.JWTAuth.CheckJWT(w, r)

		// If there was an error, do not continue.
		if err != nil {
			return
		}
	})
}
