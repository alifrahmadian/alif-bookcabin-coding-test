package models

type SeatsItineraryPart struct {
	ID          int64     `json:"id"`
	SegmentID   int64     `json:"segmentId"`
	Segment     Segment   `json:"segment"`
	PassengerID int64     `json:"passengerId"`
	Passenger   Passenger `json:"passenger"`
}
