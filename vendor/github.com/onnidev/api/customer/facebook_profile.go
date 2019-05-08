package customer

import (
	"net/http"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// FacebookProfile sdkjfn
func FacebookProfile(w http.ResponseWriter, r *http.Request) {
	facebookLoginReq := r.Context().
		Value(middlewares.FacebookLoginRequestKey).(types.FacebookLoginRequest)
	_, err := onni.FacebookAppValidate(facebookLoginReq.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fbCustomer, err := onni.GetFacebookProfile(facebookLoginReq.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, fbCustomer)
}
