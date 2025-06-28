CREATE TABLE IF NOT EXISTS flights(
    id BIGSERIAL PRIMARY KEY,
    flight_number INTEGER,
    operating_flight_number INTEGER,
    airline_code TEXT,
    operating_airline_code TEXT,
    stop_airports TEXT[],
    departure_terminal TEXT,
    arrival_terminal TEXT
);