package models

type FrequentFlyer struct {
	ID          int64     `json:"id"`
	PassengerID int64     `json:"passenger_id"`
	Passenger   Passenger `json:"passenger"`
	Airline     string    `json:"airline"`
	Number      string    `json:"number"`
	TierNumber  int       `json:"tier_number"`
}
