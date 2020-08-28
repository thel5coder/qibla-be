-- +migrate Up
CREATE TABLE IF NOT EXISTS "satisfaction_categories"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "parent_id"  char(36),
    "slug"       varchar(255)         NOT NULL,
    "name"       varchar(100)         NOT NULL,
    "is_active"  bool,
    "created_at" timestamp            NOT NULL,
    "updated_at" timestamp            NOT NULL,
    "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE if exists "satisfaction_categories";