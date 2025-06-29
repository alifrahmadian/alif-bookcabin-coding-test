package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/models"
)

type FrequentFlyerRepository interface {
	GetFrequentFliersByPassengerID(id int64) ([]*models.FrequentFlyer, error)
}

type frequentFlyerRepository struct {
	DB *sql.DB
}

func NewFrequentFlyerRepository(db *sql.DB) FrequentFlyerRepository {
	return &frequentFlyerRepository{
		DB: db,
	}
}

func (r *frequentFlyerRepository) GetFrequentFliersByPassengerID(id int64) ([]*models.FrequentFlyer, error) {
	var frequentFliers []*models.FrequentFlyer

	query := `
		SELECT
			id,
			passenger_id,
			airline,
			number,
			tier_number
		FROM
			frequent_fliers
		WHERE
			passenger_id = $1
	`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		frequentFlyer := &models.FrequentFlyer{}
		err := rows.Scan(
			&frequentFlyer.ID,
			&frequentFlyer.PassengerID,
			&frequentFlyer.Airline,
			&frequentFlyer.Number,
			&frequentFlyer.TierNumber,
		)

		if err != nil {
			return nil, err
		}

		frequentFliers = append(frequentFliers, frequentFlyer)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return frequentFliers, nil
}
