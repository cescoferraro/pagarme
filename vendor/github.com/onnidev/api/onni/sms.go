package onni

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/onnidev/api/types"
	"github.com/sfreiberg/gotwilio"
)

// SendSMSTotalInvoice TODO: NEEDS COMMENT INFO
func SendSMSTotalInvoice(phone, password string) error {
	sms := types.SMSRequest{Phone: phone, Msg: password}
	ibyt, err := json.Marshal(sms)
	if err != nil {
		return err
	}
	URL, err := url.Parse("https://api.totalvoice.com.br/sms")
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", URL.String(), bytes.NewReader(ibyt))
	if err != nil {
		return err
	}
	req.Header.Set("Access-Token", "6144c8b7bc5002acc08365e4a9206390")
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(byt))
	return nil
}

// SendSMSTwillio TODO: NEEDS COMMENT INFO
func SendSMSTwillio(phone, password string) error {
	accountSid := "AC7b863fc20f842970b2bfa4e351561538"
	authToken := "15acb11008885f1a34a270671c3a122c"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)
	from := "+13344599103"
	phone = strings.Replace(phone, " ", "", -1)
	to := "+55" + phone
	message := password
	_, _, err := twilio.SendSMS(from, to, message, "", "")
	return err
}
