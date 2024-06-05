-- +goose Up
-- +goose StatementBegin
ALTER TABLE todos ADD COLUMN user_id INT NOT NULL;
ALTER TABLE todos ADD FOREIGN KEY (user_id) REFERENCES users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todos DROP COLUMN user_id;
-- +goose StatementEnd
