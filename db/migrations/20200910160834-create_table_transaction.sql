-- +migrate Up
CREATE TABLE IF NOT EXISTS "transactions"
(
    "id"                  char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "user_id"             char(36)             NOT NULL,
    "invoice_number"      varchar(20),
    "trx_id"              varchar(30),
    "due_date"            date,
    "due_date_period"     int,
    "payment_status"      payment_status_enum,
    "payment_method_code" char(3),
    "va_number"           varchar(20),
    "bank_name"           varchar(10),
    "direction"           transaction_direction_enum,
    "transaction_type"    transaction_type_enum,
    "paid_date"           date,
    "transaction_date"    timestamp,
    "updated_at"          timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "transactions";
