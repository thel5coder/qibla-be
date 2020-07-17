-- +migrate Up
CREATE TABLE iF NOT EXISTS "menus"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "menu_id"    char(3)              NOT NULL,
    "name"       varchar(20)          NOT NULL,
    "url"        text                 NOT NULL,
    "parent_id"  char(36),
    "is_active"  boolean,
    "created_at" timestamp            NOT NULL,
    "updated_at" timestamp            NOT NULL,
    "deleted_at" timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "menus";