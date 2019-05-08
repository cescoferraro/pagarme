package banner_test

import (
	"encoding/json"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/banner"
	"github.com/onnidev/api/club"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/file"
	"github.com/onnidev/api/infra"
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
	club.Routes(r)
	party.Routes(r)
	banner.Routes(r)
	file.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()

	loginHelper := tester.LogInCescoUserClub(t)

	image := tester.UploadImage(t)

	clubs := tester.ListAllClubs(t, loginHelper)
	parties := tester.ListAllParties(t, loginHelper)

	for i := 1; i <= 10; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		clubID := clubs[rand.Intn(len(clubs))].ID.Hex()
		partyID := parties[rand.Intn(len(parties))].ID.Hex()
		tester.CreateClubBanner(t, loginHelper, image.FileID.Hex(), clubID, "CLUB")
		tester.CreateClubBanner(t, loginHelper, image.FileID.Hex(), partyID, "PARTY")
	}

	var publishedBanners []types.Banner
	t.Run("List Published Banners", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/banner",
			Headers: loginHelper.Headers(),
		}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &publishedBanners)
		assert.NoError(t, err)
	})
	// log.Println("sdkjfn")
	// log.Println(publishedBanners)
	// log.Println(len(publishedBanners))
	// id := publishedBanners[rand.Intn(len(publishedBanners))].ID.Hex()
	// t.Run("Read Banners", func(t *testing.T) {
	// 	ajax := infra.Ajax{
	// 		Method:  "GET",
	// 		Path:    infra.TestServer.URL + "/banner/" + id,
	// 		Headers: loginHelper.Headers(),
	// 	}
	// 	_, body := infra.NewTestRequest(t, ajax)
	// 	var readBanners types.Banner
	// 	log.Println(string(body))
	// 	err := json.Unmarshal(body, &readBanners)
	// 	assert.NoError(t, err)
	// })
	// t.Run("Delet Banners", func(t *testing.T) {
	// 	ajax := infra.Ajax{
	// 		Method:  "DELETE",
	// 		Path:    infra.TestServer.URL + "/banner/" + id,
	// 		Headers: loginHelper.Headers(),
	// 	}
	// 	_, body := infra.NewTestRequest(t, ajax)
	// 	var deleteID string
	// 	err := json.Unmarshal(body, &deleteID)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, id, deleteID)
	// })
}
