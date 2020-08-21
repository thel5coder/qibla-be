-- +migrate Up
CREATE TABLE IF NOT EXISTS "video_contents"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "channel"    varchar(100)         NOT NULL,
    "links"      varchar(255)         NOT NULL,
    "created_at" timestamp            NOT NULL,
    "updated_at" timestamp            NOT NULL,
    "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "video_contents";