package models

import "time"

type Passenger struct {
	ID                  int64  `json:"id"`
	PassengerIndex      int64  `json:"passengerIndex"`
	PassengerNameNumber string `json:"passengerNameNumber"`

	// Passenger Details Fields
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	// Passenger Info Fields
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	Type        string    `json:"type"`
	Emails      []string  `json:"emails"`
	Phones      []string  `json:"phones"`

	// Address Fields
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	Postcode    string `json:"postcode"`
	State       string `json:"state"`
	City        string `json:"city"`
	Country     string `json:"country"`
	AddressType string `json:"addressType"`

	// Preferences Fields
	MealPreference               string   `json:"mealPreference"`
	SeatPreference               string   `json:"seatPreference"`
	SpecialRequests              []string `json:"specialRequests"`
	SpecialServiceRequestRemarks []string `json:"specialServiceRequestRemarks"`

	// Frequent Flyer Fields
	Airline    string `json:"airline"`
	Number     string `json:"number"`
	TierNumber int    `json:"tierNumber"`

	// Document Info Fields
	IssuingCountry string `json:"issuingCountry"`
	CountryOfBirth string `json:"countryOfBirth"`
	DocumentType   string `json:"documentType"`
	Nationality    string `json:"nationality"`
}
