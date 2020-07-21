-- +migrate Up
CREATE TABLE IF NOT EXISTS "gallery_images"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "gallery_id" char(36)             NOT NULL,
    "file_id"    char(36)             NOT NULL,
    "created_at" timestamp            NOT NULL,
    "updated_at" timestamp            NOT NULL,
    "deleted_at" timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "gallery_images";