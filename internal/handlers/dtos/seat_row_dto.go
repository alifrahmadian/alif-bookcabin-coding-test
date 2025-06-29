package dtos

type SeatRow struct {
	RowNumber uint     `json:"rowNumber"`
	SeatCodes []string `json:"seatCodes"`
	Seats     []Seat   `json:"seats"`
}
