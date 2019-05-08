package antitheft

import (
	"errors"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// List is dskjfn
func List(w http.ResponseWriter, r *http.Request) {
	banners, err := onni.PendingAntithefts(r.Context())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	result := []types.FullTheft{}
	repo, ok := r.Context().Value(middlewares.BanRepoKey).(interfaces.BansRepo)
	if !ok {
		err := errors.New("sakjd")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	for _, item := range banners {
		them := types.FullTheft{
			AntiTheftResult: item,
			Banned:          false,
		}
		_, err := repo.IsCustomerBanned(item.CustomerID.Hex())
		if err != nil {
			result = append(result, them)
			continue
		}
		them.Banned = true
		result = append(result, them)
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
