-- +migrate Up
CREATE TABLE IF NOT EXISTS "faq_lists"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "faq_id"     char(36)             NOT NULL,
    "question"   varchar(255)         NOT NULL,
    "answer"     text                 NOT NULL,
    "created_at" timestamp            NOT NULL,
    "updated_at" timestamp            NOT NULL,
    "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "faq_lists";