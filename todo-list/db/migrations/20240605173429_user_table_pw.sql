-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN password TEXT NOT NULL DEFAULT 'pw';
ALTER TABLE users RENAME name TO username;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN password;
ALTER TABLE users RENAME username TO name;
-- +goose StatementEnd
