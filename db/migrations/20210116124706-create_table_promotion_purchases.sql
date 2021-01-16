-- +migrate Up
CREATE TABLE IF NOT EXISTS "promotion_purchases"
(
    "id"                        char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "promotion_id"              char(36)             NOT NULL,
    "tour_package_promotion_id" char(36)             NOT NULL,
    "transaction_id"            char(36)             NOT NULL,
    "start_date"                date,
    "end_date"                  date,
    "price"                     float4,
    "status"                    tour_package_promotion_status_enum,
    "is_active"                 boolean,
    "created_at"                timestamp            NOT NULL,
    "updated_at"                timestamp            NOT NULL,
    "deleted_at"                timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "promotion_purchases";