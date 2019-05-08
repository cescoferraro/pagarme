package userClub

import (
	"errors"
	"net/http"
	"strings"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Check is the shit
func Check(w http.ResponseWriter, r *http.Request) {
	customerCheck := r.Context().Value(middlewares.CustomerCheckKey).(types.CustomerCheck)
	repo, ok := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if !shared.Contains(
		[]string{
			"firstName",
			"lastName",
			"mail",
			"phone",
			"username",
			// "documentNumber",
			"facebookId",
		}, customerCheck.Type) {
		err := errors.New("not checkable key")
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	exists, err := repo.ExistsByKey(customerCheck.Type, strings.ToLower(customerCheck.Payload))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, types.CustomerCheckResponse{Result: exists})
}
