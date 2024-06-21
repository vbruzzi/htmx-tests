-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN login_key_salt TEXT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN login_key_salt;
-- +goose StatementEnd
