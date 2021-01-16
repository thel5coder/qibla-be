-- +migrate Up
CREATE TABLE IF NOT EXISTS "tour_packages"
(
    "id"                         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "partner_id"                 char(36)             NOT NULL,
    "odoo_package_id"            int2                 NOT NULL,
    "name"                       varchar(100)         NOT NULL,
    "package_type"               varchar(100),
    "program_day"                int4,
    "departure_airport"          varchar(100),
    "destination_airport"        varchar(100),
    "return_departure_airport"   varchar(100),
    "return_destination_airport" varchar(100),
    "departure_date"             date,
    "return_date"                date,
    "description"                text,
    "notes"                      text,
    "quota"                      int4,
    "image"                      text,
    "created_at"                 timestamp            NOT NULL,
    "updated_at"                 timestamp            NOT NULL,
    "deleted_at"                 timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "tour_packages";