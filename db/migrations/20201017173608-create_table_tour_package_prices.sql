-- +migrate Up
CREATE TABLE IF NOT EXISTS "tour_package_prices"
(
    "id"            char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "room_type"     varchar(30)          NOT NULL,
    "room_capacity" int,
    "price"         float4,
    "promo_price"   float4,
    "airline_class" varchar(30),
    "created_at"    timestamp            NOT NULL,
    "updated_at"    timestamp            NOT NULL,
    "deleted_at"    timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "tour_package_prices";