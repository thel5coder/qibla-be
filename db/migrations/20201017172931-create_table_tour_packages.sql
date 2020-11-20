-- +migrate Up
CREATE TABLE IF NOT EXISTS "tour_packages"
(
    "id"                      char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "partner_id"              char(36)             NOT NULL,
    "odoo_package_id"         int2                 NOT NULL,
    "name"                    varchar(100)         NOT NULL,
    "odoo_package_program_id" int4,
    "package_program"         varchar(100),
    "program_day"             int4,
    "departure_date"          date,
    "return_date"             date,
    "description"             text,
    "notes"                   text,
    "created_at"              timestamp            NOT NULL,
    "updated_at"              timestamp            NOT NULL,
    "deleted_at"              timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "tour_packages";