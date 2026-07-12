-- +goose Up
CREATE TABLE IF NOT EXISTS roles
(
    id              UUID PRIMARY KEY,
    organization_id UUID         NOT NULL,

    title           VARCHAR(100) NOT NULL,

    created_at      TIMESTAMPTZ  NOT NULL,
    updated_at      TIMESTAMPTZ  NOT NULL,
    deleted_at      TIMESTAMPTZ,

    CONSTRAINT uq_roles_id_organization
        UNIQUE (id, organization_id),

    CONSTRAINT fk_roles_organization
        FOREIGN KEY (organization_id)
            REFERENCES organizations (id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS role_permissions
(
    role_id          UUID         NOT NULL REFERENCES roles (id) ON UPDATE CASCADE ON DELETE CASCADE,
    permission_value VARCHAR(100) NOT NULL,

    PRIMARY KEY (role_id, permission_value)
);

-- +goose Down
DROP TABLE IF EXISTS role_permissions;

DROP TABLE IF EXISTS roles;
