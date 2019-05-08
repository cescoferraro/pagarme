package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// FinanceQueryReq sdfkjn
var FinanceQueryReq key = "transaction-query"

// ReadChangePasswordRequestFromBody skjdfn
func ReadFinanceQueryFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.FinanceQuery
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
		ctx := context.WithValue(r.Context(), FinanceQueryReq, card)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
