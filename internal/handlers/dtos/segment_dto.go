package dtos

type Segment struct {
	Type                        string                  `json:"@type"`
	SegmentOfferInformation     SegmentOfferInformation `json:"segmentOfferInformation"`
	Duration                    uint                    `json:"duration"`
	CabinClass                  string                  `json:"cabinClass"`
	Equipment                   string                  `json:"equipment"`
	Flight                      Flight                  `json:"flight"`
	Origin                      string                  `json:"string"`
	Destination                 string                  `json:"destination"`
	Departure                   string                  `json:"departure"`
	Arrival                     string                  `json:"arrival"`
	BookingClass                string                  `json:"bookingClass"`
	LayoverDuration             uint                    `json:"layoverDuration"`
	FareBasis                   string                  `json:"fareBasis"`
	SubjectToGovernmentApproval bool                    `json:"subjectToGovernmentApproval"`
	SegmentRef                  string                  `json:"segmentRef"`
}

type SegmentOfferInformation struct {
	FlightMiles uint `json:"flightMiles"`
	AwardFare   bool `json:"awardFare"`
}
