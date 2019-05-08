package userClub_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/userClub"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// TestSpec isk kjsdf
func TestSpecReturn(t *testing.T) {
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	userClub.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	cescoAttendent := tester.LeitorLoginTestingHelperFactory(t, types.LoginRequest{
		Email:    "francescoaferraro+2@gmail.com",
		Password: "descriptor8"})
	headers := map[string]string{
		"JWT_TOKEN":    cescoAttendent.Token,
		"Content-Type": "application/json"}
	t.Run("List Voucher read by userClub", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/userClub/vouchers",
			Headers: headers,
		}
		_, body := infra.NewTestRequest(t, ajax)
		var readByMeVouchers []types.Voucher
		err := json.Unmarshal(body, &readByMeVouchers)
		assert.NoError(t, err)
	})
}
