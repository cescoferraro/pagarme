package loop

import (
	"net/http"

	"github.com/pressly/chi/render"
)

type LoopInfo struct {
	Count int
}

// Info TODO: NEEDS COMMENT INFO
func Info(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, LoopInfo{
		Count: COUNTER,
	})
}
