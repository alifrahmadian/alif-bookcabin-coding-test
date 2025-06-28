package dtos

type FrequentFlyer struct {
	Airline    string `json:"airline"`
	Number     string `json:"number"`
	TierNumber uint   `json:"tierNumber"`
}
