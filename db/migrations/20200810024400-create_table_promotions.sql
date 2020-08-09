-- +migrate Up
CREATE TABLE IF NOT EXISTS "promotions"
(
    "id"                   char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "promotion_package_id" char(36)             NOT NULL,
    "package_promotion"    package_promotion_enum        DEFAULT 'per_hari',
    "start_date"           date,
    "end_date"             date,
    "platform"             platform_enum                 DEFAULT 'semua',
    "position"             promotion_position_enum       DEFAULT 'slider_1',
    "price"                int4                 NOT NULL,
    "description"          text                 NOT NULL,
    "created_at"           timestamp            NOT NULL,
    "updated_at"           timestamp            NOT NULL,
    "deleted_at"           timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "promotions";