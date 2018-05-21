CREATE TABLE notes (
    id BIGSERIAL,
    name varchar NOT NULL,
    note text NOT NULL,
    filename varchar NOT NULL,
    created_at timestamp without time zone NOT NULL default timezone('utc', now()),
	updated_at timestamp without time zone NOT NULL default timezone('utc', now()),
    PRIMARY KEY(id)
);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = timezone('utc', now());
	RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER set_notes_updated_at BEFORE UPDATE ON notes FOR EACH ROW EXECUTE PROCEDURE  set_updated_at();