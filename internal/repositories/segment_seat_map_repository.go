package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/models"
)

type SegmentSeatMapRepository interface {
	GetSegmentSeatMapsBySeatsItineraryPartID(id int64) ([]*models.SegmentSeatMap, error)
}

type segmentSeatMapRepository struct {
	DB *sql.DB
}

func NewSegmentSeatMapRepository(db *sql.DB) SegmentSeatMapRepository {
	return &segmentSeatMapRepository{
		DB: db,
	}
}

func (r *segmentSeatMapRepository) GetSegmentSeatMapsBySeatsItineraryPartID(id int64) ([]*models.SegmentSeatMap, error) {
	var segmentSeatMaps []*models.SegmentSeatMap

	query := `
		SELECT
			id,
			seats_itinerary_part_id,
			segment_id
		FROM 
			segment_seat_maps
		WHERE
			seats_itinerary_part_id = $1
	`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		segmentSeatMap := &models.SegmentSeatMap{}
		err := rows.Scan(
			&segmentSeatMap.ID,
			&segmentSeatMap.SeatsItineraryPartID,
			&segmentSeatMap.SegmentID,
		)

		if err != nil {
			return nil, err
		}

		segmentSeatMaps = append(segmentSeatMaps, segmentSeatMap)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return segmentSeatMaps, nil
}
