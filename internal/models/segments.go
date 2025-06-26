package models

import "time"

type Segment struct {
	ID                          int64                   `json:"id"`
	Type                        string                  `json:"@type"`
	SegmentOfferInformation     SegmentOfferInformation `json:"segmentOfferInformation"`
	Duration                    uint                    `json:"duration"`
	CabinClass                  string                  `json:"cabinClass"`
	Equipment                   string                  `json:"equipment"`
	Flight                      Flight                  `json:"flight"`
	Origin                      string                  `json:"origin"`
	Destination                 string                  `json:"destination"`
	Departure                   time.Time               `json:"departure"`
	Arrival                     time.Time               `json:"arrival"`
	BookingClass                string                  `json:"bookingClass"`
	LayoverDuration             uint                    `json:"layoverDuration"`
	FareBasis                   string                  `json:"fareBasis"`
	SubjectToGovernmentApproval bool                    `json:"subjectToGovernmentApproval"`
	SegmentRef                  string                  `json:"segmentRef"`
}

type SegmentOfferInformation struct {
	FlightMiles uint64 `json:"flightMiles"`
	AwardFlight bool   `json:"awardFlight"`
}

type Flight struct {
	FlightNumber          uint     `json:"flightNumber"`
	OperatingFlightNumber uint     `json:"operatingFlightNumber"`
	AirlineCode           string   `json:"airlineCode"`
	OperatingAirlineCode  string   `json:"operatingAirlineCode"`
	StopAirports          []string `json:"stopAirports"`
	DepartureTerminal     string   `json:"departureTerminal"`
	ArrivalTerminal       string   `json:"arrivalTerminal"`
}
