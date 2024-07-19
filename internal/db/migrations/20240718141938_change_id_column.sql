-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
ALTER TABLE products
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
