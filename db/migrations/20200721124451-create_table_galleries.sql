-- +migrate Up
CREATE TABLE IF NOT EXISTS "galleries"
(
    "id"                      char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "web_content_category_id" char(36)             NOT NULL,
    "gallery_name"            varchar(30)          NOT NULL,
    "created_at"              timestamp            NOT NULL,
    "updated_at"              timestamp            NOT NULL,
    "deleted_at"              timestamp            NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS "galleries";