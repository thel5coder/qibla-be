-- +migrate Up
CREATE TABLE IF NOT EXISTS "disbursements"
(
    "id"                char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "transaction_id"    char(36)             NOT NULL,
    "contact_id"        char(36)             NOT NULL,
    "total"             float4,
    "status"            status_transaction_enum,
    "disbursement_type" disbursement_type_enum,
    "start_period"      timestamp,
    "end_period"        timestamp,
    "account_number"    varchar(20)          NOT NULL,
    "account_name"      varchar(50)          NOT NULL,
    "account_bank_name" varchar(20)          NOT NULL,
    "account_bank_code" varchar(5)           NOT NULL,
    "payment_details"   jsonb,
    "created_at"        timestamp            NOT NULL,
    "updated_at"        timestamp            NOT NULL,
    "disburse_at"       timestamp,
    "deleted_at"        timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "disbursements";