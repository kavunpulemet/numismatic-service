-- +goose Up
-- +goose StatementBegin
CREATE TABLE coins (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    country TEXT NOT NULL,
    year INT NOT NULL,
    denomination TEXT NOT NULL,
    material TEXT NOT NULL,
    weight NUMERIC(10, 2),
    diameter NUMERIC(10, 2),
    thickness NUMERIC(10, 2),
    condition TEXT,
    mintmark TEXT,
    historicalinfo TEXT,
    value NUMERIC(12, 2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  coins;
-- +goose StatementEnd
