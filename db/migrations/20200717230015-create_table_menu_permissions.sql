-- +migrate Up
CREATE TABLE IF NOT EXISTS "menu_permissions"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "menu_id"    char(36)             NOT NULL,
    "permission" menu_permission_enums,
    "created_at" timestamp            NOT NULL,
    "updated_at" timestamp            NOT NULL,
    "deleted_at" timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "menu_permissions";