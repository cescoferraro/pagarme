package club

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// PersistToDB TODO: NEEDS COMMENT INFO
func PersistToDB(w http.ResponseWriter, r *http.Request) {
	req, ok := r.Context().Value(middlewares.ClubPostRequestKey).(types.ClubPostRequest)
	if !ok {
		err := errors.New("new assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	id := bson.NewObjectId()
	log.Println(req.Recipient)
	log.Println(req.Recipient.BankDocumentNumber)
	recipient := types.RecipientPost{
		ClubID:     id.Hex(),
		PersonType: req.ProductionType,
		// Status
		BankCode:        req.Recipient.BankCode,
		BankBranch:      req.Recipient.BankBranch,
		BankBranchVC:    req.Recipient.BankBranchVC,
		BankAccount:     req.Recipient.BankAccount,
		BankAccountVC:   req.Recipient.BankAccountVC,
		BankAccountName: req.Recipient.BankAccountName,
		BankAccountType: req.Recipient.BankAccountType,
		DocumentNumber:  req.Recipient.BankDocumentNumber,
	}
	api := pagarme.New(viper.GetString("PAGARME"))
	pagarmerecipient, _, err := api.RecipientCreate(r.Context(), recipient)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println("before creaTING onni RECIPIENT[]")
	onniRecipient, err := onni.CreateONNiRecipient(r.Context(), recipient, pagarmerecipient)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	horario := types.Timestamp(time.Now())
	img := types.Image{
		FileID:       bson.NewObjectId(),
		MimeType:     "IMAGE_PNG",
		CreationDate: &horario,
	}
	zero := float64(0.0)
	club := types.Club{
		ID:               id,
		CreationDate:     &horario,
		Name:             req.Name,
		Description:      req.Description,
		NameSearchable:   strings.ToLower(req.Name),
		BankLegalAddress: req.Recipient.BankLegalAddress,
		Mail:             req.Mail,
		OperationType:    "ONNI_APP",
		Featured:         true,
		MusicStyles:      []types.Style{},

		AverageExpendituresProduct: &zero,
		AverageExpendituresTicket:  &zero,

		PercentDrink:  95.0,
		PercentTicket: 90.0,

		ProductionType: req.ProductionType,

		Location: types.Location{
			Type:        "Point",
			Coordinates: [2]float64{req.Latitude, req.Longitude},
		},
		Address: types.Address{
			City:    req.City,
			State:   req.State,
			Country: req.Country,
			Street:  req.Street,
			Number:  req.Number,
			Unit:    req.Unit,
		},
		Image:              img,
		BackgroundImage:    &img,
		PagarMeRecipientID: onniRecipient.ID,
		Status:             "PENDING",
		FlatProducts:       false,
		RegisterOrigin:     "SITE",
	}
	err = onni.PersistClub(r.Context(), club)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, club)
}
