package location

import (
	"errors"
	"log"
	"net/http"
	"strings"

	geo "github.com/martinlindhe/google-geolocate"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Locate is commented
func Locate(w http.ResponseWriter, r *http.Request) {
	geolocationReq := r.Context().
		Value(middlewares.ReadGeoLocalizationPostKey).(types.GeoLocationPostRequest)
	client := geo.NewGoogleGeo(viper.GetString("GOOGLEMAPSTOKEN"))
	p := geo.Point{
		Lat: geolocationReq.Lat,
		Lng: geolocationReq.Long,
	}
	res, err := client.ReverseGeocode(&p)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	rr := strings.Split(res, ",")
	log.Println("*****************")
	log.Println("*****************")
	log.Println(res)
	log.Println(len(res))
	log.Println("*****************")
	log.Println("*****************")
	if len(rr) < 4 {
		http.Error(w,
			errors.New("Not the expected response from google api").Error(),
			http.StatusBadRequest)
		return
	}
	response := strings.Split(rr[2], "-")
	render.Status(r, http.StatusOK)
	render.JSON(w, r, types.GeoResponse{
		City:  strings.TrimSpace(response[0]),
		State: strings.TrimSpace(response[1]),
	})
}
