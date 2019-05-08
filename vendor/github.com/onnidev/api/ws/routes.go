package ws

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// Routes TODO: NEEDS COMMENT INFO
func Routes(r chi.Router) {
	hub := DashboardHub
	go hub.Run()
	r.Route("/ws/dashboard", func(n chi.Router) {
		n.HandleFunc("/", Handler(shared.RedisPrefixer("dashboard"), hub))
	})
	staff := StaffHub
	go staff.Run()
	r.Route("/ws/staff", func(n chi.Router) {
		n.HandleFunc("/", Handler(shared.RedisPrefixer("staff"), staff))
	})

	app := AppHub
	go app.Run()
	r.Route("/ws/app", func(n chi.Router) {
		n.HandleFunc("/", Handler(shared.RedisPrefixer("app"), app))
	})
	r.Route("/ws/test", func(n chi.Router) {
		n.Get("/", func(w http.ResponseWriter, r *http.Request) {

			Publish("onni/dashboard", string("test"))
			render.Status(r, http.StatusOK)
			render.JSON(w, r, 33)
		})
	})
}
