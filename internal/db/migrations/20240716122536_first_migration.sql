-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products   (
    id uuid PRIMARY KEY NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
