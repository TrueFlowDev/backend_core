-- +goose Up
ALTER TABLE roles
    ADD COLUMN IF NOT EXISTS is_owner BOOLEAN NOT NULL DEFAULT false;

CREATE UNIQUE INDEX IF NOT EXISTS idx_roles_org_title ON roles (organization_id, title) WHERE deleted_at IS NULL;

-- +goose Down
DROP INDEX IF EXISTS idx_roles_org_title;

ALTER TABLE roles
    DROP COLUMN IF EXISTS is_owner;
