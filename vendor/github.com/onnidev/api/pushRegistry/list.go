package pushRegistry

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// ListPushRegistry is dskjfn
func ListPushRegistry(w http.ResponseWriter, r *http.Request) {
	bannerRepo := r.Context().Value(middlewares.PushRegistryRepoKey).(interfaces.PushRegistryRepo)
	publishdPushRegistrys, err := bannerRepo.List()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(publishdPushRegistrys[0], "", "    ")
		log.Println("All published pushRegistry")
		log.Println(string(j))
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, publishdPushRegistrys)
}
