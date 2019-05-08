package invitedCustomer

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/shared"
	"github.com/spf13/viper"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// FB TODO: NEEDS COMMENT INFO
func FB(w http.ResponseWriter, r *http.Request) {
	oauthConf := &oauth2.Config{
		ClientID:     "450924951907199",
		ClientSecret: "4d6657f8e7649d5ad71784d56740761d",
		RedirectURL:  redirect(),
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
	oauthStateString := chi.URLParam(r, "customerId")
	url, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	parameters := url.Query()
	parameters.Set("client_id", oauthConf.ClientID)
	parameters.Set("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Set("redirect_uri", oauthConf.RedirectURL)
	parameters.Set("response_type", "code")
	parameters.Set("state", oauthStateString)
	url.RawQuery = parameters.Encode()
	log.Println(url.String())
	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
}

func redirect() string {
	if viper.GetString("env") == "homolog" {
		log.Println("homolog")
		return "https://canary.onni.live/invitedCustomer/return"
	}
	if viper.GetString("env") == "production" {
		log.Println("production")
		return "https://api.onni.live/invitedCustomer/return"
	}
	log.Println("dev")
	return "http://localhost:7000/invitedCustomer/return"
}
