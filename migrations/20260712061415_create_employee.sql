-- +goose Up
CREATE TYPE membership_status AS ENUM (
    'active',
    'on_leave',
    'suspended',
    'resigned',
    'terminated'
    );

CREATE TYPE employment_type AS ENUM (
    'full_time',
    'part_time',
    'contract',
    'intern',
    'temporary',
    'consultant'
    );

CREATE TABLE IF NOT EXISTS employees
(
    id                UUID PRIMARY KEY,
    user_id           UUID              NOT NULL,
    organization_id   UUID              NOT NULL,
    role_id           UUID              NOT NULL,


    job_title         VARCHAR(100)      NOT NULL,
    membership_status membership_status NOT NULL,
    employment_type   employment_type   NOT NULL,

    created_at        TIMESTAMPTZ       NOT NULL,
    updated_at        TIMESTAMPTZ       NOT NULL,
    deleted_at        TIMESTAMPTZ,

    CONSTRAINT fk_employees_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,

    CONSTRAINT fk_employees_organization
        FOREIGN KEY (organization_id)
            REFERENCES organizations (id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,

    CONSTRAINT fk_employees_role
        FOREIGN KEY (role_id, organization_id)
            REFERENCES roles (id, organization_id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_employees_organization_id_user_id
    ON employees (organization_id, user_id) WHERE deleted_at IS NULL;

-- +goose Down
DROP TABLE IF EXISTS employees;

DROP TYPE IF EXISTS employment_type;
DROP TYPE IF EXISTS membership_status;
