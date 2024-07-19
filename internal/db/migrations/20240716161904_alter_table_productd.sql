-- +goose Up
-- +goose StatementBegin
ALTER TABLE products
    ADD COLUMN name VARCHAR(255) NOT NULL,
    ADD COLUMN description VARCHAR(255),
    ADD COLUMN price numeric(20, 2) NOT NULL,
    ADD COLUMN brand VARCHAR(255),
    ADD COLUMN visible BOOLEAN,
    ADD COLUMN amount numeric(20, 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
