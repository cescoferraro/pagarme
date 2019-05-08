package notification_test

import (
	"bytes"
	"encoding/json"
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
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/party"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/userClub"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	// . "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/cobra"
)

// TestSpec isk kjsdf
func TestSpecPublish(t *testing.T) {
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
	var allParties []types.Party
	t.Run("List Parties", func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/party",
			Headers: map[string]string{
				"JWT_TOKEN":    loginHelper.Token,
				"Content-Type": "application/json"}}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &allParties)
		assert.NoError(t, err)
	})

	rand.Seed(time.Now().UTC().UnixNano())
	randomParty := allParties[rand.Intn(len(allParties)-1)]
	not := types.NotificationPostRequest{
		ID:      bson.NewObjectId().Hex(),
		Title:   randomParty.Name,
		Text:    randomParty.Description,
		PartyID: randomParty.ID.Hex(),
	}
	notificationCreated := types.Notification{}
	t.Run("Create Single Notifications", func(t *testing.T) {
		bolB, _ := json.Marshal(not)
		ajax := infra.Ajax{
			Method:  "POST",
			Path:    infra.TestServer.URL + "/notification",
			Body:    bytes.NewBuffer(bolB),
			Headers: loginHelper.Headers(),
		}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &notificationCreated)
		assert.NoError(t, err)

	})

	t.Run("Publish Manual Notifications", func(t *testing.T) {
		var iosIDS []string
		// Cesco 14/05
		iosIDS = append(iosIDS, "1686539249ddeb3e255f0091bd96710b3f4b056107369a4b155cb764d8c2e68d")
		// Thales 14/05 afternoon
		iosIDS = append(iosIDS, "b2881990c3b2f4871323728bfac2ffc0e47287349443e93ce683eca8cb794891")
		// Gian 14/05 night
		iosIDS = append(iosIDS, "83a76eaa5729113516dc27b5cbb21c896b7e8ef47579a9bcfa7dcb32bb5a9e1d")
		err := onni.IosPushProductionPushNotification(notificationCreated, iosIDS)
		assert.NoError(t, err)
		ids := []string{}
		// kerpen
		ids = append(ids, "dCoRK-nXI7Y:APA91bFzdnoLm--teJ7LVAt2EbDWU9WfxIhoX5yfdZW1_ctBjdt1YsSSm3c-0mAqdfIEMFYrBorP4QELMLbNsRaFPXmdv8QMjJLqrv5zcw-RUKQIHu50cEE7LGZsIEszYUSivsoa9Cwk")
		// thales
		ids = append(ids, "fU98lpgnsgw:APA91bErCyguDyy47UnUksa1OCZ0xI4RRq_ASYv6tRPeVArR1ZwpuJV-xofQmgTXdeF_xacCpcANhNzXME3wqHUePQYKPtqMRnpUcHTEHZJEVaEpKXmNWOyFHruoyMETMm64M-IRICJP")
		_, err = onni.AndroidPushDevelopmentPushNotification(notificationCreated, ids)
		assert.NoError(t, err)
	})
	// var deleted string
	// t.Run("Delete Single Notifications", func(t *testing.T) {
	// 	ajax := infra.Ajax{
	// 		Method:  "DELETE",
	// 		Path:    infra.TestServer.URL + "/notification/" + id,
	// 		Headers: loginHelper.Headers()}
	// 	_, body := infra.NewTestRequest(t, ajax)
	// 	err := json.Unmarshal(body, &deleted)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, id, deleted)
	// })
}
