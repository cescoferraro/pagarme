package notification_test

import (
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
	"github.com/onnidev/api/file"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/notification"
	"github.com/onnidev/api/party"
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
	party.Routes(r)
	notification.Routes(r)
	file.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoUserClub(t)
	cesco := tester.LogInCescoCustomer(t)
	log.Println(loginHelper.Token)
	var allNotifications []types.Notification
	t.Run("Get Notifications", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/notification",
			Headers: loginHelper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &allNotifications)
		// log.Println(allNotifications)
		assert.NoError(t, err)
	})
	t.Run("Get Customer Notifications", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/notification/customer/" + cesco.ID,
			Headers: cesco.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		var allNotifications []types.Notification
		err := json.Unmarshal(body, &allNotifications)
		// log.Println(err.Error())
		assert.NoError(t, err)
		log.Println(allNotifications)
	})
	id := allNotifications[rand.Intn(len(allNotifications))].ID.Hex()
	single := types.Notification{}
	t.Run("Read Single Notifications", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/notification/" + id,
			Headers: loginHelper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &single)
		assert.NoError(t, err)
		assert.Equal(t, id, single.ID.Hex())
	})
	// var publishedNotification types.Notification
	// t.Run("Publish Single Notifications", func(t *testing.T) {
	// 	ajax := infra.Ajax{
	// 		Method:  "POST",
	// 		Path:    infra.TestServer.URL + "/notification/publish/" + single.ID.Hex(),
	// 		Headers: loginHelper.Headers(),
	// 	}
	// 	_, body := infra.NewTestRequest(t, ajax)
	// 	err := json.Unmarshal(body, &publishedNotification)
	// 	log.Println(string(body))
	// 	assert.NoError(t, err)
	// 	// assert.Equal(t, single.ID.Hex(), publishedNotification.ID.Hex())
	// })
}
