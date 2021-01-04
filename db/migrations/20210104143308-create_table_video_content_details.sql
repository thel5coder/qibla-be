-- +migrate Up
CREATE TABLE IF NOT EXISTS "video_content_details"
(
    "id"               char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "youtube_video_id" varchar(20),
    "video_content_id" char(36)             NOT NULL,
    "title"            varchar(255)         not null,
    "channel_title"    varchar(255)         not null,
    "embedded_url"     varchar(255),
    "thumbnails"       json,
    "description"      text,
    "published_at"     varchar(100),
    "created_at"       timestamp            NOT NULL,
    "updated_at"       timestamp            NOT NULL,
    "deleted_at"       timestamp
);
-- +migrate Down
drop table if exists "video_content_details";