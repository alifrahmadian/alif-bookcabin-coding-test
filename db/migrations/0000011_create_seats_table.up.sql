CREATE TABLE IF NOT EXISTS seats(
    id BIGSERIAL PRIMARY KEY,
    seat_row_id BIGINT,
    segment_id BIGINT,
    slot_characteristics TEXT[],
    storefront_slot_code TEXT,
    available BOOLEAN,
    entitled BOOLEAN,
    fee_waived BOOLEAN,
    free_of_charge BOOLEAN,
    originally_selected BOOLEAN,
    code TEXT,
    designations TEXT[],
    entitled_rule_id BIGINT,
    fee_waive_rule_id BIGINT,
    seat_characteristics TEXT[],
    limitations TEXT[],
    refund_indicator TEXT,
    raw_seat_characteristics TEXT[],
    price_amount REAL,
    price_currency TEXT,
    tax_amount REAL,
    tax_currency TEXT,
    total_amount REAL,
    total_currency TEXT,

    FOREIGN KEY (seat_row_id) REFERENCES seat_rows(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (segment_id) REFERENCES segments(id)
        ON DELETE CASCADE ON UPDATE CASCADE
);