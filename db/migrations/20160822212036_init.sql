
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
	id serial NOT NULL PRIMARY KEY,
	name text,
	age int
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE users;