package models

type SeatRow struct {
	ID        int64    `json:"id"`
	CabinID   int64    `json:"cabin_id"`
	RowNumber uint     `json:"row_number"`
	SeatCodes []string `json:"seat_codes"`
}
