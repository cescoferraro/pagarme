package voucher

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"log"

	"github.com/JKhawaja/sendinblue"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// MailPartyReport chocalate
func MailPartyReport(userClub types.UserClub, party types.Party, attachmentBuffer *bytes.Buffer) {
	log.Println(userClub, party.Name)
	sibClient, err := sib.NewClient(viper.GetString("sendinblue"))
	if err != nil {
		log.Println(err)
		return
	}
	buf, err := PartyReportHTMLGenerator(party)
	if err != nil {
		log.Println(err)
		return
	}
	myTemplate := &sib.Template{
		Template_name: "Party Report Excel",
		Html_content:  buf.String(),
		Subject:       "Relatório do evento: " + party.Name,
		From_email:    "onni@onni.live",
		Status:        1}
	createResponse, err := sibClient.CreateTemplate(myTemplate)
	if err != nil {
		log.Println(err)
		return
	}
	options := sib.NewEmailOptions("", "", []string{""}, []string{""})
	options.Attachment[party.Name+".xlsx"] = base64.StdEncoding.EncodeToString(attachmentBuffer.Bytes())
	userList := []string{userClub.Mail}
	_, err = sibClient.SendTemplateEmail(createResponse.Data.ID, userList, options)
	if err != nil {
		log.Println(err)
		return
	}

}

// PartyReportHTMLGenerator return a html template
func PartyReportHTMLGenerator(party types.Party) (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := shared.Asset("report/templates/index.html")
	if err != nil {
		return response, err
	}
	tmpl := template.New("HHH")
	tmpl, err = tmpl.Parse(string(bb))
	if err != nil {
		return response, err
	}
	tmpl.Execute(response, "")
	return response, nil
}
