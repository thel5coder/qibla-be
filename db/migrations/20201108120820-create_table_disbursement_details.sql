-- +migrate Up
CREATE TABLE IF NOT EXISTS "disbursement_details"
(
    "id"             char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "disbursement_id" char(36)             NOT NULL,
    "transaction_id" char(36)             NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS "disbursement_details";