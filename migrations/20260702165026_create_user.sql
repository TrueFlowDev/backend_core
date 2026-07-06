-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    id         UUID PRIMARY KEY,

    phone      VARCHAR(20) NOT NULL,
    password   VARCHAR(255),

    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_users_phone ON users(phone) WHERE deleted_at IS NULL;

-- +goose Down
DROP INDEX IF EXISTS uq_users_phone;
DROP TABLE IF EXISTS users;
