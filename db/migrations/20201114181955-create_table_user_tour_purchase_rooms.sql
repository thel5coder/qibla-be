-- +migrate Up
CREATE TABLE IF NOT EXISTS "user_tour_purchase_rooms"
(
    "id"                    char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "user_tour_purchase_id" char(36)             NOT NULL,
    "tour_package_price_id" char(36)             NOT NULL,
    "price"                 float4,
    "quantity"              int2,
    "created_at"            timestamp            NOT NULL,
    "updated_at"            timestamp            NOT NULL,
    "deleted_at"            timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "user_tour_purchase_rooms";