--
-- create unique index
ALTER TABLE fieldvalues ADD UNIQUE INDEX unique_fieldvalues( field_name, field_value );