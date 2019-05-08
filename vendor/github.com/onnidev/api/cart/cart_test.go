package cart_test

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/cart"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
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
	cart.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoCustomer(t)
	// cart := tester.GetCart(t, loginHelper)
	// loginHelper := tester.LogInKerpenCustomer(t)
	cart := tester.GetCart(t, loginHelper)
	// byt, err := json.MarshalIndent(cart, "", "    ")
	byt2, err := json.MarshalIndent(cart.Promotions, "", "    ")
	assert.NoError(t, err)
	log.Println(string(byt2))
}
