package userClub_test

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/icrowley/fake"
	"github.com/onnidev/api/club"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/userClub"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

// TestSpecUserClub isk kjsdf
func TestSpecUserClub(t *testing.T) {
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	userClub.Routes(r)
	club.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoUserClub(t)
	t.Run("Change Back Password",
		userClub.ChangePassword(loginHelper,
			types.ChangePasswordRequest{
				Password:    "descriptor8",
				NewPassword: "onni",
			}))
	t.Run("Change Back Password",
		userClub.ChangePassword(loginHelper,
			types.ChangePasswordRequest{
				Password:    "onni",
				NewPassword: "descriptor8",
			}))

	t.Run("Regex", func(t *testing.T) {
		db, err := infra.Cloner()
		assert.NoError(t, err)
		collection, err := interfaces.NewUserClubCollection(db)
		assert.NoError(t, err)
		_, err = collection.FacebookExtraUserRegex("francescoaferraro@gmail.com")
		assert.NoError(t, err)
	})
	clubs := tester.ListAllClubs(t, loginHelper)

	headers := map[string]string{
		"JWT_TOKEN":    loginHelper.Token,
		"Content-Type": "application/json"}
	t.Run("Create userClub", func(t *testing.T) {
		userClubReq := types.UserClubPostRequest{
			ID:      bson.NewObjectId().Hex(),
			Name:    fake.FirstName(),
			Email:   fake.FirstName() + "@gmail.com",
			Profile: "ADMIN",
			Clubs:   []string{clubs[rand.Intn(len(clubs))].ID.Hex()},
		}
		byt, err := json.Marshal(userClubReq)
		assert.NoError(t, err)
		ajax := infra.Ajax{
			Method:  "POST",
			Path:    infra.TestServer.URL + "/userClub",
			Body:    bytes.NewBuffer(byt),
			Headers: headers,
		}
		_, body := infra.NewTestRequest(t, ajax)
		var userCreated types.UserClub
		err = json.Unmarshal(body, &userCreated)
		assert.NoError(t, err)
	})
}
