-- +migrate Up
CREATE TABLE IF NOT EXISTS "promotion_packages"
(
    "id"           char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "slug"         text,
    "package_name" varchar(30)          NOT NULL,
    "is_active"    boolean                       DEFAULT true,
    "created_at"   timestamp            NOT NULL,
    "updated_at"   timestamp            NOT NULL,
    "deleted_at"   timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS master_promotions;