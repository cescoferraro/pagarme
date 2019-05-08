package promotionalCustomer

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/JKhawaja/sendinblue"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Create is a comented function
func Create(w http.ResponseWriter, r *http.Request) {
	req, ok := r.Context().Value(middlewares.PromotionalCustomerPostQueryKey).(types.PromotionalCustomerPost)
	if !ok {
		err := errors.New("1bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	repo, ok := r.Context().Value(middlewares.PromotionalCustomerRepoKey).(interfaces.PromotionalCustomerRepo)
	if !ok {
		err := errors.New("2bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	userClub, ok := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	if !ok {
		err := errors.New("3bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	customersRepo, ok := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	if !ok {
		err := errors.New("4bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	productsCollection, ok := r.Context().Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("4bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	nocustomer := []string{}
	already := []string{}
	promotionalCustomers := []types.PromotionalCustomer{}
	for _, promotion := range req.PromotionIDS {
		partyP, promo, err := productsCollection.GetPromotion(promotion)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		party, err := onni.Party(r.Context(), partyP.PartyID.Hex())
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		log.Println(party.ID.Hex())
		for _, mail := range req.Mails {
			email := shared.NormalizeEmail(mail)
			customer, err := customersRepo.GetByEmail(email)
			if err != nil {
				if err.Error() == "not found" {
					log.Println("not found this email")
					// ver se tem invitecustomer ...
					inviterepo, ok := r.Context().Value(middlewares.InvitedCustomerRepoKey).(interfaces.InvitedCustomerRepo)
					if !ok {
						log.Println("bug problem")
						nocustomer = append(nocustomer, promo.Name+" - "+email)
						continue
					}
					invite, err := inviterepo.GetByMail(email)
					if err != nil {
						log.Println("nbow achei por emIl")
						nocustomer = append(nocustomer, promo.Name+" - "+email)
						continue
					}
					log.Println("found an invivte for this email")
					if invite.Done {
						// talvez eu mando
						if invite.LinkedCustomer != nil {
							id := *invite.LinkedCustomer
							customer, err := customersRepo.GetByID(id.Hex())
							if err != nil {
								log.Println("nao peguei por idd linked")
								nocustomer = append(nocustomer, promo.Name+" - "+email)
								continue
							}
							exists, err := repo.CustomerHasThisPromotion(bson.ObjectIdHex(promotion), customer.ID.Hex())
							if err != nil {
								continue
							}
							if exists {
								already = append(already, promo.Name+" - "+mail)
								continue
							}
							now := types.Timestamp(time.Now())
							prom := types.PromotionalCustomer{
								ID:           bson.NewObjectId(),
								CreationDate: &now,
								PromotionID:  bson.ObjectIdHex(promotion),
								PromoterID:   userClub.ID,
								CustomerID:   customer.ID,
								CustomerMail: customer.Mail,
								CustomerName: customer.FirstName + " " + customer.LastName,
								PromoterName: userClub.Name,
							}
							go sendPromotionMail(party, partyP, prom, promo)
							promotionalCustomers = append(promotionalCustomers, prom)
							continue
						}
						if invite.AssignedMail != nil {
							customer, err := customersRepo.GetByEmail(*invite.AssignedMail)
							if err != nil {
								log.Println("nao peguei por idd linked")
								nocustomer = append(nocustomer, promo.Name+" - "+email)
								continue
							}
							exists, err := repo.CustomerHasThisPromotion(bson.ObjectIdHex(promotion), customer.ID.Hex())
							if err != nil {
								continue
							}
							if exists {
								already = append(already, promo.Name+" - "+mail)
								continue
							}
							now := types.Timestamp(time.Now())
							prom := types.PromotionalCustomer{
								ID:           bson.NewObjectId(),
								CreationDate: &now,
								PromotionID:  bson.ObjectIdHex(promotion),
								PromoterID:   userClub.ID,
								CustomerID:   customer.ID,
								CustomerMail: customer.Mail,
								CustomerName: customer.FirstName + " " + customer.LastName,
								PromoterName: userClub.Name,
							}
							go sendPromotionMail(party, partyP, prom, promo)
							promotionalCustomers = append(promotionalCustomers, prom)
							continue
						}
						continue
					}
					exists, err := repo.MailHasThisPromotion(promo.ID, email)
					if err != nil {
						continue
					}
					if exists {
						already = append(already, promo.Name+" - "+mail)
						continue
					}
					nocustomer = append(nocustomer, promo.Name+" - "+email)
					continue
				}
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			exists, err := repo.CustomerHasThisPromotion(bson.ObjectIdHex(promotion), customer.ID.Hex())
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			if !exists {
				now := types.Timestamp(time.Now())
				prom := types.PromotionalCustomer{
					ID:           bson.NewObjectId(),
					CreationDate: &now,
					PromotionID:  bson.ObjectIdHex(promotion),
					PromoterID:   userClub.ID,
					CustomerID:   customer.ID,
					CustomerMail: customer.Mail,
					CustomerName: customer.FirstName + " " + customer.LastName,
					PromoterName: userClub.Name,
				}
				go sendPromotionMail(party, partyP, prom, promo)
				promotionalCustomers = append(promotionalCustomers, prom)
				continue
			}
			already = append(already, promo.Name+" - "+mail)
			continue
		}
	}

	for _, prom := range promotionalCustomers {
		// send email
		err := repo.Collection.Insert(prom)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
	}

	if len(nocustomer) != 0 || len(already) != 0 {
		render.Status(r, 206)
	} else {
		render.Status(r, 200)
	}
	render.JSON(w, r, struct {
		MailsAlreadyAdded       []string `json:"mailsAlreadyAdded"`
		MailsNotFounds          []string `json:"mailsNotFounds"`
		PromotionsNotPermission []string `json:"promotionsNotPermission"`
	}{
		MailsAlreadyAdded:       already,
		MailsNotFounds:          nocustomer,
		PromotionsNotPermission: []string{},
	})
}

func sendPromotionMail(party types.Party, partyP types.PartyProduct, promoCus types.PromotionalCustomer, promotion types.Promotion) {
	sibClient, err := sib.NewClient(viper.GetString("sendinblue"))
	if err != nil {
		log.Println(err)
		return
	}
	buf, err := PromotionHTMLGenerator(party, partyP, promotion)
	if err != nil {
		log.Println(err)
		return
	}
	myTemplate := &sib.Template{
		Template_name: "promotion template",
		Html_content:  buf.String(),
		Subject:       "Você ganhou uma promoção",
		From_email:    "onni@onni.live",
		Status:        1}
	createResponse, err := sibClient.CreateTemplate(myTemplate)
	if err != nil {
		log.Println(err)
		return
	}
	options := sib.NewEmailOptions("", "", []string{""}, []string{""})
	userList := []string{promoCus.CustomerMail}
	_, err = sibClient.SendTemplateEmail(createResponse.Data.ID, userList, options)
	if err != nil {
		log.Println(err)
		return
	}

}

// PromotionHTMLGenerator return a html template
func PromotionHTMLGenerator(party types.Party, partyP types.PartyProduct, promotion types.Promotion) (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := shared.Asset("templates/promotion.html")
	if err != nil {
		return response, err
	}
	tmpl := template.New("HHH")
	tmpl, err = tmpl.Parse(string(bb))
	if err != nil {
		return response, err
	}
	date := party.StartDate.Time()
	resp := struct {
		PartyName        string
		ProductName      string
		ValorOriginal    string
		PartyStartDate   string
		ValorComDesconto string
	}{
		PartyName:        party.Name,
		ProductName:      partyP.Name,
		PartyStartDate:   strconv.Itoa(date.Day()) + " " + shared.Month(int(date.Month())) + " às " + strconv.Itoa(date.Hour()) + ":" + min(strconv.Itoa(date.Minute())),
		ValorOriginal:    shared.AddCents(strconv.FormatFloat(partyP.MoneyAmount.Value, 'f', -1, 64)),
		ValorComDesconto: shared.AddCents(strconv.FormatFloat(promotion.Price.Value, 'f', -1, 64)),
	}
	tmpl.Execute(response, resp)
	return response, nil
}

func min(time string) string {
	if time == "0" {
		return "00"
	}
	return time
}
