-- +migrate Up
CREATE TABLE IF NOT EXISTS "setting_products"
(
    "id"                    char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "product_id"            char(36)             NOT NULL,
    "price"                 int4                 NOT NULL,
    "price_unit"            price_unit_enum,
    "maintenance_price"     int4,
    "discount"              int4,
    "discount_type"         discount_type_enum   NOT NULL DEFAULT 'fix',
    "discount_period_start" date,
    "discount_period_end"   date,
    "description"           text                 NOT NULL,
    "sessions"              varchar(30),
    "created_at"            timestamp            NOT NULL,
    "updated_at"            timestamp            NOT NULL,
    "deleted_at"            timestamp
);


-- +migrate Down
DROP TABLE IF EXISTS "setting_products";