package dtos

type Seat struct {
	SlotCharacteristics    []string `json:"slotCharacteristics"`
	StorefrontSlotCode     []string `json:"storefrontSlotCode"`
	Available              bool     `json:"available"`
	Entitled               bool     `json:"entitled"`
	FeeWaived              bool     `json:"feeWaived"`
	FreeOfCharge           bool     `json:"freeOfCharge"`
	OriginallySelected     bool     `json:"originallySelected"`
	Code                   string   `json:"code"`
	Designations           []string `json:"designations"`
	EntitledRuleId         int64    `json:"entitledRuleId"`
	FeeWaivedRuleId        int64    `json:"feeWaivedRuleId"`
	SeatCharacteristics    []string `json:"seatCharacteristics"`
	Limitations            []string `json:"limitations"`
	RefundIndicator        string   `json:"refundIndicator"`
	Prices                 Prices   `json:"prices"`
	Taxes                  Taxes    `json:"taxes"`
	Total                  Total    `json:"total"`
	RawSeatCharacteristics []string `json:"rawSeatCharacteristics"`
}

type Prices struct {
	Alternatives [][]Alternative `json:"alternatives"`
}

type Taxes struct {
	Alternatives [][]Alternative `json:"alternatives"`
}

type Total struct {
	Alternatives [][]Alternative `json:"alternatives"`
}

type Alternative struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
