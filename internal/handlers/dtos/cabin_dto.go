package dtos

type Cabin struct {
	Deck        string    `json:"deck"`
	SeatColumns []string  `json:"seatColumns"`
	SeatRows    []SeatRow `json:"seatRows"`
	FirstRow    uint      `json:"firstRow"`
	LastRow     uint      `json:"lastRow"`
}
