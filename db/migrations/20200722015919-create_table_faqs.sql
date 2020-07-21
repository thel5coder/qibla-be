-- +migrate Up
CREATE TABLE IF NOT EXISTS "faqs"
(
    "id"                      char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "web_content_category_id" char(36)             NOT NULL,
    "faq_category_name"       varchar(50)          NOT NULL,
    "created_at"              timestamp            NOT NULL,
    "updated_at"              timestamp            NOT NULL,
    "deleted_at"              timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "faqs";