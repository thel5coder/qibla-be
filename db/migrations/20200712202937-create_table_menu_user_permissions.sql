-- +migrate Up
CREATE TABLE IF NOT EXISTS "menu_user_permissions"
(
    "id"          char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "user_id"     char(36)             NOT NULL,
    "menu_id"     char(36)             NOT NULL,
    "menu_scopes" varchar(100)         NOT NULL,
    "created_at"  timestamp            NOT NULL,
    "updated_at"  timestamp            NOT NULL,
    "deleted_at"  timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "menu_user_permissions";