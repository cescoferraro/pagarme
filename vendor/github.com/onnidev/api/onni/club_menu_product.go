package onni

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ClubMenuProduct TODO: NEEDS COMMENT INFO
func ClubMenuProduct(ctx context.Context, id string) (types.ClubMenuProduct, error) {
	result := types.ClubMenuProduct{}
	repo, ok := ctx.Value(middlewares.ClubMenuProductRepoKey).(interfaces.ClubMenuProductRepo)
	if !ok {
		err := errors.New("assert")
		return result, err
	}
	result, err := repo.GetByID(id)
	if err != nil {

		return result, err
	}
	return result, nil
}

// CreateClubMenuProduct TODO: NEEDS COMMENT INFO
func CreateClubMenuProduct(ctx context.Context, req types.SoftPartyPostRequest) (types.ClubMenuProduct, error) {
	log.Println("##### Creating a clubMenuProduct")
	result := types.ClubMenuProduct{}
	repo, ok := ctx.Value(middlewares.ClubMenuProductRepoKey).(interfaces.ClubMenuProductRepo)
	if !ok {
		err := errors.New("assert")
		return result, err
	}
	log.Println("creating menu club")
	newProducts, err := ConvertProductsReqToMenuProducts(ctx, req.Products)
	if err != nil {
		return result, err
	}
	now := types.Timestamp(time.Now())
	name := req.Name
	if req.ClubMenuProduct != nil {
		name = req.ClubMenuProduct.Name
	}
	zero := types.Timestamp(time.Time{})
	menu := types.ClubMenuProduct{
		ID:           bson.NewObjectId(),
		CreationDate: &now,
		UpdateDate:   &zero,
		ClubID:       bson.ObjectIdHex(req.ClubID),
		Name:         name,
		MenuDefault:  false,
		Status:       "ACTIVE",
		Products:     newProducts,
	}
	log.Println("before inserting menu")
	err = repo.Collection.Insert(menu)
	if err != nil {
		return result, err
	}
	return menu, nil
}

// UpdateClubMenuProduct TODO: NEEDS COMMENT INFO
func UpdateClubMenuProduct(ctx context.Context, req types.SoftPartyPostRequest) (types.ClubMenuProduct, error) {
	log.Println("##### Updating a clubMenuProduct")
	menuid := bson.ObjectIdHex(*req.ClubMenuProduct.ID)
	result := types.ClubMenuProduct{}
	repo, ok := ctx.Value(middlewares.ClubMenuProductRepoKey).(interfaces.ClubMenuProductRepo)
	if !ok {
		err := errors.New("assert")
		return result, err
	}
	result, err := repo.GetByID(menuid.Hex())
	if err != nil {
		return result, err
	}
	log.Println("received menu Product from Mongo id:", result.ID)
	newProducts, err := ConvertProductsReqToMenuProducts(ctx, req.Products)
	if err != nil {
		return result, err
	}
	log.Println("got the new products for the menu", result.ID)
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate": &now,
				"products":   newProducts,
				"name":       req.ClubMenuProduct.Name,
			}},
		ReturnNew: true,
	}
	patchedMenu := types.ClubMenuProduct{}
	_, err = repo.Collection.Find(bson.M{"_id": result.ID}).Apply(change, &patchedMenu)
	if err != nil {
		return patchedMenu, err
	}
	return patchedMenu, nil
}

// AddMenuProduct TODO: NEEDS COMMENT INFO
func AddMenuProduct(ctx context.Context, ticket types.SoftProductPostPartyProduct, party types.Party, req types.SoftPartyPostRequest) error {
	log.Println("##### Creating a menuProduct")
	result := types.ClubMenuProduct{}
	repo, ok := ctx.Value(middlewares.ClubMenuProductRepoKey).(interfaces.ClubMenuProductRepo)
	if !ok {
		err := errors.New("assert")
		return err
	}
	menuID := *party.ClubMenuProductID
	result, err := repo.GetByID(menuID.Hex())
	if err != nil {
		return err
	}
	now := types.Timestamp(time.Now())
	menuticket, err := SoftProductPartyProductSingleConverter(ctx, ticket)
	if err != nil {
		return err
	}
	set := bson.M{
		"updateDate": &now,
		"products":   append(result.Products, menuticket),
	}
	if req.ClubMenuProduct != nil {
		menu := *req.ClubMenuProduct
		set["name"] = menu.Name
	}
	change := mgo.Change{Update: bson.M{"$set": set}, ReturnNew: true}
	patchedMenu := types.ClubMenuProduct{}
	_, err = repo.Collection.Find(bson.M{"_id": result.ID}).Apply(change, &patchedMenu)
	if err != nil {
		return err
	}
	all, err := SingleBrandNewProduct(ctx, menuticket, party.ID)
	if err != nil {
		return err
	}
	err = InsertPartyProducts(ctx, []types.PartyProduct{all})
	if err != nil {
		return err
	}
	log.Println("CESCO:", len(patchedMenu.Products))
	return nil
}

// EditMenuProduct TODO: NEEDS COMMENT INFO
func EditMenuProduct(ctx context.Context, newProduct types.MenuProduct, party types.Party, newname string) error {
	if party.ClubMenuProductID != nil {
		result := types.ClubMenuProduct{}
		repo, ok := ctx.Value(middlewares.ClubMenuProductRepoKey).(interfaces.ClubMenuProductRepo)
		if !ok {
			err := errors.New("assert")
			return err
		}
		menuID := *party.ClubMenuProductID
		result, err := repo.GetByID(menuID.Hex())
		if err != nil {
			return err
		}
		now := types.Timestamp(time.Now())
		products := []types.MenuProduct{}
		for _, tk := range result.Products {
			if tk.ID.Hex() == newProduct.ID.Hex() {
				newProduct.Image = tk.Image
				j, _ := json.MarshalIndent(newProduct, "", "    ")
				log.Println("[TICKET]")
				log.Println(string(j))
				j, _ = json.MarshalIndent(tk, "", "    ")
				log.Println("[WAS ALREADY]")
				log.Println(string(j))
				products = append(products, newProduct)
				continue
			}
			products = append(products, tk)

		}
		log.Println("######### the new name is ", newname)
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"updateDate": &now,
					"products":   products,
					"name":       newname,
				}},
			ReturnNew: true,
		}
		patchedMenu := types.ClubMenuProduct{}
		_, err = repo.Collection.Find(bson.M{"_id": result.ID}).Apply(change, &patchedMenu)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("party is editinig a clubProduct But iit has no menu? how come?")
}

// MenuProduct TODO: NEEDS COMMENT INFO
func MenuProduct(ctx context.Context, partyID, id string) (types.MenuProduct, error) {
	repo, ok := ctx.Value(middlewares.ClubMenuProductRepoKey).(interfaces.ClubMenuProductRepo)
	if !ok {
		err := errors.New("assert")
		return types.MenuProduct{}, err
	}
	partyRepo, ok := ctx.Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	if !ok {
		err := errors.New("bug")
		return types.MenuProduct{}, err
	}
	party, err := partyRepo.GetByID(partyID)
	if err != nil {
		return types.MenuProduct{}, err
	}
	result, err := repo.GetByClub(party.ClubID.Hex())
	if err != nil {
		return types.MenuProduct{}, err
	}
	for _, menu := range result {
		for _, ticket := range menu.Products {
			if ticket.ID.Hex() == id {
				return ticket, nil
			}
		}
	}
	return types.MenuProduct{}, errors.New("not found")
}
