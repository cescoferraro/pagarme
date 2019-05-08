package types

import (
	"strconv"
	"time"

	"github.com/onnidev/api/shared"

	"gopkg.in/mgo.v2/bson"
)

// BuyPost TODO: NEEDS COMMENT INFO
type BuyPost struct {
	CustomerID string                        `json:"customerId" bson:"customerId"`
	PartyID    string                        `json:"partyId" bson:"partyId"`
	ClubID     string                        `json:"clubId" bson:"clubId"`
	Products   []BuyPartyProductsItemRequest `json:"itens" json:"itens"`
	Log        Log                           `json:"log" bson:"log"`
}

// BuyPartyProductsItemRequest TODO: NEEDS COMMENT INFO
type BuyPartyProductsItemRequest struct {
	PartyProductID string `json:"partyProductId" bson:"partyProductId"`
	PromotionID    string `json:"promotionId" bson:"promotionId"`
	Quantity       int64  `json:"quantity" bson:"quantity"`
}

// BuyPartyProductsItem TODO: NEEDS COMMENT INFO
type BuyPartyProductsItem struct {
	PartyProductID string       `json:"partyProductId" bson:"partyProductId"`
	InvoiceItem    InvoiceItem  `json:"invoiceItem" bson:"invoiceItem"`
	PartyProduct   PartyProduct `json:"partyProduct" bson:"partyProduct"`
	PromotionID    string       `json:"promotionId" bson:"promotionId"`
	Promotion      *Promotion   `json:"promotion" bson:"promotion"`
	Quantity       int64        `json:"quantity" bson:"quantity"`
}

// VoucherProduct TODO: NEEDS COMMENT INFO
func (item BuyPartyProductsItem) VoucherProduct() VoucherProduct {
	horario := Timestamp(time.Now())
	result := VoucherProduct{
		Name:  item.PartyProduct.Name,
		Type:  item.PartyProduct.Type,
		Image: Image{FileID: bson.ObjectIdHex("586e3741cc922d0223ffcecc"), MimeType: "IMAGE_PNG", CreationDate: &horario},
	}
	if item.PartyProduct.Image != nil {
		result.Image = *item.PartyProduct.Image
	}
	return result
}

// PartyPPrice sadkjfn
func PartyPPrice(partyP PartyProduct, promotion *Promotion, party Party, club Club) float64 {
	if partyP.Type == "TICKET" {
		result := partyP.ValueToClubTickets(party, club) + partyP.ValueToONNiTickets(party, club)
		if promotion != nil {
			promo := *promotion
			if promo.Valid() {
				result = partyP.ValueToClubTicketsPromo(party, club, promo) + partyP.ValueToONNiTicketsPromo(party, club, promo)
			}
		}
		return floor(result)
	}
	result := partyP.ValueToClubDrinks(party, club) + partyP.ValueToONNiDrinks(party, club)
	if promotion != nil {
		promo := *promotion
		if promo.Valid() {
			result = partyP.ValueToClubDrinksPromo(party, club, promo) + partyP.ValueToONNiDrinksPromo(party, club, promo)
		}
	}
	return floor(result)
}

// Price sadkjfn
func (item BuyPartyProductsItem) Price(party Party, club Club) string {
	result := PartyPPrice(item.PartyProduct, item.Promotion, party, club)
	return shared.AddCents(strconv.FormatFloat(floor(result), 'f', -1, 64))
}

// BuyPostList TODO: NEEDS COMMENT INFO
type BuyPostList []BuyPartyProductsItem

// Tickets TODO: NEEDS COMMENT INFO
func (list BuyPostList) Tickets() BuyPostList {
	result := []BuyPartyProductsItem{}
	for _, item := range list {
		if item.PartyProduct.Type == "TICKET" {
			result = append(result, item)
		}
	}
	return BuyPostList(result)
}

// DoesNotHavePromotion TODO: NEEDS COMMENT INFO
func (list BuyPostList) DoesNotHavePromotion() BuyPostList {
	result := []BuyPartyProductsItem{}
	for _, item := range list {
		if item.Promotion == nil {
			result = append(result, item)
		}
	}
	return BuyPostList(result)
}

// HasPromotion TODO: NEEDS COMMENT INFO
func (list BuyPostList) HasPromotion() BuyPostList {
	result := []BuyPartyProductsItem{}
	for _, item := range list {
		if item.Promotion != nil {
			result = append(result, item)
		}
	}
	return BuyPostList(result)
}

// Sum TODO: NEEDS COMMENT INFO
func (list BuyPostList) Sum(party Party, club Club) float64 {
	return list.SumTickets(party, club) + list.SumDrinks(party, club)
}

