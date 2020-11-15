-- +migrate Up
CREATE TABLE IF NOT EXISTS "user_tour_purchase_participants"
(
    "id"                          char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "user_tour_purchase_id"             char(36)             NOT NULL,
    "user_id"                     char(36)             NOT NULL,
    "is_new_jamaah"               bool                          DEFAULT true,
    "identity_type"               identity_type_enum   NOT NULL DEFAULT 'ktp',
    "identity_number"             varchar(20)          NOT NULL,
    "full_name"                   varchar(50)          NOT NULL,
    "sex"                         sex_enum             NOT NULL,
    "birth_date"                  date                 NOT NULL,
    "birth_place"                 varchar(100)         NOT NULL,
    "phone_number"                varchar(20)          NOT NULL,
    "city_id"                     char(36)             NOT NULL,
    "marital_status"              marital_status_enum  NOT NULL DEFAULT 'kawin',
    "address"                     text                 NOT NULL,
    "kk_number"                   varchar(20)          NOT NULL,
    "passport_number"             varchar(20),
    "passport_name"               varchar(50),
    "immigration_office"          varchar(50)          NOT NULL,
    "passport_validity_period"    date,
    "national_id_file"            char(36)             NOT NULL,
    "kk_file"                     char(36)             NOT NULL,
    "birth_certificate"           char(36),
    "marriage_certificate"        char(36),
    "photo_3x4"                   char(36)             NOT NULL,
    "photo_4x6"                   char(36)             NOT NULL,
    "meningitis_free_certificate" char(36),
    "passport_file"               char(36),
    "is_depart"                   bool,
    "status"                      varchar(50),
    "created_at"                  timestamp            NOT NULL,
    "updated_at"                  timestamp            NOT NULL,
    "deleted_at"                  timestamp
);


-- +migrate Down
DROP TABLE user_tour_purchase_participants;