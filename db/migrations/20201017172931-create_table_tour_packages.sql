-- +migrate Up
CREATE TABLE IF NOT EXISTS "tour_packages"
(
    "id"              char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "odoo_package_id" int2                 NOT NULL,
    "Name"            varchar(100)         NOT NULL,
    "departure_date"  date,
    "return_date"     date,
    "description"     text,
    "created_at"      timestamp            NOT NULL,
    "updated_at"      timestamp            NOT NULL,
    "deleted_at"      timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "tour_packages";