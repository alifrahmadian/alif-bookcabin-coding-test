package models

type SegmentSeatMap struct {
	ID                   int64              `json:"id"`
	SeatsItineraryPartID int64              `json:"seats_itinerary_part_id"`
	SeatsItineraryPart   SeatsItineraryPart `json:"seats_itinerary_part"`
	SegmentID            int64              `json:"segment_id"`
	Segment              Segment            `json:"segment"`
}
