package customer

import (
	"log"
	"net/http"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// FacebookLogin sdkjfn
func FacebookLogin(w http.ResponseWriter, r *http.Request) {
	facebookLoginReq := r.Context().
		Value(middlewares.FacebookLoginRequestKey).(types.FacebookLoginRequest)
	validation, err := onni.FacebookAppValidate(facebookLoginReq.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println("******")
	log.Println("foi validado pelo facebook")
	log.Println(validation)
	log.Println(validation.Data)
	log.Println(validation.Data.UserID)
	log.Println("******")
	customer, err := onni.FindCustomerByFacebookID(r.Context(), validation.Data.UserID)
	if err != nil {
		log.Println("foi achei ess customer pelo facebook id")
		fbCustomer, err := onni.GetFacebookProfile(facebookLoginReq.AccessToken)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		log.Println("******")
		log.Println("peguei o fb profile desse customer")
		log.Println(fbCustomer)
		log.Println(fbCustomer.ID)
		log.Println("******")
		customer, err = onni.FindCustomerByMail(r.Context(), fbCustomer.Email)
		if err != nil {
			// TODO signup
			log.Println("não achei pelo email")
			http.Error(w, err.Error(), 400)
			return
		}
		log.Println("achei pelo email")
		customer, err = onni.
			SetCustomerFacebookID(r.Context(), customer.ID.Hex(), validation.Data.UserID)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
	log.Println("achei pelo fbID")
	response, err := customer.LogInCustomer()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
