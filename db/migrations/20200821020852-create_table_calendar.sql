-- +migrate Up
CREATE TABLE IF NOT EXISTS "calendars"
(
    "id"          char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "title"       varchar(50)          NOT NULL,
    "start"  timestamp            NOT NULL,
    "end"    timestamp,
    "description" text,
    "created_at"  timestamp            NOT NULL,
    "updated_at"  timestamp            NOT NULL,
    "deleted_at"  timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "calendars";