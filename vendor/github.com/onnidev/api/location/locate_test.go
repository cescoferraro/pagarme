package location_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/location"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/userClub"
	"github.com/spf13/cobra"
)

// TestSpec isk kjsdf
func TestSpecReturn(t *testing.T) {
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	userClub.Routes(r)
	location.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoUserClub(t)
	t.Run("Test Geolocation", func(t *testing.T) {
		bolB, _ := json.Marshal(
			types.GeoLocationPostRequest{
				Lat:  -30.0277,
				Long: -51.217,
			})
		ajax := infra.Ajax{
			Method: "POST",
			Path:   infra.TestServer.URL + "/locate",
			Body:   bytes.NewBuffer(bolB),
			Headers: map[string]string{
				"Content-Type": "application/json",
				"JWT_TOKEN":    loginHelper.Token},
		}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
	})
}
