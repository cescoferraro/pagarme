package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// BuyPostKey is the shit
var BuyPostKey key = "cards-request-repo"

// ReadCreateBuyPostFromBody skjdfn
func ReadCreateBuyPostFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var buy types.BuyPost
		err := decoder.Decode(&buy)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(buy, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), BuyPostKey, buy)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
