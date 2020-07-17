-- +migrate Up
CREATE TABLE IF NOT EXISTS "term_conditions"
(
    "id"          char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "term_name"   varchar(50)          NOT NULL,
    "term_type"   varchar(30)          NOT NULL,
    "description" text                 NOT NULL,
    "created_at"  timestamp            NOT NULL,
    "updated_at"  timestamp            NOT NULL,
    "deleted_at"  timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "term_conditions";