package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/models"
)

type PassengerRepository interface {
	GetPassengerByID(id int64) (*models.Passenger, error)
}

type passengerRepository struct {
	DB *sql.DB
}

func NewPassengerRepository(db *sql.DB) PassengerRepository {
	return &passengerRepository{
		DB: db,
	}
}

func (r *passengerRepository) GetPassengerByID(id int64) (*models.Passenger, error) {
	passenger := &models.Passenger{}

	query := `
		SELECT
			passengers.id,
			passengers.passenger_name_number,
			passengers.first_name,
			passengers.last_name,
			passengers.date_of_birth,
			passengers.gender,
			passengers.type,
			passengers.emails,
			passengers.phones,
			passengers.meal_preference,
			passengers.seat_preference,
			passengers.special_requests,
			passengers.special_service_request_remarks,
			passengers.issuing_country,
			passengers.country_of_birth,
			passengers.document_type,
			passengers.nationality,
			addresses.street_1,
			addresses.street_2,
			addresses.postcode,
			addresses.state,
			addresses.city,
			addresses.country,
			addresses.address_type
		FROM
			passengers
		JOIN
			addresses ON addresses.passenger_id = passengers.id
		WHERE
			passengers.id = $1
	`

	err := r.DB.
		QueryRow(query, id).
		Scan(
			&passenger.ID,
			&passenger.PassengerNameNumber,
			&passenger.PassengerDetails.FirstName,
			&passenger.PassengerDetails.LastName,
			&passenger.PassengerInfo.DateOfBirth,
			&passenger.PassengerInfo.Gender,
			&passenger.PassengerInfo.Type,
			&passenger.PassengerInfo.Emails,
			&passenger.PassengerInfo.Phones,
			&passenger.Preferences.MealPreference,
			&passenger.Preferences.SeatPreference,
			&passenger.Preferences.SpecialRequests,
			&passenger.Preferences.SpecialServiceRequestRemarks,
			&passenger.DocumentInfo.IssuingCountry,
			&passenger.DocumentInfo.CountryOfBirth,
			&passenger.DocumentInfo.DocumentType,
			&passenger.DocumentInfo.Nationality,
			&passenger.Address.Street1,
			&passenger.Address.Street2,
			&passenger.Address.Postcode,
			&passenger.Address.State,
			&passenger.Address.City,
			&passenger.Address.Country,
			&passenger.Address.AddressType,
		)

	if err != nil {
		return nil, err
	}

	return passenger, nil
}
