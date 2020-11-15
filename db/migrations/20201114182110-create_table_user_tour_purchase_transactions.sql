-- +migrate Up
CREATE TABLE IF NOT EXISTS "user_tour_purchase_transactions"
(
    "id"                    char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "transaction_id"        char(36)             NOT NULL,
    "user_tour_purchase_id" char(36)             NOT NULL,
    "created_at"            timestamp            NOT NULL,
    "updated_at"            timestamp            NOT NULL,
    "deleted_at"            timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "user_tour_purchase_transaction";