-- +migrate Up
CREATE TABLE IF NOT EXISTS "tour_package_airlines"
(
    "id"                       char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name"                     varchar(50)          NOT NULL,
    "odoo_product_template_id" int4,
    "created_at"               timestamp            NOT NULL,
    "updated_at"               timestamp            NOT NULL,
    "deleted_at"               timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "tour_package_airlines";