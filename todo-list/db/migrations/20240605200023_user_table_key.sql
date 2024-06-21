-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN login_key TEXT NULL UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN login_key;
-- +goose StatementEnd
