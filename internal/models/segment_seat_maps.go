package models

type SegmentSeatMap struct {
	ID        int64   `json:"id"`
	SegmentID int64   `json:"segmentId"`
	Segment   Segment `json:"segment"`
}
