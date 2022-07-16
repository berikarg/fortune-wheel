-- +goose Up
-- +goose StatementBegin
CREATE TABLE spin_results (
    id SERIAL PRIMARY KEY,
    result VARCHAR(255),
    time TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE spin_results;
-- +goose StatementEnd
