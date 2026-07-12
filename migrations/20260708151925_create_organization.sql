-- +goose Up
CREATE TYPE organization_category AS ENUM (
    'technology',
    'finance',
    'retail',
    'manufacturing',
    'other'
    );

CREATE TABLE IF NOT EXISTS organizations
(
    id         UUID PRIMARY KEY,

    category   organization_category NOT NULL,
    name       VARCHAR(100)          NOT NULL,
    active     bool                  NOT NULL,

    created_at TIMESTAMPTZ           NOT NULL,
    updated_at TIMESTAMPTZ           NOT NULL,
    deleted_at TIMESTAMPTZ
);


-- +goose Down
DROP TABLE IF EXISTS organizations;
DROP TYPE IF EXISTS organization_category;
