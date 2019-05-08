package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// ReadGeoLocalizationPostKey sdkjfn
var ReadGeoLocalizationPostKey key = "read-geo-post-req"

// ReadGeoLocalizationPostRequestFromBody skjdfn
func ReadGeoLocalizationPostRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var userClubReq types.GeoLocationPostRequest
		err := decoder.Decode(&userClubReq)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(userClubReq, "", "    ")
			log.Println("GeoLocalization Post Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ReadGeoLocalizationPostKey, userClubReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
