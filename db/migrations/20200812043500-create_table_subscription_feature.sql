-- +migrate Up
CREATE TABLE "subscription_features"
(
    "id"                 char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "setting_product_id" char(36)             NOT NULL,
    "feature_name"       varchar(30)          NOT NULL,
    "created_at"         timestamp            NOT NULL,
    "updated_at"         timestamp            NOT NULL,
    "deleted_at"         timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS setting_product_features;
