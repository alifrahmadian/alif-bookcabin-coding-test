package dtos

type Seat struct {
	SlotCharacteristics    []string `json:"slotCharacteristics,omitempty"`
	StorefrontSlotCode     string   `json:"storefrontSlotCode,omitempty"`
	Available              bool     `json:"available"`
	Entitled               bool     `json:"entitled"`
	FeeWaived              bool     `json:"feeWaived"`
	FreeOfCharge           bool     `json:"freeOfCharge"`
	OriginallySelected     bool     `json:"originallySelected"`
	Code                   string   `json:"code,omitempty"`
	Designations           []string `json:"designations,omitempty"`
	EntitledRuleId         string   `json:"entitledRuleId,omitempty"`
	FeeWaivedRuleId        string   `json:"feeWaivedRuleId,omitempty"`
	SeatCharacteristics    []string `json:"seatCharacteristics,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
	RefundIndicator        string   `json:"refundIndicator,omitempty"`
	Prices                 *Prices  `json:"prices,omitempty"`
	Taxes                  *Taxes   `json:"taxes,omitempty"`
	Total                  *Total   `json:"total,omitempty"`
	RawSeatCharacteristics []string `json:"rawSeatCharacteristics,omitempty"`
}

type Prices struct {
	Alternatives [][]Alternative `json:"alternatives,omitempty"`
}

type Taxes struct {
	Alternatives [][]Alternative `json:"alternatives,omitempty"`
}

type Total struct {
	Alternatives [][]Alternative `json:"alternatives,omitempty"`
}

type Alternative struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
