package models

type PassengerSeatMap struct {
	ID                         int64          `json:"id"`
	SeatSelectionEnabledForPax bool           `json:"seat_selection_enabled_for_pax"`
	PassengerID                int64          `json:"passenger_id"`
	Passenger                  Passenger      `json:"passenger"`
	SegmentSeatMapID           int64          `json:"segment_seat_map_id"`
	SegmentSeatMap             SegmentSeatMap `json:"segment_seat_map"`
	SeatMapID                  int64          `json:"seat_map_id"`
	SeatMap                    SeatMap        `json:"seat_map"`
}
