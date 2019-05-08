package file_test

import (
	"log"
	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/file"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/userClub"
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
	userClub.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoUserClub(t)
	image := tester.UploadImage(t)
	img := tester.PublishtoS3Image(t, loginHelper, image)
	log.Println(img)
}
