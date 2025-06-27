package models

type Cabin struct {
	ID          int64    `json:"id"`
	SeatMapID   int64    `json:"seat_map_id"`
	Deck        string   `json:"deck"`
	SeatColumns []string `json:"seat_columns"`
	FirstRow    uint     `json:"first_row"`
	LastRow     uint     `json:"last_row"`
}
