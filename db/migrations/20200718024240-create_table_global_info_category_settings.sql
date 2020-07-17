-- +migrate Up
CREATE TABLE IF NOT EXISTS "global_info_category_settings"
(
    "id"                      char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "global_info_category_id" char(36)             NOT NULL,
    "description"             text                 NOT NULL,
    "is_active"               boolean,
    "created_at"              timestamp            NOT NULL,
    "updated_at"              timestamp            NOT NULL,
    "deleted_at"              timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "global_info_category_settings";