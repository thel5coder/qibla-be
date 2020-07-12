-- +migrate Up
CREATE TABLE IF NOT EXISTS "roles"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name"       varchar(30)          NOT NULL,
    "created_at" timestamp            NOT NULL,
    "updated_at" timestamp            NOT NULL,
    "deleted_at" timestamp
);


-- +migrate Down
DROP TABLE IF EXISTS "roles";