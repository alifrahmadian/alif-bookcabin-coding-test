package dtos

type SeatMapResponse struct {
	SeatsItineraryParts []SeatsItineraryPart `json:"seatsItineraryParts"`
	SelectedSeats       []SelectedSeat       `json:"selected_seat"`
}

type SeatsItineraryPart struct {
	SegmentSeatMaps []SegmentSeatMap `json:"segmentSeatMap"`
}

type SelectedSeat struct{}

type SegmentSeatMap struct {
	PassengerSeatMaps []PassengerSeatMap `json:"passengerSeatMaps"`
	Segment           Segment            `json:"segment"`
}

type PassengerSeatMap struct {
	SeatSelectionEnabledForPax bool      `json:"seatSelectionEnabledForPax"`
	SeatMap                    SeatMap   `json:"seatMap"`
	Passenger                  Passenger `json:"passenger"`
}

type SeatMap struct {
	RowsDisabledCauses []string `json:"rowsDisabledCauses"`
	Aircraft           string   `json:"aircraft"`
	Cabins             []Cabin  `json:"cabins"`
}
