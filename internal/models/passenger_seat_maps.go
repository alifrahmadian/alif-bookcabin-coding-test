package models

type PassengerSeatMap struct {
	ID                         int64     `json:"id"`
	SegmentID                  int64     `json:"segmentId"`
	Segment                    Segment   `json:"segment"`
	SeatSelectionEnabledForPax bool      `json:"seatSelectionEnabledForPax"`
	PassengerID                int64     `json:"passengerId"`
	Passenger                  Passenger `json:"passenger"`
}
