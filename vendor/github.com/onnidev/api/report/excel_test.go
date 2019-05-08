package report_test

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/report"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/userClub"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestSpec(t *testing.T) {
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	report.Routes(r)
	userClub.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInGmailCescoUserClub(t)
	partyID := "5a9ebbddcc922d68ac5b0aaf"
	t.Run("ENDPOINT TEST", func(t *testing.T) {
		t.Run("Download Excel Report", func(t *testing.T) {
			ajax := infra.Ajax{
				Method:  "GET",
				Path:    infra.TestServer.URL + "/report/download/" + partyID,
				Headers: loginHelper.Headers()}
			response, _ := infra.NewTestRequest(t, ajax)
			content := response.Header["Content-Type"][0]
			assert.Equal(t, content, "application/octet-stream", "")
		})
		t.Run("Send Excel Report through Email", func(t *testing.T) {
			ajax := infra.Ajax{
				Method:  "POST",
				Path:    infra.TestServer.URL + "/report/mail/" + partyID,
				Headers: loginHelper.Headers()}
			var hey string
			_, body := infra.NewTestRequest(t, ajax)
			json.Unmarshal(body, &hey)
			log.Println(hey)
		})
	})
}
