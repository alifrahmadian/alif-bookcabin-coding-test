CREATE TABLE IF NOT EXISTS segment_seat_maps(
    id BIGSERIAL PRIMARY KEY,
    seats_itinerary_part_id BIGINT,
    segment_id BIGINT,

    FOREIGN KEY (seats_itinerary_part_id) REFERENCES seats_itinerary_parts(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (segment_id) REFERENCES segments(id)
        ON DELETE SET NULL ON UPDATE CASCADE
);