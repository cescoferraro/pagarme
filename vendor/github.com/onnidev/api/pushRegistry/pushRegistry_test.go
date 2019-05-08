package pushRegistry_test

import (
	"bytes"
	"encoding/json"
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
	"github.com/onnidev/api/pushRegistry"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/userClub"
	"github.com/stretchr/testify/assert"
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
	customer.Routes(r)
	pushRegistry.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()

	loginHelper := tester.LogInCescoCustomer(t)

	var pushRegistrys []types.PushRegistry
	t.Run("List PushRegistrys", func(t *testing.T) {
		ajax := infra.Ajax{
			Path:    infra.TestServer.URL + "/pushRegistry",
			Headers: loginHelper.Headers(),
		}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &pushRegistrys)
		assert.NoError(t, err)
	})
	t.Run("List PushRegistrys", func(t *testing.T) {
		pushRegistry := types.PushRegistry{
			Platform:    "ANDROID",
			DeviceUUID:  "sdjknf2398ufewjfwe",
			DeviceToken: "skdjnfsff",
		}
		push, _ := json.Marshal(pushRegistry)
		ajax := infra.Ajax{
			Method:  "POST",
			Body:    bytes.NewBuffer(push),
			Path:    infra.TestServer.URL + "/pushRegistry",
			Headers: loginHelper.Headers(),
		}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
		var resultPush types.PushRegistry
		err := json.Unmarshal(body, &resultPush)
		assert.NoError(t, err)
	})
}
