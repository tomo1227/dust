-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
    id int NOT NULL,
    name text,
    age int,
    PRIMARY KEY(id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE user;

-- +goose StatementEnd