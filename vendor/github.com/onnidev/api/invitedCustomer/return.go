package invitedCustomer

import (
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/spf13/viper"
)

// Return TODO: NEEDS COMMENT INFO
func Return(w http.ResponseWriter, r *http.Request) {
	oauthConf := &oauth2.Config{
		ClientID:     "450924951907199",
		ClientSecret: "4d6657f8e7649d5ad71784d56740761d",
		RedirectURL:  redirect(),
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
	code := r.FormValue("code")

	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	validation, err := onni.FacebookAppValidate(token.AccessToken)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	exists, err := onni.ExistCustomerWithFacebookID(r.Context(), validation.Data.UserID)
	if err != nil {
		log.Println("deu pau no existssss")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(exists, "passou do exists")
	if exists {
		log.Println("existssss nas que eraaaa")
		url := returnFBError()
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		return
	}
	customer := r.FormValue("state")
	inviterepo, ok := r.Context().Value(middlewares.InvitedCustomerRepoKey).(interfaces.InvitedCustomerRepo)
	if !ok {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	invite, err := inviterepo.GetByID(customer)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	profile, err := onni.GetFacebookProfile(token.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println("email")
	log.Println(profile.Email)
	newinvite, err := inviterepo.AddFacebookID(invite.ID.Hex(), validation, profile)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	url := returnTo(newinvite.ID.Hex())
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func returnTo(customer string) string {
	if viper.GetString("env") == "homolog" {
		return "https://sigma.onni.live/invite/" + customer
	}
	if viper.GetString("env") == "production" {
		return "https://www.onni.live/invite/" + customer
	}
	return "http://localhost:3000/invite/" + customer
}

func returnFBError() string {
	if viper.GetString("env") == "homolog" {
		return "https://sigma.onni.live/error/fb"
	}
	if viper.GetString("env") == "production" {
		return "https://www.onni.live/error/fb"
	}
	return "http://localhost:3000/error/fb"
}
