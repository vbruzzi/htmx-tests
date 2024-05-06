-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos (
    id SERIAL NOT NULL,
    todo TEXT NOT NULL,
    created_on TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos;
-- +goose StatementEnd
