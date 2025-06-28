package dtos

type Passenger struct {
	PassengerIndex      int64            `json:"passengerIndex"`
	PassengerNameNumber string           `json:"passengerNameNumber"`
	PassengerDetails    PassengerDetails `json:"passengerDetails"`
	PassengerInfo       PassengerInfo    `json:"passengerInfo"`
	Preferences         Preference       `json:"preferences"`
	DocumentInfo        DocumentInfo     `json:"documentInfo"`
}

type PassengerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type PassengerInfo struct {
	DateOfBirth string   `json:"dateOfBirth"`
	Gender      string   `json:"gender"`
	Type        string   `json:"type"`
	Emails      []string `json:"emails"`
	Phones      []string `json:"phones"`
	Address     Address  `json:"address"`
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
	SpecialPreferences SpecialPreference `json:"specialPreferences"`
	FrequentFlyer      []FrequentFlyer   `json:"frequentFlyer"`
}

type SpecialPreference struct {
	MealPreference               string   `json:"mealPreference"`
	SeatPreference               string   `json:"seatPreference"`
	SpecialRequests              []string `json:"specialRequests"`
	SpecialServiceRequestRemarks []string `json:"specialServiceRequestRemarks"`
}

type DocumentInfo struct {
	IssuingCountry string `json:"issuingCountry"`
	CountryOfBirth string `json:"countryOfBirth"`
	DocumentType   string `json:"documentType"`
	Nationality    string `json:"nationality"`
}
