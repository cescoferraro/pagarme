package banner

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Read sdkfjn
func Read(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().Value(middlewares.BannerRepoKey).(interfaces.BannerRepo)
	id := chi.URLParam(r, "bannerID")
	banner, err := repo.GetByID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, 22)
		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(banner, "", "    ")
		log.Println("Banner you requested from MongoDB")
		log.Println(string(j))
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, banner)
}
