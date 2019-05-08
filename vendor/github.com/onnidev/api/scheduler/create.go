package scheduler

import (
	"net/http"

	"github.com/pressly/chi/render"
)

// CreateTask is commented
func CreateTask(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, 33)
}
