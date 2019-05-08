package types

// GeoLocationPostRequest djfhs
type GeoLocationPostRequest struct {
	Lat  float64
	Long float64
}

// GeoResponse djfhs
type GeoResponse struct {
	State string
	City  string
}
