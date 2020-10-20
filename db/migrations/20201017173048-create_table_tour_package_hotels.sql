-- +migrate Up
CREATE TABLE "tour_package_hotels"
(
    "id"              char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "Name"            varchar(50)          NOT NULL,
    "facility_rating" int2,
    "location"        varchar(50),
    "created_at"      timestamp            NOT NULL,
    "updated_at"      timestamp            NOT NULL,
    "deleted_at"      timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "tour_package_hotels";