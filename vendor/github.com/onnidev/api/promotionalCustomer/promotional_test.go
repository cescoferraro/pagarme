package promotionalCustomer_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/promotionalCustomer"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
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
	promotionalCustomer.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoCustomer(t)
	t.Run("List PromotionalCustomers", func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/promotionalCustomer",
			Headers: map[string]string{
				"JWT_TOKEN":    loginHelper.Token,
				"Content-Type": "application/json",
			},
		}
		_, body := infra.NewTestRequest(t, ajax)
		var all []types.PromotionalCustomer
		err := json.Unmarshal(body, &all)
		assert.NoError(t, err)
	})
}
