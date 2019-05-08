package onni

import (
	"bytes"
	"html/template"
	"log"
	"strconv"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JKhawaja/sendinblue"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// SendBuySuccessMail TODO: NEEDS COMMENT INFO
func SendBuySuccessMail(party types.Party, club types.Club, products types.BuyPostList, customer types.Customer) error {
	sibClient, err := sib.NewClient(viper.GetString("sendinblue"))
	if err != nil {
		return err
	}
	buf, err := BuySuccessMail(party, club, products)
	if err != nil {
		return err
	}
	myTemplate := &sib.Template{
		Template_name: "recibos onni",
		Html_content:  buf.String(),
		Subject:       "Recibos ONNi",
		From_email:    "onni@onni.live",
		Status:        1}
	createResponse, err := sibClient.CreateTemplate(myTemplate)
	if err != nil {
		return err
	}
	options := sib.NewEmailOptions("", "", []string{""}, []string{""})
	userList := []string{customer.Mail}
	_, err = sibClient.SendTemplateEmail(createResponse.Data.ID, userList, options)
	if err != nil {
		return err
	}
	return nil
}

// BuyMailList TODO: NEEDS COMMENT INFO
type BuyMailList struct {
	Name     string
	Quantity int64
	Price    string
}

// BuySuccessMail return a html template
func BuySuccessMail(party types.Party, club types.Club, products types.BuyPostList) (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := shared.Asset("public/buy.html")
	if err != nil {
		return response, err
	}
	tmpl := template.New("HHH")
	tmpl, err = tmpl.Parse(string(bb))
	if err != nil {
		return response, err
	}
	date := party.StartDate.Time()
	ccluburl := clubURL(party)
	log.Println(ccluburl)

	cart := []BuyMailList{}
	for _, item := range products {
		if item.Quantity != 0 {
			cart = append(cart, BuyMailList{
				Name:     item.PartyProduct.Name,
				Quantity: item.Quantity,
				Price:    item.Price(party, club),
			})
		}
	}
	tmpl.Execute(response, struct {
		Products       types.BuyPostList
		Cart           []BuyMailList
		Total          string
		PartyName      string
		PartyStartDate string
		URL            string
	}{
		Products:       products,
		Total:          products.SumBuyString(party, club),
		Cart:           cart,
		PartyName:      party.Name,
		PartyStartDate: strconv.Itoa(date.Day()) + " " + shared.Month(int(date.Month())) + " às " + strconv.Itoa(date.Hour()) + ":" + strconv.Itoa(date.Minute()),
		URL:            ccluburl,
	})
	return response, nil
}
func clubURL(party types.Party) string {
	if viper.GetString("env") == "homolog" {
		log.Println("homolog")
		return "https://canary.onni.live/club/image/" + party.ClubID.Hex()
	}
	if viper.GetString("env") == "production" {
		log.Println("production")
		return "https://api.onni.live/club/image/" + party.ClubID.Hex()
	}
	log.Println("dev")
	return "https://api.onni.live/club/image/" + party.ClubID.Hex()
}

// PartyProductsBuySuccess TODO: NEEDS COMMENT INFO
func PartyProductsBuySuccess(products types.BuyPostList) error {
	for _, product := range products {
		err := PatchPartyProductBuySuccess(product)
		if err != nil {
			return err
		}
	}
	return nil
}

// PatchPartyProductBuySuccess TODO: NEEDS COMMENT INFO
func PatchPartyProductBuySuccess(product types.BuyPartyProductsItem) error {
	db, err := infra.Cloner()
	if err != nil {
		return err
	}
	defer db.Session.Close()
	repo, err := interfaces.NewPartyProductsCollection(db)
	if err != nil {
		return err
	}
	now := types.Timestamp(time.Now())
	promotions := []types.Promotion{}
	set := bson.M{
		"updateDate":        &now,
		"quantityPurchased": product.PartyProduct.QuantityPurchased + product.Quantity,
	}
	if product.PartyProduct.PromotionalPrices != nil {
		for _, promotion := range *product.PartyProduct.PromotionalPrices {
			if product.Promotion != nil {
				promo := *product.Promotion
				if promotion.ID.Hex() == promo.ID.Hex() {
					promotion.QuantityPurchased = promotion.QuantityPurchased + product.Quantity
				}
			}
			log.Println("putting promotion append")
			promotions = append(promotions, promotion)
			continue
		}
		set["promotionalPrices"] = &promotions
	}
	change := mgo.Change{Update: bson.M{"$set": set}, ReturnNew: true}
	patchedPartyProduct := types.PartyProduct{}
	_, err = repo.Collection.Find(bson.M{"_id": product.PartyProduct.ID}).Apply(change, &patchedPartyProduct)
	if err != nil {
		return err
	}
	return nil
}
