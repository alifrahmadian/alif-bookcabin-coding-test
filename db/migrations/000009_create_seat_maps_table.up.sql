CREATE TABLE IF NOT EXISTS seat_maps(
    ID BIGSERIAL PRIMARY KEY,
    segment_id BIGINT,
    rows_disabled_causes TEXT[],
    aircraft TEXT,

    FOREIGN KEY (segment_id) REFERENCES segments(id)
        ON DELETE CASCADE ON UPDATE CASCADE
);