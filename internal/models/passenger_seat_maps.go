package models

type PassengerSeatMap struct {
	ID                         int64     `json:"id"`
	SegmentID                  int64     `json:"segment_id"`
	Segment                    Segment   `json:"segment"`
	SeatSelectionEnabledForPax bool      `json:"seat_selection_enabled_for_pax"`
	PassengerID                int64     `json:"passenger_id"`
	Passenger                  Passenger `json:"passenger"`
	SeatMapID                  int64     `json:"seat_map_id"`
	SeatMap                    SeatMap   `json:"seat_map"`
}
