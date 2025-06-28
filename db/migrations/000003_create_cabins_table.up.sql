CREATE TABLE IF NOT EXISTS cabins(
    id BIGSERIAL PRIMARY KEY,
    deck TEXT,
    seat_columns TEXT[],
    first_row INTEGER,
    last_row INTEGER,
    aircraft TEXT
);