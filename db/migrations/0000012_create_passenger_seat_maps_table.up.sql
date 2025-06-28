CREATE TABLE IF NOT EXISTS passenger_seat_maps(
    ID BIGSERIAL PRIMARY KEY,
    seat_selection_enabled_for_pax BOOLEAN,
    passenger_id BIGINT,
    segment_seat_map_id BIGINT,
    seat_map_id BIGINT,

    FOREIGN KEY (passenger_id) REFERENCES passengers(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (segment_seat_map_id) REFERENCES segment_seat_maps(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (seat_map_id) REFERENCES seat_maps(id)
        ON DELETE CASCADE ON UPDATE CASCADE
);