package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/models"
	"github.com/lib/pq"
)

type SeatMapRepository interface {
	GetSeatMapByID(id int64) (*models.SeatMap, error)
}

type seatMapRepository struct {
	DB *sql.DB
}

func NewSeatMapRepository(db *sql.DB) SeatMapRepository {
	return &seatMapRepository{
		DB: db,
	}
}

func (r *seatMapRepository) GetSeatMapByID(id int64) (*models.SeatMap, error) {
	seatMap := &models.SeatMap{}

	query := `
		SELECT
			id,
			segment_id,
			rows_disabled_causes,
			aircraft
		FROM
			seat_maps
		WHERE
			id = $1
	`

	err := r.DB.
		QueryRow(query, id).
		Scan(
			&seatMap.ID,
			&seatMap.SegmentID,
			pq.Array(&seatMap.RowsDisabledCauses),
			&seatMap.Aircraft,
		)
	if err != nil {
		return nil, err
	}

	return seatMap, nil
}
