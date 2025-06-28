package dtos

type SeatRow struct {
	RowNumber uint     `json:"rowNumber"`
	SearCodes []string `json:"seatCodes"`
	Seats     []Seat   `json:"seats"`
}
