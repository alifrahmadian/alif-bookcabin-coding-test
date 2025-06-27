package models

type SeatMap struct {
	ID                 int64    `json:"id"`
	SegmentID          int64    `json:"segment_id"`
	Segment            Segment  `json:"segment"`
	RowsDisabledCauses []string `json:"rows_disabled_causes"`
	Aircraft           string   `json:"aircraft"`
}
