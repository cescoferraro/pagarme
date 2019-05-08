package customer

import (
	"log"
	"net/http"
	"strings"

	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// Redirect TODO: NEEDS COMMENT INFO
func Redirect(w http.ResponseWriter, r *http.Request) {
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
	if !exists {
		http.Redirect(w, r, newcomer(validation.Data.UserID, r.FormValue("state")), http.StatusTemporaryRedirect)
		return
	}
	customer, err := onni.FindCustomerByFacebookID(r.Context(), validation.Data.UserID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	jwttoken, err := customer.GenerateToken()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if !customer.Ready() {
		http.Redirect(w, r, fixup(customer.ID.Hex(), jwttoken, r.FormValue("state")), http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, logger(jwttoken, r.FormValue("state")), http.StatusTemporaryRedirect)
	return
}

func fixup(id, token, togo string) string {
	token = strings.Replace(token, ".", "@", -1)
	split := strings.Split(togo, "@")
	log.Println(split)
	if shared.Contains(split, "evento") || shared.Contains(split, "label") {
		if viper.GetString("env") == "homolog" {
			return "https://sigma.onni.live/fixup/" + id + "/" + token + "/" + togo
		}
		if viper.GetString("env") == "production" {
			return "https://www.onni.live/fixup/" + id + "/" + token + "/" + togo
		}
		return "http://localhost:3000/fixup/" + id + "/" + token + "/" + togo
	}
	log.Println("http://localhost:3000/fixup/home/" + id + "/" + token)
	if viper.GetString("env") == "homolog" {
		return "https://sigma.onni.live/fixup/home/" + id + "/" + token
	}
	if viper.GetString("env") == "production" {
		return "https://www.onni.live/fixup/home/" + id + "/" + token
	}
	return "http://localhost:3000/fixup/home/" + id + "/" + token
}

func newcomer(fbid string, togo string) string {
	fbid = strings.Replace(fbid, ".", "@", -1)
	split := strings.Split(togo, "@")
	log.Println(split)
	if shared.Contains(split, "evento") || shared.Contains(split, "label") {
		if viper.GetString("env") == "homolog" {
			return "https://sigma.onni.live/newcomer/" + fbid + "/" + togo
		}
		if viper.GetString("env") == "production" {
			return "https://www.onni.live/newcomer/" + fbid + "/" + togo
		}
		return "http://localhost:3000/newcomer/" + fbid + "/" + togo
	}
	log.Println("to mandando na home")
	log.Println("to mandando na home")
	log.Println("to mandando na home")
	log.Println("to mandando na home")
	log.Println("to mandando na home")
	log.Println("http://localhost:3000/newcomer/home/" + fbid)
	if viper.GetString("env") == "homolog" {
		return "https://sigma.onni.live/newcomer/home/" + fbid
	}
	if viper.GetString("env") == "production" {
		return "https://www.onni.live/newcomer/home/" + fbid
	}
	return "http://localhost:3000/newcomer/home/" + fbid
}

func logger(token string, togo string) string {
	token = strings.Replace(token, ".", "@", -1)
	if viper.GetString("env") == "homolog" {
		return "https://sigma.onni.live/logger/" + token + "/" + togo
	}
	if viper.GetString("env") == "production" {
		return "https://www.onni.live/logger/" + token + "/" + togo
	}
	return "http://localhost:3000/logger/" + token + "/" + togo
}
