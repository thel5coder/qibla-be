-- +migrate Up
CREATE TABLE IF NOT EXISTS "promotion_positions"
(
    "id"                    char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "promotion_platform_id" char(36)             NOT NULL,
    "position"              varchar(30)
);
-- +migrate Down
DROP TABLE IF EXISTS "promotion_positions";