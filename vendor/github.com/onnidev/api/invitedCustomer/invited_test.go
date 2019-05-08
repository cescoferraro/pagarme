package invitedCustomer

import (
	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/card"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
	// . "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TestSpec isk kjsdf
func TestSpec(t *testing.T) {
	viper.SetDefault("verbose", true)
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	customer.Routes(r)
	card.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	_ = tester.LogInCescoCustomer(t)
}
