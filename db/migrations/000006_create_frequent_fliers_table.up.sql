CREATE TABLE IF NOT EXISTS frequent_fliers(
    ID BIGSERIAL PRIMARY KEY,
    passenger_id BIGINT,
    airline TEXT,
    number TEXT,
    tier_number INTEGER,

    FOREIGN KEY (passenger_id) REFERENCES passengers(id)
        ON DELETE CASCADE ON UPDATE CASCADE
);