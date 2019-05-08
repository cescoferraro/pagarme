package antitheft_test

import (
	"log"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/antitheft"
	"github.com/onnidev/api/club"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/file"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/party"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/userClub"
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
	userClub.Routes(r)
	club.Routes(r)
	party.Routes(r)
	antitheft.Routes(r)
	file.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoUserClub(t)
	t.Run("Get ANtithefts", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/antitheft",
			Headers: loginHelper.Headers(),
		}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
		log.Println(string(body))
	})
}
