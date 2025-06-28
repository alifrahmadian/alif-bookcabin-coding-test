CREATE TABLE IF NOT EXISTS passengers(
    id BIGSERIAL PRIMARY KEY,
    passenger_name_number VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    date_of_birth DATE,
    gender TEXT,
    type TEXT,
    emails TEXT[],
    phones TEXT[],
    meal_preference TEXT,
    seat_preference TEXT,
    special_requests TEXT[],
    special_service_request_remarks TEXT[],
    issuing_country TEXT,
    country_of_birth TEXT,
    document_type TEXT,
    nationality TEXT
);  