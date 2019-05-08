package club_test

import (
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/club"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/userClub"
	"github.com/spf13/cobra"
)

// TestSpec isk kjsdf
func TestSpecReturn(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	club.Routes(r)
	userClub.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoUserClub(t)
	clubs := tester.ListAllClubs(t, loginHelper)
	id := tester.RandomClubID(clubs)
	_ = tester.ReadClub(t, loginHelper, id)
	_ = tester.ReadClubDashboard(t, loginHelper, id)
	_ = tester.ReadClubImage(t, loginHelper, id)

}
