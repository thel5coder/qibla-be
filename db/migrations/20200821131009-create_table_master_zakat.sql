-- +migrate Up
CREATE TABLE IF NOT EXISTS "master_zakats"
(
    "id"                 char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "Slug"               text,
    "type_zakat"         type_zakat_enum,
    "name"               varchar(50),
    "description"        text,
    "amount"             int4,
    "current_gold_price" int4,
    "gold_nishab"        int4,
    "created_at"         timestamp            NOT NULL,
    "updated_at"         timestamp            NOT NULL,
    "deleted_at"         timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "master_products";