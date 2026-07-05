-- +goose Up
CREATE TABLE IF NOT EXISTS users_profile
(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    email TEXT NOT NULL,

    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    headline VARCHAR(100) NOT NULL DEFAULT '',
    bio TEXT NOT NULL DEFAULT '',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,

    CONSTRAINT fk_users_profile_user
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE UNIQUE INDEX uq_users_profile_user_id ON users_profile(user_id);
CREATE UNIQUE INDEX uq_users_profile_email ON users_profile(email);


-- +goose Down
DROP INDEX IF EXISTS uq_users_profile_user_id;
DROP INDEX IF EXISTS uq_users_profile_email;
DROP TABLE IF EXISTS users_profile;
