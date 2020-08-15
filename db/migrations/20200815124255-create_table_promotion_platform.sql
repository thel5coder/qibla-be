-- +migrate Up
CREATE TABLE IF NOT EXISTS "promotion_platforms"
(
    "id"           char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "promotion_id" char(36)             NOT NULL,
    "platform"     varchar(30)
);

-- +migrate Down
DROP TABLE IF EXISTS "promotion_platforms";