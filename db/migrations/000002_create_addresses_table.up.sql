CREATE TABLE IF NOT EXISTS addresses (
    ID BIGSERIAL PRIMARY KEY,
    passenger_id BIGINT,
    street_1 VARCHAR(255),
    street_2 VARCHAR(255),
    postcode TEXT,
    state TEXT,
    city TEXT,
    country TEXT,
    address_type TEXT,

    FOREIGN KEY (passenger_id) REFERENCES passengers(id)
        ON DELETE CASCADE ON UPDATE CASCADE
);