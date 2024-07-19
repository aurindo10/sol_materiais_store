-- +goose Up
-- +goose StatementBegin
ALTER TABLE products
    ADD COLUMN category VARCHAR(255) NOT NULL,
    ADD COLUMN weight numeric(20, 2);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