// SumBuy TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumBuy(party Party, club Club) float64 {
	sum := float64(100) * (list.Sum(party, club))
	return floor(sum)
}

// SumBuyString TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumBuyString(party Party, club Club) string {
	return shared.AddCents(strconv.FormatFloat(floor(list.Sum(party, club)), 'f', -1, 64))
}

// SumTickets TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumTickets(party Party, club Club) float64 {
	return list.SumTicketsONNi(party, club) + list.SumTicketsClub(party, club)
}

// SumONNiString TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumONNiString(party Party, club Club) string {
	return strconv.FormatFloat(floor((list.SumONNi(party, club))*float64(100)), 'f', -1, 64)
}

// SumONNi TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumONNi(party Party, club Club) float64 {
	return floor((list.SumTicketsONNi(party, club) + list.SumDrinksONNi(party, club)))
}

// SumTicketsONNi TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumTicketsONNi(party Party, club Club) float64 {
	sum := float64(0)
	for _, item := range list {
		if item.PartyProduct.Type == "TICKET" {
			if item.Promotion != nil {
				promo := *item.Promotion
				value := float64(item.Quantity) * item.PartyProduct.ValueToONNiTicketsPromo(party, club, promo)
				sum = sum + value
				continue
			}
			value := float64(item.Quantity) * item.PartyProduct.ValueToONNiTickets(party, club)
			sum = sum + value
		}
	}
	return floor(sum)
}

// SumTicketsClub TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumTicketsClub(party Party, club Club) float64 {
	sum := float64(0)
	for _, item := range list {
		if item.PartyProduct.Type == "TICKET" {
			if item.Promotion != nil {
				promo := *item.Promotion
				value := float64(item.Quantity) * item.PartyProduct.ValueToClubTicketsPromo(party, club, promo)
				sum = sum + value
				continue
			}
			value := float64(item.Quantity) * item.PartyProduct.ValueToClubTickets(party, club)
			sum = sum + floor(value)
		}
	}
	return floor(sum)
}

// SumClubString TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumClubString(party Party, club Club) string {
	return strconv.FormatFloat(floor(float64(100)*(list.SumClub(party, club))), 'f', -1, 64)
}

// SumClub TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumClub(party Party, club Club) float64 {
	ticket := list.SumTicketsClub(party, club)
	drinks := list.SumDrinksClub(party, club)
	return floor(drinks + ticket)
}

// SumDrinks TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumDrinks(party Party, club Club) float64 {
	return list.SumDrinksClub(party, club) + list.SumDrinksONNi(party, club)
}

// SumDrinksONNi TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumDrinksONNi(party Party, club Club) float64 {
	sumfee := float64(0)
	for _, item := range list {
		if item.PartyProduct.Type == "DRINK" {
			if item.Promotion != nil {
				promo := *item.Promotion
				value := item.PartyProduct.ValueToONNiDrinksPromo(party, club, promo)
				sumfee = sumfee + (float64(item.Quantity) * value)
				continue
			}
			value := float64(item.Quantity) * item.PartyProduct.ValueToONNiDrinks(party, club)
			sumfee = sumfee + value
		}
	}
	return floor(sumfee)
}

// SumDrinksClub TODO: NEEDS COMMENT INFO
func (list BuyPostList) SumDrinksClub(party Party, club Club) float64 {
	sum := float64(0)
	for _, item := range list {
		if item.PartyProduct.Type == "DRINK" {
			if item.Promotion != nil {
				promo := *item.Promotion
				value := float64(item.Quantity) * item.PartyProduct.ValueToClubDrinksPromo(party, club, promo)
				sum = sum + floor(value)
				continue
			}
			cvlub := item.PartyProduct.ValueToClubDrinks(party, club)
			value := float64(item.Quantity) * cvlub
			sum = sum + floor(value)

		}
	}
	return floor(sum)
}

// InvoiceItens TODO: NEEDS COMMENT INFO
func (list BuyPostList) InvoiceItens() []InvoiceItem {
	result := []InvoiceItem{}
	for _, item := range list {
		result = append(result, item.InvoiceItem)
	}
	return result
}

// Drinks TODO: NEEDS COMMENT INFO
func (list BuyPostList) Drinks() BuyPostList {
	result := []BuyPartyProductsItem{}
	for _, item := range list {
		if item.PartyProduct.Type == "DRINK" {
			result = append(result, item)
		}
	}
	return BuyPostList(result)
}

// Size TODO: NEEDS COMMENT INFO
func (list BuyPostList) Size() int {
	return len(list)
}
