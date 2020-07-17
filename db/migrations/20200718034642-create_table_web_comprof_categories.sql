-- +migrate Up
CREATE TABLE IF NOT EXISTS "web_comprof_categories"
(
    "id"            char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "slug"          text                 NOT NULL,
    "name"          varchar(30)          NOT NULL,
    "category_type" web_content_category_type_enum,
    "created_at"    timestamp,
    "updated_at"    timestamp,
    "deleted_at"    timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "web_comprof_categories";