package models

type Flight struct {
	ID                    int64    `json:"id"`
	FlightNumber          uint     `json:"flightNumber"`
	OperatingFlightNumber uint     `json:"operatingFlightNumber"`
	AirlineCode           string   `json:"airlineCode"`
	OperatingAirlineCode  string   `json:"operatingAirlineCode"`
	StopAirports          []string `json:"stopAirports"`
	DepartureTerminal     string   `json:"departureTerminal"`
	ArrivalTerminal       string   `json:"arrivalTerminal"`
}
