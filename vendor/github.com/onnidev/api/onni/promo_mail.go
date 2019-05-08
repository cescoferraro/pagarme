package onni

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"strconv"

	"github.com/JKhawaja/sendinblue"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// PromoNewMail chocalate
func PromoNewMail(ctx context.Context, party types.Party, partyProduct types.PartyProduct, prom types.PromotionalCustomer, email string) {
	sibClient, err := sib.NewClient(viper.GetString("sendinblue"))
	if err != nil {
		log.Println(err)
		return
	}
	buf, err := PromoHTMLGenerator(party, partyProduct, prom.CustomerID.Hex())
	if err != nil {
		log.Println(err)
		return
	}
	myTemplate := &sib.Template{
		Template_name: "Party Report Excel",
		Html_content:  buf.String(),
		Subject:       "Você recebeu uma promoção da ONNi para a " + party.Name,
		From_email:    "onni@onni.live",
		Status:        1}
	createResponse, err := sibClient.CreateTemplate(myTemplate)
	if err != nil {
		log.Println(err)
		return
	}
	options := sib.NewEmailOptions("", "", []string{""}, []string{""})
	log.Println("sending email to ", email)
	userList := []string{email}
	_, err = sibClient.SendTemplateEmail(createResponse.Data.ID, userList, options)
	if err != nil {
		log.Println(err)
		return
	}

}

// PromoHTMLGenerator return a html template
func PromoHTMLGenerator(party types.Party, partyProduct types.PartyProduct, id string) (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := shared.Asset("promotionalCustomer/templates/promo.html")
	if err != nil {
		return response, err
	}
	tmpl := template.New("HHH")
	tmpl, err = tmpl.Parse(string(bb))
	if err != nil {
		return response, err
	}

	date := party.StartDate.Time()
	actual := struct {
		ClubName    string
		ProductName string
		PartyName   string
		PartyDate   string
		// VoucherType    string
		WebURL         string
		BaseURL        string
		CustomerID     string
		PartyID        string
		PartyProductID string
	}{
		ClubName:    party.Club.Name,
		WebURL:      webURL(),
		ProductName: partyProduct.Name,
		PartyName:   party.Name,
		BaseURL:     baseURL(),
		// VoucherType:    transformProductType(req.Type),
		CustomerID:     id,
		PartyID:        party.ID.Hex(),
		PartyProductID: partyProduct.ID.Hex(),
		PartyDate:      strconv.Itoa(date.Day()) + "/" + strconv.Itoa(int(date.Month())) + " às " + strconv.Itoa(date.Hour()) + ":" + strconv.Itoa(date.Minute()),
	}
	tmpl.Execute(response, actual)
	return response, nil
}
