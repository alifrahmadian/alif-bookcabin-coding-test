package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/models"
	"github.com/lib/pq"
)

type SeatRepository interface {
	GetSeatsBySeatRowIDAndSegmentID(seatRowID, segmentID int64) ([]*models.Seat, error)
}

type seatRepository struct {
	DB *sql.DB
}

func NewSeatRepository(db *sql.DB) SeatRepository {
	return &seatRepository{
		DB: db,
	}
}

func (r *seatRepository) GetSeatsBySeatRowIDAndSegmentID(seatRowID, segmentID int64) ([]*models.Seat, error) {
	var seats []*models.Seat

	query := `
		SELECT
			id,
			seat_row_id,
			segment_id,
			slot_characteristics,
			storefront_slot_code,
			available,
			entitled,
			fee_waived,
			free_of_charge,
			originally_selected,
			code,
			designations,
			entitled_rule_id,
			fee_waive_rule_id,
			seat_characteristics,
			limitations,
			refund_indicator,
			raw_seat_characteristics,
			price_amount,
			price_currency,
			tax_amount,
			tax_currency,
			total_amount,
			total_currency
		FROM
			seats
		WHERE
			seat_row_id = $1 AND segment_id = $2
	`

	rows, err := r.DB.Query(query, seatRowID, segmentID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		seat := &models.Seat{}
		err := rows.Scan(
			&seat.ID,
			&seat.SeatRowID,
			&seat.SegmentID,
			pq.Array(&seat.SlotCharacteristics),
			&seat.StorefrontSlotCode,
			&seat.Available,
			&seat.Entitled,
			&seat.FeeWaived,
			&seat.FreeOfCharge,
			&seat.OriginallySelected,
			&seat.Code,
			pq.Array(&seat.Designations),
			&seat.EntitledRuleID,
			&seat.FeeWaiveRuleID,
			pq.Array(&seat.SeatCharacteristics),
			pq.Array(&seat.Limitations),
			&seat.RefundIndicator,
			pq.Array(&seat.RawSeatCharacteristics),
			&seat.PriceAmount,
			&seat.PriceCurrency,
			&seat.TaxAmount,
			&seat.TaxCurrency,
			&seat.TotalAmount,
			&seat.TotalCurrency,
		)

		if err != nil {
			return nil, err
		}

		seats = append(seats, seat)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return seats, nil
}
