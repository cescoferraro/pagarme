package types

// Cart TODO: NEEDS COMMENT INFO
type Cart struct {
	GroupedProducts map[string][]CartPartyProduct `json:"groupedProducts" bson:"groupedProducts"`
	Promotions      []CartPartyProduct            `json:"promotions" bson:"promotions"`
	Party           SmallParty                     `json:"party" bson:"party"`
	Club            SmallClub                      `json:"club" bson:"club"`
	ShoppingCounter int                            `json:"shoppingCounter" bson:"shoppingCounter"`
	Balance         *float64                       `json:"balance" bson:"balance"`
	TicketStatus    string                         `json:"ticketStatus" bson:"ticketStatus"`
	DrinkStatus     string                         `json:"drinkStatus" bson:"drinkStatus"`
}
