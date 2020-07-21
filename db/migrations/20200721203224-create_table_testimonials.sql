-- +migrate Up
CREATE TABLE IF NOT EXISTS "testimonials"
(
    "id"                      char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "web_content_category_id" char(36)             NOT NULL,
    "file_id"                 char(36)             NOT NULL,
    "customer_name"           varchar(50)          NOT NULL,
    "job_position"            varchar(30)          NOT NULL,
    "testimoni"               text                 NOT NULL,
    "rating"                  int,
    "is_active"               boolean                       DEFAULT true,
    "created_at"              timestamp            NOT NULL,
    "updated_at"              timestamp            NOT NULL,
    "deleted_at"              timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "testimonials";