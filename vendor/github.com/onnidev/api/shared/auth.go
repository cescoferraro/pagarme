package shared

import (
	"net/http"

	jwtmiddleware "github.com/cescoferraro/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// RedisPrefixer dskjfn
func RedisPrefixer(channel string) string {
	prefix := "onni/"
	if viper.GetString("env") == "homolog" {
		prefix = prefix + "homolog/"
	}
	return prefix + channel
}

// JWTAuth asda
var JWTAuth = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwtsecret")), nil
	},
	Debug:         false,
	Extractor:     fromAuthHeader,
	SigningMethod: jwt.SigningMethodHS256,
})

func fromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("JWT_TOKEN")
	if authHeader == "" {
		return "", nil // No error, just no token
	}
	return authHeader, nil
}

// SamePayload helps
var SamePayload = func(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
