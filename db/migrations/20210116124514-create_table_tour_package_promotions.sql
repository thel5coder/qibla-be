-- +migrate Up
CREATE TABLE IF NOT EXISTS "tour_package_promotions"
(
    "id"              char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "tour_package_id" char(36)             NOT NULL,
    "created_at"      timestamp            NOT NULL,
    "updated_at"      timestamp            NOT NULL,
    "deleted_at"      timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "tour_package_promotions";