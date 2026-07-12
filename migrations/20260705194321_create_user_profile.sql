-- +goose Up
CREATE TABLE IF NOT EXISTS users_profile
(
    user_id UUID PRIMARY KEY,
    email TEXT NOT NULL,

    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    headline VARCHAR(100) NOT NULL DEFAULT '',
    bio TEXT NOT NULL DEFAULT '',

    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_users_profile_user
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE UNIQUE INDEX uq_users_profile_email ON users_profile(email) WHERE deleted_at IS NULL;


-- +goose Down
DROP INDEX IF EXISTS uq_users_profile_email;
DROP TABLE IF EXISTS users_profile;
