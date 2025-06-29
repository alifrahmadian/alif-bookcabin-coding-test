package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/models"
	"github.com/lib/pq"
)

type CabinRepository interface {
	GetCabinsByAircraft(aircraft string) ([]*models.Cabin, error)
}

type cabinRepository struct {
	DB *sql.DB
}

func NewCabinRepository(db *sql.DB) CabinRepository {
	return &cabinRepository{
		DB: db,
	}
}

func (r *cabinRepository) GetCabinsByAircraft(aircraft string) ([]*models.Cabin, error) {
	var cabins []*models.Cabin

	query := `
		SELECT
			id,
			deck,
			seat_columns,
			first_row,
			last_row,
			aircraft
		FROM
			cabins
		WHERE aircraft = $1
	`

	rows, err := r.DB.Query(query, aircraft)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		cabin := &models.Cabin{}
		err := rows.Scan(
			&cabin.ID,
			&cabin.Deck,
			pq.Array(&cabin.SeatColumns),
			&cabin.FirstRow,
			&cabin.LastRow,
			&cabin.Aircraft,
		)
		if err != nil {
			return nil, err
		}

		cabins = append(cabins, cabin)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cabins, nil
}
