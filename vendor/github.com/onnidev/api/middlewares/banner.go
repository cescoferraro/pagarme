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

// BannerRepoKey is the shit
var BannerRepoKey key = "banner-repo"

// AttachBannerCollection skjdfn
func AttachBannerCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		collection, err := interfaces.NewBannerCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}

		ctx := context.WithValue(r.Context(), BannerRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// BannerPatchRequestKey is the shit
var BannerPatchRequestKey key = "banner-patch-request"

// ReadPatchBannerRequestFromBody skjdfn
func ReadPatchBannerRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var banner types.BannerPatch
		err := decoder.Decode(&banner)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(banner, "", "    ")
			log.Println("[BANNER] Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), BannerPatchRequestKey, banner)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// BannerRequestKey is the shit
var BannerRequestKey key = "banner-request"

// ReadCreateBannerRequestFromBody skjdfn
func ReadCreateBannerRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.BannerPostRequest
		err := decoder.Decode(&card)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(card, "", "    ")
			log.Println("[BANNER] Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), BannerRequestKey, card)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
