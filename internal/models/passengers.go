package models

import "time"

type Passenger struct {
	ID                  int64           `json:"id"`
	PassengerIndex      int64           `json:"passenger_index"`
	PassengerNameNumber string          `json:"passenger_name_number"`
	PassengerDetails    PassengerDetail `json:"passenger_details"`
	PassengerInfo       PassengerInfo   `json:"passenger_info"`
	Preferences         Preference      `json:"preferences"`
	DocumentInfo        DocumentInfo    `json:"document_info"`
}

type PassengerDetail struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type PassengerInfo struct {
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Type        string    `json:"type"`
	Emails      []string  `json:"emails"`
	Phones      []string  `json:"phones"`
}

type Preference struct {
	MealPreference               string   `json:"meal_preference"`
	SeatPreference               string   `json:"seat_preference"`
	SpecialRequests              []string `json:"special_request"`
	SpecialServiceRequestRemarks []string `json:"special_service_request_remarks"`
}

type DocumentInfo struct {
	IssuingCountry string `json:"issuing_country"`
	CountryOfBirth string `json:"country_of_birth"`
	DocumentType   string `json:"document_type"`
	Nationality    string `json:"nationality"`
}
