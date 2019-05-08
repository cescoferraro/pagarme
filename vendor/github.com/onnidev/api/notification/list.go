package notification

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"math/rand"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// ListBanners is dskj
func ListNotifications(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	notificationRepo := r.Context().
		Value(middlewares.NotificationRepoKey).(interfaces.NotificationRepo)
	allNotifications, err := notificationRepo.List()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if viper.GetBool("verbose") {
		var thing types.Notification
		var ii types.Notification
		for _, not := range allNotifications {
			if not.Type == "PROMOTION" {
				thing = not
				break
			}
		}
		for _, not := range allNotifications {
			if not.Type == "TEXT" {
				ii = not
				break
			}
		}
		log.Println("PROMOTION NOTIFICATION")
		j, _ := json.MarshalIndent(thing, "", "    ")
		log.Println(string(j))
		log.Println("TEXT NOTIFICATION")
		k, _ := json.MarshalIndent(ii, "", "    ")
		log.Println(string(k))
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allNotifications)
}
