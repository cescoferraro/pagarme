package proxy_test

import (
	"log"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/proxy"
	"github.com/onnidev/api/tester"
	// . "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/cobra"
)

// TestSpec isk kjsdf
func TestSpec(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	customer.Routes(r)
	proxy.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoCustomer(t)
	t.Run("Proxy request", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/proxy/app/v4/home/" + loginHelper.ID,
			Headers: loginHelper.SoftHeaders(),
		}
		log.Println(loginHelper.Token)
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
	})
	t.Run("Buy Proxy request", func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/proxy/buy",
		}
		log.Println(loginHelper.Token)
		_, body := infra.NewTestRequest(t, ajax)
		// assert.Equal(t, string(body), "buy")
		log.Println(string(body))
	})
}
