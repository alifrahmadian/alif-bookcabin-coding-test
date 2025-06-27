package models

type Address struct {
	ID          int64     `json:"id"`
	PassengerID int64     `json:"passenger_id"`
	Passenger   Passenger `json:"passenger"`
	Street1     string    `json:"street1"`
	Street2     string    `json:"street2"`
	Postcode    string    `json:"postcode"`
	State       string    `json:"state"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	AddressType string    `json:"address_type"`
}
