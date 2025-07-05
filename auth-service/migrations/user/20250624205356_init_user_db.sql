-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    password_salt VARCHAR(10),
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    short_name VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS roles(
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    permissions JSONB NOT NULL
);

CREATE TABLE IF NOT EXISTS user_roles(
    user_id UUID NOT NULL,
    role_id UUID NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_roles;
DROP TABLE roles;
DROP TABLE users;
-- +goose StatementEnd
