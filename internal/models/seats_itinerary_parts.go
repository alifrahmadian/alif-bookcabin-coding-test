package models

type SeatsItineraryPart struct {
	ID          int64     `json:"id"`
	SegmentID   int64     `json:"segment_id"`
	Segment     Segment   `json:"segment"`
	PassengerID int64     `json:"passenger_id"`
	Passenger   Passenger `json:"passenger"`
}
