-- +migrate Up
CREATE TABLE IF NOT EXISTS "crm_boards"
(
    "id"                 char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "crm_story_id"       char(36)             NOT NULL,
    "contact_id"         char(36)             NOT NULL,
    "opportunity"        varchar(255),
    "profit_expectation" float4,
    "created_at"         timestamp            NOT NULL,
    "updated_at"         timestamp            NOT NULL,
    "deleted_at"         timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "crm_boards";