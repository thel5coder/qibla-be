-- +migrate Up
CREATE TABLE IF NOT EXISTS "partner_extra_products"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "partner_id" char(36)             NOT NULL,
    "product_id" char(36)             NOT NULL,
    "created_at" timestamp            NOT NULL,
    "updated_at" timestamp            NOT NULL,
    "deleted_at" timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS partner_extra_products;