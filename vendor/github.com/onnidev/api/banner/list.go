package banner

import (
	"net/http"

	"github.com/onnidev/api/onni"
	"github.com/pressly/chi/render"
)

// PublishedBanners is dskjfn
func PublishedBanners(w http.ResponseWriter, r *http.Request) {
	banners, err := onni.PublishedBanners(r.Context())
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, banners)
}

// AllBanners is dskjfn
func AllBanners(w http.ResponseWriter, r *http.Request) {
	banners, err := onni.AllBanners(r.Context())
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, banners)
}
