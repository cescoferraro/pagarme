package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Sms TODO: NEEDS COMMENT INFO
func Sms(w http.ResponseWriter, r *http.Request) {
	phone := chi.URLParam(r, "phone")
	password := shared.RangeIn(1000, 9999)
	phone = strings.Replace(phone, " ", "", -1)
	log.Println("printing the phone: ", phone)
	err := onni.SendSMSTotalInvoice(phone, password)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.JSON(w, r, types.SMSResponse{Password: password})
}

// Twilio TODO: NEEDS COMMENT INFO
func Twilio(w http.ResponseWriter, r *http.Request) {
	phone := chi.URLParam(r, "phone")
	password := shared.RangeIn(1000, 9999)
	phone = strings.Replace(phone, " ", "", -1)
	log.Println("[TWILIO] printing the phone: ", phone)
	err := onni.SendSMSTwillio(phone, password)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.JSON(w, r, types.SMSResponse{Password: password})
}

// TotalVoice TODO: NEEDS COMMENT INFO
func TotalVoice(w http.ResponseWriter, r *http.Request) {
	phone := chi.URLParam(r, "phone")
	password := shared.RangeIn(1000, 9999)
	phone = strings.Replace(phone, " ", "", -1)
	log.Println("[TOTALVOICE] printing the phone: ", phone)
	err := onni.SendSMSTotalInvoice(phone, password)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.JSON(w, r, types.SMSResponse{Password: password})
}
