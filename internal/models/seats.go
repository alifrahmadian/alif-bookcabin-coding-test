package models

type Seat struct {
	ID                     int64    `json:"id"`
	SeatRowID              int64    `json:"seat_row_id"`
	SeatRow                SeatRow  `json:"seat_row"`
	SegmentID              int64    `json:"segment_id"`
	Segment                Segment  `json:"segment"`
	SlotCharacteristics    []string `json:"slot_characteristic"`
	StorefrontSlotCode     string   `json:"storefront_slot_code"`
	Available              bool     `json:"available"`
	Entitled               bool     `json:"entitled"`
	FeeWaived              bool     `json:"fee_waived"`
	FreeOfCharge           bool     `json:"free_of_charge"`
	OriginallySelected     bool     `json:"originally_selected"`
	Code                   string   `json:"code"`
	Designations           []string `json:"designations"`
	EntitledRuleID         int64    `json:"entitled_rule_id"`
	FeeWaiveRuleID         int64    `json:"fee_waive_rule_id"`
	SeatCharacteristics    []string `json:"seat_characteristics"`
	Limitations            []string `json:"limintations"`
	RefundIndicator        string   `json:"refund_indicator"`
	RawSeatCharacteristics []string `json:"raw_seat_characteristics"`
	PriceAmount            float64  `json:"price_amount"`
	PriceCurrency          string   `json:"price_currency"`
	TaxAmount              float64  `json:"tax_amount"`
	TaxCurrency            string   `json:"tax_currency"`
	TotalAmount            float64  `json:"total_amount"`
	TotalCurrency          string   `json:"total_currency"`
}
