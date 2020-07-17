
-- +migrate Up
CREATE TABLE IF NOT EXISTS "global_info_categories" (
                                          "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
                                          "name" varchar(30) NOT NULL,
                                          "created_at" timestamp NOT NULL,
                                          "updated_at" timestamp NOT NULL,
                                          "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "global_info_categories";