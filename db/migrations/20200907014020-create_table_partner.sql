-- +migrate Up
CREATE TABLE IF NOT EXISTS "partners"
(
    "id"                      char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "user_id"                 char(36)             NOT NULL,
    "contact_id"              char(36)             NOT NULL,
    "contract_number"         varchar(20),
    "webinar_status"          boolean,
    "website_status"          boolean,
    "domain_site"             varchar(30),
    "domain_erp"              varchar(30),
    "database"                varchar(20),
    "invoice_publish_date"    date,
    "due_date"                date,
    "due_date_aging"          int,
    "verified_status"         partner_status_enum           DEFAULT 'unverified',
    "is_active"               boolean,
    "reason"                  varchar(255),
    "product_id"              char(36)             NOT NULL,
    "subscription_period"     int2,
    "subscription_expired_at" timestamp,
    "is_subscription_expired" boolean,
    "created_at"              timestamp            NOT NULL,
    "updated_at"              timestamp            NOT NULL,
    "deleted_at"              timestamp,
    "verified_at"             timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "partners";