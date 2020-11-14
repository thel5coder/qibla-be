-- +migrate Up
CREATE TABLE IF NOT EXISTS "user_tour_purchases"
(
    "id"                     char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "tour_package_id"        char(36)             NOT NULL,
    "customer_name"          varchar(50)          NOT NULL,
    "customer_identity_type" char(3),
    "identity_number"        varchar(20),
    "full_name"              varchar(50),
    "sex"                    sex_enum,
    "birth_date"             date,
    "birth_place"            varchar(100),
    "phone_number"           varchar(20),
    "city_id"                char(36),
    "marital_status"         marital_status_enum,
    "customer_address"       text,
    "user_id"                char(36)             NOT NULL,
    "contact_id"             char(36)             NOT NULL,
    "airline_odoo_id"        int2,
    "airlines_class"         varchar(50),
    "airline_price"          float4,
    "created_at"             timestamp            NOT NULL,
    "updated_at"             timestamp            NOT NULL,
    "deleted_at"             timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "user_tour_purchases";