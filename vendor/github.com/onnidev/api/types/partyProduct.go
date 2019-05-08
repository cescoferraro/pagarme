package types

import (
	"log"
	"math"

	"gopkg.in/mgo.v2/bson"
)

// PartyProduct FKNJSD  iiii qnj q
type PartyProduct struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	PartyID      bson.ObjectId `json:"partyId" bson:"partyId"`

	Name        string   `json:"exhibitionName" bson:"exhibitionName"`
	Product     *Product `json:"product,omitempty" bson:"product,omitempty"`
	MoneyAmount Price    `json:"moneyAmount" bson:"moneyAmount"`

	QuantityPurchased int64 `json:"quantityPurchased" bson:"quantityPurchased"`
	QuantityFree      int64 `json:"quantityFree" bson:"quantityFree"`

	PercentCustom    *float64 `json:"percentCustom" bson:"percentCustom"`
	Image            *Image   `bson:"image" json:"image"`
	CreditConsumable *float64 `json:"creditConsumable" bson:"creditConsumable"`

	MenuTicketID      *bson.ObjectId       `json:"menuTicketId" bson:"menuTicketId"`
	MenuProductID     *bson.ObjectId       `json:"menuProductId" bson:"menuProductId"`
	DoNotSellMoreThan int64                `json:"doNotSellMoreThan" bson:"doNotSellMoreThan"`
	DoNotSellBefore   *Timestamp           `bson:"doNotSellBefore,omitempty" json:"doNotSellBefore,omitempty"`
	DoNotSellAfter    *Timestamp           `bson:"doNotSellAfter,omitempty" json:"doNotSellAfter,omitempty"`
	AcceptsEvaluation bool                 `json:"acceptsEvaluation" bson:"acceptsEvaluation"`
	Featured          bool                 `json:"featured" bson:"featured"`
	Deprecated        bool                 `json:"deprecated" bson:"deprecated"`
	InviteVendorID    *bson.ObjectId       `json:"inviteVendorId" bson:"inviteVendorId"`
	Type              string               `json:"type" bson:"type"`
	Category          *string              `json:"category" bson:"category"`
	Status            string               `json:"status" bson:"status"`
	PromotionalPrices *[]Promotion         `json:"promotionalPrices,omitempty" bson:"promotionalPrices,omitempty"`
	Combo             *[]Combo             `json:"combo" bson:"combo"`
	Batches           []ObjectPartyProduct `json:"batches" bson:"batches"`
	OwnerBatchID      *bson.ObjectId       `json:"ownerBatchId" bson:"ownerBatchId"`
}

// ValueToClubTickets TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) ValueToClubTickets(party Party, club Club) float64 {
	value := partyP.MoneyAmount.Value
	log.Println("==========================")
	log.Println("partyP.MoneyAmount.Value", value)
	log.Println("club.PercentTicket", club.PercentTicket)
	if !party.AssumeServiceFee {
		price := value
		log.Println("ValueToClubTickets", price)
		return floor(price)
	}
	log.Println("vvalue", value)
	onni := partyP.ValueToONNiTickets(party, club)
	log.Println("onni", onni)
	price := value - onni
	log.Println("ValueToClubTickets", price)
	return floor(price)
}

// ValueToONNiTickets TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) ValueToONNiTickets(party Party, club Club) float64 {
	value := partyP.MoneyAmount.Value
	log.Println("==========================")
	log.Println("partyP.MoneyAmount.Value", value)
	log.Println("club.PercentTicket", club.PercentTicket)
	price := value * ((100 - club.PercentTicket) / 100)
	log.Println("ValueToONNiTickets", price)
	return floor(price)
}

// Round returns the nearest integer, rounding ties away from zero.
func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

func floor(value float64) float64 {
	return math.Round(value*float64(100)) / float64(100)
}

// ValueToClubTicketsPromo TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) ValueToClubTicketsPromo(party Party, club Club, promotion Promotion) float64 {
	value := promotion.Price.Value
	log.Println("==========================")
	log.Println("promotion.Price.Value", value)
	log.Println("club.PercentTicket", club.PercentTicket)
	if !party.AssumeServiceFee {
		price := value
		log.Println("ValueToClubTickets", price)
		return floor(price)
	}
	price := (value - partyP.ValueToONNiTicketsPromo(party, club, promotion))
	log.Println("ValueToClubTickets", price)
	return floor(price)
}

// ValueToONNiTicketsPromo TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) ValueToONNiTicketsPromo(party Party, club Club, promotion Promotion) float64 {
	value := promotion.Price.Value
	log.Println("==========================")
	log.Println("partyP.MoneyAmount.Value", value)
	log.Println("club.PercentTicket", club.PercentTicket)
	if party.AssumeServiceFee {
		price := value * ((100 - club.PercentTicket) / 100)
		log.Println("ValueToONNiTickets", price)
		return floor(price)
	}
	price := value * ((100 - club.PercentTicket) / 100)
	log.Println("ValueToONNiTickets", price)
	return floor(price)
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

// ValueToClubDrinks TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) ValueToClubDrinks(party Party, club Club) float64 {
	price := partyP.MoneyAmount.Value - partyP.ValueToONNiDrinks(party, club)
	return floor(price)
}

// ValueToONNiDrinks TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) ValueToONNiDrinks(party Party, club Club) float64 {
	price := ((100 - club.PercentDrink) * partyP.MoneyAmount.Value) / 100
	return floor(price)
}

// ValueToClubDrinksPromo TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) ValueToClubDrinksPromo(party Party, club Club, promotion Promotion) float64 {
	price := promotion.Price.Value * 100
	onniShare := partyP.ValueToONNiDrinksPromo(party, club, promotion)
	result := promotion.Price.Value - onniShare
	log.Println("price", price)
	log.Println("onniShare", onniShare)
	log.Println("result", result)
	return floor(result)
}

// ValueToONNiDrinksPromo TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) ValueToONNiDrinksPromo(party Party, club Club, promotion Promotion) float64 {
	price := ((100 - club.PercentDrink) * promotion.Price.Value) / 100
	return floor(price)
}

// GetFee TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) GetFee(hasDiscount bool, party Party, club Club, promotion Promotion) float64 {
	if partyP.Type == "DRINK" {
		return 0.0
	}
	if party.AssumeServiceFee {
		return 0.0
	}
	if hasDiscount {
		return ((100 - club.PercentTicket) * promotion.Price.Value) / 100
	}
	return ((100 - club.PercentTicket) * partyP.MoneyAmount.Value) / 100
}

// ObjectPartyProduct TODO: NEEDS COMMENT INFO
type ObjectPartyProduct struct {
	PartyProductID bson.ObjectId `json:"partyProductId" bson:"partyProductId"`
}

// Combo TODO: NEEDS COMMENT INFO
type Combo struct {
	ProductID   bson.ObjectId `json:"productId" bson:"productId"`
	ProductName string        `json:"productName" bson:"productName"`
	Quantity    int64         `json:"quantity" bson:"quantity"`
	Image       Image         `bson:"image" json:"image"`
}
