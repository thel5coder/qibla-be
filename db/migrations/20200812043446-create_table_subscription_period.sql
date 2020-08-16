-- +migrate Up
CREATE TABLE IF NOT EXISTS "subscription_period"
(
    "id"                 char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "setting_product_id" char(36)             NOT NULL,
    "period"             int4                 NOT NULL,
    "created_at"         timestamp            NOT NULL,
    "updated_at"         timestamp            NOT NULL,
    "deleted_at"         timestamp
);

-- +migrate Down
DROP TABLE if exists setting_product_periods;