-- +migrate Up
CREATE TABLE IF NOT EXISTS "user_scopes"
(
    "id"                 char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "user_id"            char(36)             NOT NULL,
    "menu_permission_id" char(36)             NOT NULL,
    "created_at"         timestamp            NOT NULL,
    "updated_at"         timestamp            NOT NULL,
    "deleted_at"         timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "user_scopes";