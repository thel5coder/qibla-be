-- +migrate Up
CREATE TABLE IF NOT EXISTS "calendar_participants"
(
    "id"          char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "calendar_id" char(36)             NOT NULL,
    "email"       varchar(50)          NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS "calendar_participants";