-- +migrate Up
CREATE TABLE IF NOT EXISTS "app_complaints"
(
    "id"             char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "FullName"       varchar(50)          NOT NULL,
    "Email"          varchar(50)          NOT NULL,
    "ticket_number"  char(6)              NOT NULL,
    "complaint_type" varchar(30),
    "complaint"      text,
    "solution"       text,
    "is_open"        complaint_status_enum         DEFAULT 'open',
    "created_at"     timestamp            NOT NULL,
    "updated_at"     timestamp            NOT NULL,
    "deleted_at"     timestamp
);

-- +migrate Down

DROP TABLE IF EXISTS "app_complaints";
