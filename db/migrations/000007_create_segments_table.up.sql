CREATE TABLE IF NOT EXISTS segments (
    id BIGSERIAL PRIMARY KEY,
    type TEXT,
    flight_miles BIGINT,
    award_flight BOOLEAN,
    duration INTEGER,
    cabin_class TEXT,
    equipment TEXT,
    flight_id BIGINT,
    origin TEXT,
    destination TEXT,
    departure DATE,
    arrival DATE,
    booking_class TEXT,
    layover_duration INTEGER,
    fare_basis TEXT,
    subject_to_government_approval BOOLEAN,
    segment_ref TEXT,

    FOREIGN KEY (flight_id) REFERENCES flights(id)
        ON DELETE CASCADE ON UPDATE CASCADE
)