package party

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// Create TODO: NEEDS COMMENT INFO
func Create(w http.ResponseWriter, r *http.Request) {
	log.Println(3233)
	req, ok := r.Context().Value(middlewares.SoftPartyPostRequestKey).(types.SoftPartyPostRequest)
	if !ok {
		err := errors.New("assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	repo, ok := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	if !ok {
		err := errors.New("assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	all, err := onni.MusicStyles(r.Context(), req.MusicStyles)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	partyID := bson.NewObjectId()
	horario := types.Timestamp(time.Now())
	update := types.Timestamp(time.Time{})
	image := types.Image{}
	if viper.GetString("env") == "production" || viper.GetString("env") == "dev" {
		log.Println("eu vim aqui")
		image, err = onni.ImageONNi(r.Context(), "5af8c25d9305fe0001a91de4")
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
	}
	if viper.GetString("env") == "homolog" {
		image, err = onni.ImageONNi(r.Context(), "58c2eb58b80dff716c8a2f59")
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
	}
	log.Println("afer images")
	party := types.Party{
		ID:              partyID,
		CreationDate:    &horario,
		Address:         req.Address,
		Status:          "DRAFT",
		ClubID:          bson.ObjectIdHex(req.ClubID),
		BackgroundImage: image,
		UpdateDate:      &update,
		StartDate:       req.StartDate,
		MusicStyles:     all,
		EndDate:         req.EndDate,
		Name:            req.Name,
		Description:     req.Description,
		Location: types.Location{
			Type:        "Point",
			Coordinates: [2]float64{0.0, 0.0},
		},
	}
	if req.ClubMenuProduct != nil {
		log.Println("mandei cardapio de producto")
		if req.ClubMenuProduct.ID != nil {
			log.Println("mandei um id no cardapio de produto")
			log.Println("vouc UPDATE cardapio de produto")
			menu, err := onni.UpdateClubMenuProduct(r.Context(), req)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			log.Printf("bem aqui certinho")
			_, err = onni.BrandNewProductsFromMenu(r.Context(), menu, partyID)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			party.ClubMenuProductID = &menu.ID
		}
		if req.ClubMenuProduct.ID == nil {
			log.Println("NAO mandei um id no cardapio de produto")
			log.Println("vouc criar NOVO cardapio de produto")
			menu, err := onni.CreateClubMenuProduct(r.Context(), req)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			log.Println("antes de pegar partyProducts from menu")
			_, err = onni.BrandNewProductsFromMenu(r.Context(), menu, partyID)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			party.ClubMenuProductID = &menu.ID
		}
	}
	if req.ClubMenuTicket != nil {
		log.Println("##### mandei cardapio de ticket")
		if req.ClubMenuTicket.ID != nil {
			log.Println("###### updating an existing menu")
			menu, err := onni.UpdateClubMenuTicket(r.Context(), req)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			log.Println("before inserting partyProducts")
			_, err = onni.BrandNewTicketsFromMenu(r.Context(), menu, partyID)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			party.ClubMenuTicketID = &menu.ID
		}
		if req.ClubMenuTicket.ID == nil {
			menu, err := onni.CreateClubMenuTicket(r.Context(), req)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			log.Println("22before inserting partyProducts")
			_, err = onni.BrandNewTicketsFromMenu(r.Context(), menu, partyID)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			party.ClubMenuTicketID = &menu.ID
		}
	}
	err = repo.Collection.Insert(party)
	if err != nil {
		log.Println("deu erro em inserrir uma festa")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, party.ID.Hex())
}
