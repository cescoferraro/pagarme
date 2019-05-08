package onni

import (
	"bytes"
	"html/template"
	"log"
	"math/rand"
	"time"

	"github.com/JKhawaja/sendinblue"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// MailNewUserClub chocalate
func MailNewUserClub(club types.Club, userclub types.UserClub, password string) error {
	rand.Seed(time.Now().UTC().UnixNano())
	sibClient, err := sib.NewClient(viper.GetString("sendinblue"))
	if err != nil {
		return err
	}
	buf, err := UserClubReportHTMLGenerator(userclub, club, password)
	if err != nil {
		return err
	}
	myTemplate := &sib.Template{
		Template_name: "New UserClub Mail Template",
		Html_content:  buf.String(),
		Subject:       "Bem vindo à ONNi",
		From_email:    "onni@onni.live",
		Status:        1}
	createResponse, err := sibClient.CreateTemplate(myTemplate)
	if err != nil {
		return err
	}
	options := sib.NewEmailOptions("", "", []string{""}, []string{""})
	userList := []string{userclub.Mail}
	_, err = sibClient.SendTemplateEmail(createResponse.Data.ID, userList, options)
	if err != nil {
		return err
	}
	return nil

}

// UserClubReportHTMLGenerator TODO: NEEDS COMMENT INFO
func UserClubReportHTMLGenerator(userclub types.UserClub, club types.Club, password string) (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := UserClubReportHTMLByte(userclub)
	if err != nil {
		return response, err
	}
	tmpl := template.New("HHH")
	tmpl, err = tmpl.Parse(string(bb))
	if err != nil {
		return response, err
	}
	actual := struct {
		ClubName     string
		Mail         string
		UserClubName string
		Password     string
	}{
		ClubName:     club.Name,
		UserClubName: userclub.Name,
		Password:     password,
		Mail:         userclub.Mail,
		// PartyID:        party.ID.Hex(),
		// PartyProductID: product.ID.Hex(),
		// PartyDate:      party.StartDate.Time().String(),
	}
	tmpl.Execute(response, actual)
	return response, nil
}

// UserClubReportHTMLByte return a html template
func UserClubReportHTMLByte(userclub types.UserClub) ([]byte, error) {
	if userclub.Profile == "PROMOTER" {
		log.Println("promoterr mail")
		bb, err := shared.Asset("userClub/templates/welcome_userclub_promoter.html")
		if err != nil {
			return bb, err
		}
		return bb, nil
	}
	if userclub.Profile == "ATTENDANT" {
		log.Println("attendend mail")
		bb, err := shared.Asset("userClub/templates/welcome_userclub_attendant.html")
		if err != nil {
			return bb, err
		}
		return bb, nil
	}
	log.Println("others mail")
	bb, err := shared.Asset("userClub/templates/welcome_userclub.html")
	if err != nil {
		return bb, err
	}
	return bb, nil
}
