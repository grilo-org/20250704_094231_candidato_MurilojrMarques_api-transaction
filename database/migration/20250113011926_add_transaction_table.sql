-- +goose Up
CREATE TABLE IF NOT EXISTS transaction(
    ID UUID NOT NULL PRIMARY KEY,
    Description VARCHAR(50) NOT NULL,
    Date DATE NOT NULL,
    value NUMERIC(10, 2) CHECK (VALUE > 0)
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transaction;
-- +goose StatementEnd
