-- +migrate Up
CREATE TABLE IF NOT EXISTS "contacts"
(
    "id"                     char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "branch_name"            varchar(50)          NOT NULL,
    "travel_agent_name"      varchar(50)          NOT NULL,
    "address"                text                 NOT NULL,
    "longitude"              varchar(255),
    "latitude"               varchar(255),
    "area_code"              char(5)              NOT NULL,
    "phone_number"           char(20)             NOT NULL,
    "user_id"                char(36)             NOT NULL,
    "sk_number"              varchar(20)          NOT NULL,
    "sk_date"                date                 NOT NULL,
    "accreditation"          char(20)             NOT NULL,
    "accreditation_date"     date,
    "director_name"          varchar(50)          NOT NULL,
    "director_contact"       varchar(20)          NOT NULL,
    "pic_name"               varchar(50)          NOT NULL,
    "pic_contact"            varchar(20)          NOT NULL,
    "logo"                   char(36)             NOT NULL,
    "virtual_account_number" varchar(20),
    "account_number"         varchar(20)          NOT NULL,
    "account_name"           varchar(50)          NOT NULL,
    "account_bank_name"      varchar(20)          NOT NULL,
    "account_bank_code"      varchar(5)           NOT NULL,
    "created_at"             timestamp            NOT NULL,
    "updated_at"             timestamp            NOT NULL,
    "deleted_at"             timestamp
);


-- +migrate Down
DROP TABLE IF EXISTS "contacts";