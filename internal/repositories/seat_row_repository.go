package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/models"
	"github.com/lib/pq"
)

type SeatRowRepository interface {
	GetSeatRowsByCabinID(id int64) ([]*models.SeatRow, error)
}

type seatRowRepository struct {
	DB *sql.DB
}

func NewSeatRowRepository(db *sql.DB) SeatRowRepository {
	return &seatRowRepository{
		DB: db,
	}
}

func (r *seatRowRepository) GetSeatRowsByCabinID(id int64) ([]*models.SeatRow, error) {
	var seatRows []*models.SeatRow

	query := `
		SELECT
			id,
			cabin_id,
			row_number,
			seat_codes
		FROM
			seat_rows
		WHERE cabin_id = $1
		ORDER BY row_number ASC
	`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		seatRow := &models.SeatRow{}
		err := rows.Scan(
			&seatRow.ID,
			&seatRow.CabinID,
			&seatRow.RowNumber,
			pq.Array(&seatRow.SeatCodes),
		)

		if err != nil {
			return nil, err
		}

		seatRows = append(seatRows, seatRow)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return seatRows, nil
}
