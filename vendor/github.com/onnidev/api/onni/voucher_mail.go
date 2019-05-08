package onni

import (
	"bytes"
	"html/template"
	"log"
	"strconv"

	"github.com/JKhawaja/sendinblue"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// MailNewVoucher chocalate
func MailNewVoucher(mail string, party types.Party, product types.PartyProduct, req types.VoucherPostRequest, id string, isInvite bool) {
	sibClient, err := sib.NewClient(viper.GetString("sendinblue"))
	if err != nil {
		log.Println(err)
		return
	}
	buf, err := PartyReportHTMLGenerator(party, product, req, id, isInvite)
	if err != nil {
		log.Println(err)
		return
	}
	myTemplate := &sib.Template{
		Template_name: "Party Report Excel",
		Html_content:  buf.String(),
		Subject:       "Você recebeu um voucher da ONNi para a " + party.Name,
		From_email:    "onni@onni.live",
		Status:        1}
	createResponse, err := sibClient.CreateTemplate(myTemplate)
	if err != nil {
		log.Println(err)
		return
	}
	options := sib.NewEmailOptions("", "", []string{""}, []string{""})
	userList := []string{mail}
	_, err = sibClient.SendTemplateEmail(createResponse.Data.ID, userList, options)
	if err != nil {
		log.Println(err)
		return
	}

}

// PartyReportHTMLGenerator return a html template
func PartyReportHTMLGenerator(party types.Party, product types.PartyProduct, req types.VoucherPostRequest, id string, isInvite bool) (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := shared.Asset("voucher/templates/voucher.html")
	if err != nil {
		return response, err
	}
	tmpl := template.New("HHH")
	tmpl, err = tmpl.Parse(string(bb))
	if err != nil {
		return response, err
	}

	plural := false
	if req.Quantity > 1 {
		plural = true
	}
	date := party.StartDate.Time()
	actual := struct {
		ClubName       string
		ProductName    string
		PartyName      string
		PartyDate      string
		VoucherType    string
		VoucherQtd     int
		BaseURL        string
		CustomerID     string
		PartyID        string
		IsInvite       bool
		Plural         bool
		PartyProductID string
		WebURL         string
	}{
		ClubName:       party.Club.Name,
		WebURL:         webURL(),
		ProductName:    product.Name,
		VoucherQtd:     req.Quantity,
		Plural:         plural,
		PartyName:      party.Name,
		BaseURL:        baseURL(),
		IsInvite:       isInvite,
		VoucherType:    transformProductType(req.Type),
		CustomerID:     id,
		PartyID:        party.ID.Hex(),
		PartyProductID: product.ID.Hex(),
		PartyDate:      strconv.Itoa(date.Day()) + "/" + strconv.Itoa(int(date.Month())) + " às " + strconv.Itoa(date.Hour()) + ":" + strconv.Itoa(date.Minute()),
	}
	tmpl.Execute(response, actual)
	return response, nil
}

func transformProductType(typ string) string {
	switch typ {
	case "NORMAL":
		return "VENDA"
	case "PROMOTION":
		return "PROMOÇÃO"
	case "ANNIVERSARY":
		return "FREE ANIVERSÁRIO"
	case "FREE":
		return "CORTESIA"
	case "TRANSFERED":
		return "TRANSFERIDO"
	case "EXTERNAL_BUY":
		return "COMPRA EXTERNA"
	}
	return "ERRO EXCEL"
}

func baseURL() string {
	if viper.GetString("env") == "homolog" {
		log.Println("homolog")
		return "https://canary.onni.live"
	}
	if viper.GetString("env") == "production" {
		log.Println("production")
		return "https://api.onni.live"
	}
	return "http://localhost:7000"
}

func webURL() string {
	if viper.GetString("env") == "homolog" {
		log.Println("homolog")
		return "https://sigma.onni.live"
	}
	if viper.GetString("env") == "production" {
		log.Println("production")
		return "https://www.onni.live"
	}
	return "http://localhost:3000"
}
