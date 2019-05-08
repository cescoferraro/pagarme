package types

// WebSocketReport TODO: NEEDS COMMENT INFO
type WebSocketReport struct {
	PartyID        string     `json:"partyID"`
	PartyAddress   Address    `json:"partyAddress" bson:"partyAddress"`
	PartyStartDate *Timestamp `json:"partyStartDate" bson:"partyStartDate"`
	PartyEndDate   *Timestamp `json:"partyEndDate" bson:"partyEndDate"`
	ClubID         string     `json:"clubID"`
	PartyName      string     `json:"partyName"`
	ClubName       string     `json:"clubName"`
	Used           int        `json:"used"`
	Available      int        `json:"available"`
}
