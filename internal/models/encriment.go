package models

type EnrichmentData struct {
	Age         int
	Gender      string
	Nationality string
}

type AgeResponse struct {
	Age int `json:"age"`
}
type GenderResponse struct {
	Gender string `json:"gender"`
}
type NationalityResponse struct {
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}
