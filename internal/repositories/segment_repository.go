package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/models"
	"github.com/lib/pq"
)

type SegmentRepository interface {
	GetSegmentAndFlightBySegmentID(id int64) (*models.Segment, error)
}

type segmentRepository struct {
	DB *sql.DB
}

func NewSegmentRepository(db *sql.DB) SegmentRepository {
	return &segmentRepository{
		DB: db,
	}
}

func (r *segmentRepository) GetSegmentAndFlightBySegmentID(id int64) (*models.Segment, error) {
	query := `
		SELECT
			segments.id,
			segments.type,
			segments.flight_miles,
			segments.award_flight,
			segments.duration,
			segments.cabin_class,
			segments.equipment,
			segments.flight_id,
			flights.flight_number,
			flights.operating_flight_number,
			flights.airline_code,
			flights.operating_airline_code,
			flights.stop_airports,
			flights.departure_terminal,
			flights.arrival_terminal,
			segments.origin,
			segments.destination,
			segments.departure,
			segments.arrival,
			segments.booking_class,
			segments.layover_duration,
			segments.fare_basis,
			segments.subject_to_government_approval,
			segments.segment_ref
		FROM
			segments
		JOIN
			flights ON flights.id = segments.flight_id
		WHERE
			segments.id = $1
	`

	segment := &models.Segment{}

	err := r.DB.
		QueryRow(query, id).
		Scan(
			&segment.ID,
			&segment.Type,
			&segment.SegmentOfferInformation.FlightMiles,
			&segment.SegmentOfferInformation.AwardFlight,
			&segment.Duration,
			&segment.CabinClass,
			&segment.Equipment,
			&segment.FlightID,
			&segment.Flight.FlightNumber,
			&segment.Flight.OperatingFlightNumber,
			&segment.Flight.AirlineCode,
			&segment.Flight.OperatingAirlineCode,
			pq.Array(&segment.Flight.StopAirports),
			&segment.Flight.DepartureTerminal,
			&segment.Flight.ArrivalTerminal,
			&segment.Origin,
			&segment.Destination,
			&segment.Departure,
			&segment.Arrival,
			&segment.BookingClass,
			&segment.LayoverDuration,
			&segment.FareBasis,
			&segment.SubjectToGovernmentApproval,
			&segment.SegmentRef,
		)

	if err != nil {
		return nil, err
	}

	return segment, nil
}
