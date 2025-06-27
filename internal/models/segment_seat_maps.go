package models

type SegmentSeatMap struct {
	ID        int64   `json:"id"`
	SegmentID int64   `json:"segment_id"`
	Segment   Segment `json:"segment"`
}
