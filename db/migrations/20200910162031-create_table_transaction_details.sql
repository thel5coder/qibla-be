-- +migrate Up
CREATE TABLE IF NOT EXISTS "transaction_details"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name"       varchar(255)         NOT NULL,
    "fee"        int4,
    "price"      int4,
    "created_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS"transaction_details";