package clubLead_test

import (
	"log"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/clubLead"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/file"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
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
	clubLead.Routes(r)
	file.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	leads := tester.ListAllClubLeads(t)
	id := tester.RandomClubLeadID(leads)
	_ = tester.ReadClubLead(t, id)
	newLead := tester.CreateClubLeadRequest(t)
	createdLead := tester.CreateClubLead(t, newLead)
	log.Println(createdLead)

}
