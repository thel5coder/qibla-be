-- +migrate Up
CREATE TABLE IF NOT EXISTS "users"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "username"   varchar(50)          NOT NULL,
    "email"      varchar(50)          NOT NULL,
    "password"   varchar(128)         NOT NULL,
    "is_active"  boolean,
    "role_id"    char(36)             NOT NULL,
    "created_at" timestamp,
    "updated_at" timestamp            NOT NULL,
    "deleted_at" timestamp
);


-- +migrate Down
