package customer

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/JKhawaja/sendinblue"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Reset sdkjfn
func Reset(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	reset := r.Context().Value(middlewares.ResetKey).(types.Reset)
	password := shared.RangeIn(100000, 999999)
	exists, err := repo.ExistsByKey("mail", reset.Mail)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if !exists {
		err := errors.New("customer does not exists")
		shared.MakeONNiError(w, r, 400, err)

		return
	}

	customer, err := repo.GetByEmail(reset.Mail)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"password":   shared.EncryptPassword2(password),
				"updateDate": &now,
			}},
		ReturnNew: true,
	}
	var result types.Customer
	_, err = repo.Collection.Find(bson.M{"_id": customer.ID}).Apply(change, &result)

	sibClient, err := sib.NewClient(viper.GetString("sendinblue"))
	if err != nil {
		log.Println(err)
		return
	}

	buf, err := ResetPasswordHTML(customer.FirstName, password)
	myTemplate := &sib.Template{
		Template_name: "nova Senha Excel",
		Html_content:  buf.String(),
		Subject:       "Recuperação de Senha",
		From_email:    "onni@onni.live",
		Status:        1}
	createResponse, err := sibClient.CreateTemplate(myTemplate)
	if err != nil {
		log.Println(err)
		return
	}
	userList := []string{reset.Mail}
	options := sib.NewEmailOptions("", "", []string{""}, []string{""})
	_, err = sibClient.SendTemplateEmail(createResponse.Data.ID, userList, options)
	if err != nil {
		log.Println(err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}

// ResetPasswordHTML return a html template
func ResetPasswordHTML(name, password string) (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := shared.Asset("public/reset.html")
	if err != nil {
		return response, err
	}
	tmpl := template.New("HHH")
	tmpl, err = tmpl.Parse(string(bb))
	if err != nil {
		return response, err
	}
	resp := struct {
		Name     string
		Password string
	}{Password: password, Name: name}
	tmpl.Execute(response, resp)
	return response, nil
}
