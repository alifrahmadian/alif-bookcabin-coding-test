package repositories

import "database/sql"

type SeatsItineraryPartRepository interface {
	CheckIfExist(id int64) (bool, error)
}

type seatsItineraryPartRepository struct {
	DB *sql.DB
}

func NewSeatsItineraryPartRepository(db *sql.DB) SeatsItineraryPartRepository {
	return &seatsItineraryPartRepository{
		DB: db,
	}
}

func (r *seatsItineraryPartRepository) CheckIfExist(id int64) (bool, error) {
	query := `
		SELECT id FROM seats_itinerary_parts where id = $1
	`

	err := r.DB.QueryRow(query, id).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
