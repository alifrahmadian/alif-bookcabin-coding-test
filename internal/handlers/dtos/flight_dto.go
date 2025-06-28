package dtos

type Flight struct {
	FlightNumber          uint     `json:"flightNumber"`
	OperatingFlightNumber uint     `json:"operatingFlightNumber"`
	AirlineCode           string   `json:"airlineCode"`
	OperatingAirlineCode  string   `json:"operatingAirlineCode"`
	StopAirports          []string `json:"stopAirports"`
	DepartureTerminal     string   `json:"departureTerminal"`
	ArrivalTerminal       string   `json:"arrivalTerminal"`
}
