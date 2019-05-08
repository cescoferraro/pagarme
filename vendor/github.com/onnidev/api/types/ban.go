package types

import (
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Ban skjdnkk
type Ban struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	Type         string        `bson:"type" json:"type"`
	Payload      string        `bson:"payload" json:"payload"`
	Occurrences  *[]Occurrence `json:"occurrences,omitempty" bson:"occurrences,omitempty"`
}

// Occurrence TODO: NEEDS COMMENT INFO
type Occurrence struct {
	Date       *Timestamp    `bson:"date,omitempty" json:"date,omitempty"`
	Log        *Log          `json:"log,omitempty" bson:"log,omitempty"`
	CustomerID bson.ObjectId `json:"customerId,omitempty" bson:"customerId,omitempty"`
	UserClub   *UserClub     `json:"userClub,omitempty" bson:"userClub,omitempty"`
}

// GenerateOccurences TODO: NEEDS COMMENT INFO
func (ban Ban) GenerateOccurences(customer Customer, buylog *Log, userClub *UserClub) Ban {
	occurence := []Occurrence{ban.GetInfo(buylog, userClub, customer)}
	ban.Occurrences = &occurence
	return ban

}

// GetInfo TODO: NEEDS COMMENT INFO
func (ban Ban) GetInfo(buylog *Log, userClub *UserClub, customer Customer) Occurrence {
	now := Timestamp(time.Now())
	result := Occurrence{
		Date:       &now,
		CustomerID: customer.ID,
	}
	if buylog.DeviceSO != "" {
		log.Println("log not nil")
		result.Log = buylog
	}
	if userClub.ID.Hex() != "" {
		log.Println("userclub  not nil")
		result.UserClub = userClub
	}
	return result

}

// AntiTheftModel TODO: NEEDS COMMENT INFO
type AntiTheftModel struct {
	Picture  int     `json:"picture"`
	CardL1   int     `json:"card1"`
	CardL2   int     `json:"card2"`
	CardL3   int     `json:"card3"`
	DrinkL1  float64 `json:"drink1"`
	DrinkL2  float64 `json:"drink2"`
	DrinkL3  float64 `json:"drink3"`
	DrinkL4  float64 `json:"drink4"`
	DrinkL5  float64 `json:"drink5"`
	TicketL1 float64 `json:"ticket1"`
	TicketL2 float64 `json:"ticket2"`
	TicketL3 float64 `json:"ticket3"`
	TicketL4 float64 `json:"ticket4"`
	TicketL5 float64 `json:"ticket5"`
}

// AntiTheftModelONNi TODO: NEEDS COMMENT INFO
var AntiTheftModelONNi = AntiTheftModel{
	Picture:  8,
	CardL1:   4,
	CardL2:   5,
	CardL3:   6,
	DrinkL1:  float64(2.0),
	DrinkL2:  float64(2.5),
	DrinkL3:  float64(3.0),
	DrinkL4:  float64(3.5),
	DrinkL5:  float64(4.0),
	TicketL1: float64(2.0),
	TicketL2: float64(2.5),
	TicketL3: float64(3.0),
	TicketL4: float64(3.5),
	TicketL5: float64(4.0),
}
