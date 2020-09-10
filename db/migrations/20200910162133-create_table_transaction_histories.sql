-- +migrate Up

CREATE TABLE IF NOT EXISTS "transaction_histories"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "trx_id"     varchar(50),
    "status"     varchar(30),
    "response"   text,
    "created_at" timestamp,
    "updated_at" timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "transaction_histories";