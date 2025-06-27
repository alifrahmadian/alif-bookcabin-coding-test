package models

type Flight struct {
	ID                    int64    `json:"id"`
	FlightNumber          uint     `json:"flight_number"`
	OperatingFlightNumber uint     `json:"operating_fligth_number"`
	AirlineCode           string   `json:"airline_code"`
	OperatingAirlineCode  string   `json:"operating_airline_code"`
	StopAirports          []string `json:"stop_airports"`
	DepartureTerminal     string   `json:"departure_terminal"`
	ArrivalTerminal       string   `json:"arrival_terminal"`
}
