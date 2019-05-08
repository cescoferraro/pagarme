package types

// ClubDashboard TODO: NEEDS COMMENT INFO
type ClubDashboard struct {
	Address Address `json:"address" bson:"address"`
	Total   float64 `json:"total" bson:"total"`
}
