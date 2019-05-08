package auth_test

import (
	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/auth"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/file"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
	"github.com/spf13/cobra"
)

// TestSpec isk kjsdf
func TestSpecReturn(t *testing.T) {
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	file.Routes(r)
	customer.Routes(r)
	auth.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoCustomer(t)
	_ = tester.RefreshToken(t, loginHelper)
	tester.SmsTest(t)
	tester.FacebookPictureTest(t)
	tester.PhaseTest(t)
	// customer := tester.ReadCustomer(t, loginHelper, loginHelper.ID)
	// fake := tester.ReadCustomer(t, loginHelper, "5b5a32d3be32c5000159dc90")
	// result := tester.IsFacebookImageTest(t, customer.FacebookID)
	// result1 := tester.IsFacebookImageTest(t, fake.FacebookID)
	// cus, err := onni.IsCustomerFake(customer)
	// assert.NoError(t, err)
	// cus1, err := onni.IsCustomerFake(fake)
	// assert.NoError(t, err)
	// log.Println(customer, fake, result, result1, cus, cus1)

}
