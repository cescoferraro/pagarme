package notification

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// Publish ksjdnf
func Publish(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notification, err := onni.GetNotification(ctx, chi.URLParam(r, "notificationID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	publishedNotification, err := onni.PublishNotification(ctx, notification)
	if err != nil {
		if err.Error() == "only draft notification can be published" {
			render.Status(r, 400)
			render.JSON(w, r, publishedNotification)
			return
		}
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, 200)
	render.JSON(w, r, publishedNotification)
}

// PublishSpecific ksjdnf
func PublishSpecific(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notification, err := onni.GetNotification(ctx, chi.URLParam(r, "notificationID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	publishedNotification, err := onni.PublishNotificationForPartyProductBuyers(ctx, notification, chi.URLParam(r, "partyProductID"))
	if err != nil {
		if err.Error() == "only draft notification can be published" {
			render.Status(r, 400)
			render.JSON(w, r, publishedNotification)
			return
		}
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, 200)
	render.JSON(w, r, publishedNotification)
}
