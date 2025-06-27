package models

import "time"

type Segment struct {
	ID                          int64                   `json:"id"`
	Type                        string                  `json:"@type"`
	SegmentOfferInformation     SegmentOfferInformation `json:"segment_offer_information"`
	Duration                    uint                    `json:"duration"`
	CabinClass                  string                  `json:"cabin_class"`
	Equipment                   string                  `json:"equipment"`
	FlightID                    int64                   `json:"flight_id"`
	Flight                      Flight                  `json:"flight"`
	Origin                      string                  `json:"origin"`
	Destination                 string                  `json:"destination"`
	Departure                   time.Time               `json:"departure"`
	Arrival                     time.Time               `json:"arrival"`
	BookingClass                string                  `json:"booking_class"`
	LayoverDuration             uint                    `json:"layover_duration"`
	FareBasis                   string                  `json:"fare_basis"`
	SubjectToGovernmentApproval bool                    `json:"subject_to_government_approval"`
	SegmentRef                  string                  `json:"segment_ref"`
}

type SegmentOfferInformation struct {
	FlightMiles uint64 `json:"flight_miles"`
	AwardFlight bool   `json:"award_flight"`
}
