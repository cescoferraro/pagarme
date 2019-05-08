package voucher_test

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
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/voucher"
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
	customer.Routes(r)
	voucher.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoCustomer(t)
	t.Run("sdfdskjf", func(t *testing.T) {
		ajax := infra.Ajax{
			Path:    infra.TestServer.URL + "/voucher/customer",
			Headers: loginHelper.Headers(),
		}
		_, body := infra.NewTestRequest(t, ajax)
		var vouchers []types.CompleteVoucher
		err := json.Unmarshal(body, &vouchers)
		assert.NoError(t, err)
	})

}

// CreateVouchers TODO: NEEDS COMMENT INFO
func CreateVouchers(helper *types.UserClubLoginResponse, req types.VoucherPostRequest) func(t *testing.T) {
	return func(t *testing.T) {
		bolB, _ := json.Marshal(req)
		var vouchers []types.Voucher
		log.Println("trying to change a password using the token ", helper.Token)
		ajax := infra.Ajax{
			Method:  "POST",
			Body:    bytes.NewBuffer(bolB),
			Path:    infra.TestServer.URL + "/vouchers",
			Headers: map[string]string{"JWT_TOKEN": helper.Token},
		}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &vouchers)
		assert.NoError(t, err)
	}
}
