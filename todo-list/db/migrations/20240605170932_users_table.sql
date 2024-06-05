-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL NOT NULL,
    name TEXT NOT NULL,
    date_created TIMESTAMP NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
