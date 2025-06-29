package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/models"
)

type PassengerSeatMapRepository interface {
	GetPassengerSeatMapsBySegmentSeatMapID(id int64) ([]*models.PassengerSeatMap, error)
}

type passengerSeatMapRepository struct {
	DB *sql.DB
}

func NewPassengerSeatMapRepository(db *sql.DB) PassengerSeatMapRepository {
	return &passengerSeatMapRepository{
		DB: db,
	}
}

func (r *passengerSeatMapRepository) GetPassengerSeatMapsBySegmentSeatMapID(id int64) ([]*models.PassengerSeatMap, error) {
	var passengerSeatMaps []*models.PassengerSeatMap

	query := `
		SELECT
			id,
			seat_selection_enabled_for_pax,
			passenger_id,
			segment_seat_map_id,
			seat_map_id
		FROM
			passenger_seat_maps
		WHERE
			segment_seat_map_id = $1
	`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		passengerSeatMap := &models.PassengerSeatMap{}
		err := rows.Scan(
			&passengerSeatMap.ID,
			&passengerSeatMap.SeatSelectionEnabledForPax,
			&passengerSeatMap.PassengerID,
			&passengerSeatMap.SegmentSeatMapID,
			&passengerSeatMap.SeatMapID,
		)
		if err != nil {
			return nil, err
		}

		passengerSeatMaps = append(passengerSeatMaps, passengerSeatMap)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return passengerSeatMaps, nil
}
