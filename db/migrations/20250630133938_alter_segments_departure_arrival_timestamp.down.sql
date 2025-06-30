ALTER TABLE segments
  ALTER COLUMN departure TYPE DATE USING departure::date,
  ALTER COLUMN arrival TYPE DATE USING arrival::date;
