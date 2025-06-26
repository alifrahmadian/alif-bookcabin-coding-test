package models

import "time"

type Passenger struct {
	ID                  int64           `json:"id"`
	PassengerIndex      int64           `json:"passengerIndex"`
	PassengerNameNumber string          `json:"passengerNameNumber"`
	PassengerDetails    PassengerDetail `json:"passengerDetails"`
	PassengerInfo       PassengerInfo   `json:"passengerInfo"`
	Address             Address         `json:"address"`
	Preferences         Preference      `json:"preferences"`
	FrequentFlyer       FrequentFlyer   `json:"frequentFlyer"`
	DocumentInfo        DocumentInfo    `json:"documentInfo"`
}

type PassengerDetail struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type PassengerInfo struct {
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	Type        string    `json:"type"`
	Emails      []string  `json:"emails"`
	Phones      []string  `json:"phones"`
}

type Address struct {
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	Postcode    string `json:"postcode"`
	State       string `json:"state"`
	City        string `json:"city"`
	Country     string `json:"country"`
	AddressType string `json:"addressType"`
}

type Preference struct {
	MealPreference               string   `json:"mealPreference"`
	SeatPreference               string   `json:"seatPreference"`
	SpecialRequests              []string `json:"specialRequests"`
	SpecialServiceRequestRemarks []string `json:"specialServiceRequestRemarks"`
}

type FrequentFlyer struct {
	Airline    string `json:"airline"`
	Number     string `json:"number"`
	TierNumber int    `json:"tierNumber"`
}

type DocumentInfo struct {
	IssuingCountry string `json:"issuingCountry"`
	CountryOfBirth string `json:"countryOfBirth"`
	DocumentType   string `json:"documentType"`
	Nationality    string `json:"nationality"`
}
