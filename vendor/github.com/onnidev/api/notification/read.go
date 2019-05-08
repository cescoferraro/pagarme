package notification

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Read sdkfjn
func Read(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().
		Value(middlewares.NotificationRepoKey).(interfaces.NotificationRepo)
	id := chi.URLParam(r, "notificationID")
	notification, err := repo.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(notification, "", "    ")
		log.Println("Notification you requested from MongoDB")
		log.Println(string(j))
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, notification)
}
