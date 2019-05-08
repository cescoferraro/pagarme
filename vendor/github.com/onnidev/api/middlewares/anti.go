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

// AntiTheftModelRequestKey is the shit
var AntiTheftModelRequestKey key = "antitheft-model-request"

// ReadAntiTheftModelRequestFromBody skjdfn
func ReadAntiTheftModelRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var model types.AntiTheftModel
		err := decoder.Decode(&model)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(model, "", "    ")
			log.Println("[ANTITHEFT] Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), AntiTheftModelRequestKey, model)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AntiTheftRepoKey sdkfjn
var AntiTheftRepoKey key = "antitheft-repo"

// AttachAntiTheftCollection get the mogo collection
func AttachAntiTheftCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		collection, err := interfaces.NewAntiTheftCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		ctx := context.WithValue(r.Context(), AntiTheftRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
