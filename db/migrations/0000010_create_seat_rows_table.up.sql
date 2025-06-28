CREATE TABLE IF NOT EXISTS seat_rows(
    id BIGSERIAL PRIMARY KEY,
    cabin_id BIGINT,
    row_number INTEGER,
    seat_codes TEXT[],

    FOREIGN KEY (cabin_id) REFERENCES cabins(id)
        ON DELETE CASCADE ON UPDATE CASCADE
);