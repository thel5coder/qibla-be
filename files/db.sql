CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "register_from_enum" AS ENUM (
  'halaman_registrasi',
  'gmail',
  'facebook',
  'admin_travel',
  'registrasi_mitra'
);

CREATE TYPE "user_status_enum" AS ENUM (
  'guest',
  'calon_jamaah',
  'alumni',
  'agen',
  'travel',
  'tour_leader',
  'manajemen_travel'
);

CREATE TYPE "menu_permission_enums" AS ENUM (
  'View',
  'Add',
  'Edit',
  'Delete',
  'Approve'
);

CREATE TYPE "web_content_category_type_enum" AS ENUM (
  'gallery',
  'faq',
  'testimoni',
  'article'
);

CREATE TYPE "package_promotion_enum" AS ENUM (
  'per_hari',
  'bundle'
);

CREATE TYPE "platform_enum" AS ENUM (
  'semua',
  'mobile_apps',
  'web_compro'
);

CREATE TYPE "promotion_position_enum" AS ENUM (
  'slider_1',
  'slider_2',
  'slider_3',
  'slider_4',
  'slider_5'
);

CREATE TYPE "subcribtion_type_enum" AS ENUM (
  'subscribtion',
  'webinar',
  'website'
);

CREATE TYPE "price_unit_enum" AS ENUM (
  'jamaah',
  'hari',
  'jam'
);

CREATE TYPE "discount_type_enum" AS ENUM (
  'perhari',
  'fix'
);

CREATE TYPE "partner_status_enum" AS ENUM (
  'verified',
  'rejected',
  'unverified'
);

CREATE TYPE "status_transaction_enum" AS ENUM (
  'open',
  'close',
  'overdue'
);

CREATE TYPE "payment_status_enum" AS ENUM (
  'sukses',
  'gagal'
);

CREATE TYPE "transaction_direction_enum" AS ENUM (
  'in',
  'out'
);

CREATE TYPE "transaction_type_enum" AS ENUM (
  'tour',
  'zakat',
  'product_subscription',
  'tour_promotion'
);

CREATE TYPE "complaint_status_enum" AS ENUM (
  'open',
  'close'
);

CREATE TYPE "type_zakat_enum" AS ENUM (
  'maal',
  'penghasilan'
);

CREATE TYPE "identity_type_enum" AS ENUM (
  'ktp',
  'kk'
);

CREATE TYPE "marital_status_enum" AS ENUM (
  'belum_kawin',
  'kawin',
  'cerai_hidup',
  'cerai_mati'
);

CREATE TABLE "files" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" varchar(255),
  "path" varchar(50),
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "users" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "username" varchar(50) NOT NULL,
  "email" varchar(50) NOT NULL,
  "password" varchar(128) NOT NULL,
  "is_active" boolean,
  "role_id" char(36) NOT NULL,
  "register_at" timestamp,
  "activation_at" timestamp,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp,
  "name" varchar(50),
  "profile_picture" char(36),
  "mobile_phone" varchar(20),
  "pin" varchar(128),
  "fcm_device_token" varchar(255)
);

CREATE TABLE "roles" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "slug" varchar(50) NOT NULL,
  "name" varchar(30) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "menus" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "menu_id" char(3) NOT NULL,
  "name" varchar(20) NOT NULL,
  "url" text NOT NULL,
  "parent_id" char(36),
  "is_active" boolean,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "menu_permissions" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "menu_id" char(36) NOT NULL,
  "permission" menu_permission_enums,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "menu_permission_users" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "user_id" char(36) NOT NULL,
  "menu_permission_id" char(36) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "global_info_categories" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" varchar(30) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "global_info_category_settings" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "global_info_category_id" char(36) NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "web_comprof_categories" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "slug" text NOT NULL,
  "name" varchar(30) NOT NULL,
  "category_type" web_content_category_type_enum,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "galleries" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "web_content_category_id" char(36) NOT NULL,
  "gallery_name" varchar(30) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp NOT NULL
);

CREATE TABLE "gallery_images" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "gallery_id" char(36) NOT NULL,
  "file_id" char(36) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "testimonials" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "web_content_category_id" char(36) NOT NULL,
  "file_id" char(36) NOT NULL,
  "customer_name" varchar(50) NOT NULL,
  "job_position" varchar(30) NOT NULL,
  "testimoni" text NOT NULL,
  "rating" int,
  "is_active" boolean DEFAULT true,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "faqs" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "web_content_category_id" char(36) NOT NULL,
  "faq_category_name" varchar(50) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "faq_lists" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "faq_id" char(36) NOT NULL,
  "question" varchar(255) NOT NULL,
  "answer" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "term_conditions" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "term_name" varchar(50) NOT NULL,
  "term_type" varchar(30) NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "deleted_at" timestamp NOT NULL,
  "updated_at" timestamp
);

CREATE TABLE "promotion_packages" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "slug" text,
  "package_name" varchar(30) NOT NULL,
  "is_active" boolean DEFAULT true,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "promotions" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "promotion_package_id" char(36) NOT NULL,
  "package_promotion" package_promotion_enum DEFAULT 'perhari',
  "start_date" date,
  "end_date" date,
  "price" int4 NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "promotion_platforms" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "promotion_id" char(36) NOT NULL,
  "platform" varchar(30)
);

CREATE TABLE "promotion_positions" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "promotion_platform_id" char(36) NOT NULL,
  "position" varchar(30)
);

CREATE TABLE "tour_package_promotions" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "promotion_id" char(36) NOT NULL,
  "tour_package_id" char(36) NOT NULL,
  "partner_id" char(36) NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "tour_package_promotion_positions" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "tour_package_promotion_position_id" char(36) NOT NULL
);

CREATE TABLE "master_products" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "slug" varchar(255) NOT NULL,
  "name" varchar(30) NOT NULL,
  "subcription_type" subcribtion_type_enum NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "setting_products" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "product_id" char(36) NOT NULL,
  "price" int4 NOT NULL,
  "price_unit" price_unit_enum,
  "maintenance_price" int4,
  "discount" int4,
  "discount_type" discount_type_enum NOT NULL DEFAULT 'fix',
  "discount_period_start" date,
  "discount_period_end" date,
  "description" text NOT NULL,
  "sessions" varchar(30),
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "setting_product_periods" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "setting_product_id" char(36) NOT NULL,
  "period" int4 NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "setting_product_features" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "setting_product_id" char(36) NOT NULL,
  "feature_name" varchar(30) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "contacts" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "branch_name" varchar(50) NOT NULL,
  "travel_agent_name" varchar(50) NOT NULL,
  "address" text NOT NULL,
  "longitude" varchar(255),
  "latitude" varchar(255),
  "area_code" char(5) NOT NULL,
  "phone_number" char(20) NOT NULL,
  "sk_number" varchar(20) NOT NULL,
  "sk_date" date NOT NULL,
  "accreditation" char(20) NOT NULL,
  "accreditation_date" date,
  "director_name" varchar(50) NOT NULL,
  "director_contact" varchar(20) NOT NULL,
  "pic_name" varchar(50) NOT NULL,
  "pic_contact" varchar(20) NOT NULL,
  "logo" char(36) NOT NULL,
  "virtual_account_number" varchar(20),
  "account_number" varchar(20) NOT NULL,
  "account_name" varchar(50) NOT NULL,
  "account_bank_name" varchar(20) NOT NULL,
  "account_bank_code" varchar(5) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "partners" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "user_id" char(36) NOT NULL,
  "contact_id" char(36) NOT NULL,
  "contract_number" varchar(20),
  "web_binar_status" boolean,
  "website_status" boolean,
  "domain_site" varchar(30),
  "domain_erp" varchar(30),
  "database" varchar(20),
  "invoice_publish_date" date,
  "due_date" date,
  "due_date_aging" int,
  "verified_status" partner_status_enum DEFAULT 'unverified',
  "is_active" boolean,
  "reason" varchar(255),
  "product_id" char(36) NOT NULL,
  "subscription_period" int2,
  "subscription_expired_at" timestamp,
  "is_subscription_expired" boolean,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp,
  "verified_at" timestamp
);

CREATE TABLE "partner_extra_products" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "partner_id" char(36) NOT NULL,
  "product_id" char(36) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "transactions" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "user_id" char(36) NOT NULL,
  "invoice_number" varchar(20),
  "trx_id" varchar(30),
  "due_date" date,
  "periode_start_date" date,
  "period_end_date" data,
  "payment_status" payment_status_enum,
  "payment_method_code" char(3),
  "va_number" varchar(20),
  "bank_name" varchar(10),
  "direction" transaction_direction_enum,
  "transaction_type" transaction_type_enum,
  "paid_date" date,
  "transaction_date" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "transaction_details" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" varchar(255) NOT NULL,
  "fee" int4,
  "price" int4,
  "created_at" timestamp
);

CREATE TABLE "transaction_histories" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "trx_id" varchar(50),
  "status" varchar(30),
  "response" text,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "disbursements" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "transaction_id" char(36) NOT NULL,
  "is_disburse" boolean,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "calendars" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "title" varchar(50) NOT NULL,
  "start_date" datetime NOT NULL,
  "end_date" datetime,
  "description" text,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "calendar_participants" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "calendar_id" char(36) NOT NULL,
  "email" varchar(50) NOT NULL
);

CREATE TABLE "app_complaints" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "FullName" varchar(50) NOT NULL,
  "Email" varchar(50) NOT NULL,
  "ticket_number" char(6) NOT NULL,
  "complaint_type" varchar(30),
  "complaint" text,
  "solution" text,
  "is_open" complaint_status_enum DEFAULT 'open',
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "prays" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" varchar(50) NOT NULL,
  "file_id" char(36) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "video_contents" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "channel" varchar(100) NOT NULL,
  "links" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "master_zakats" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "slug" text,
  "type_zakat" type_zakat_enum,
  "name" varchar(50),
  "description" text,
  "amount" int4,
  "current_gold_price" int4,
  "gold_nishab" int4,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "satisfaction_categories" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "parent_id" char(36),
  "slug" varchar(255) NOT NULL,
  "name" varchar(100) NOT NULL,
  "is_active" bool,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "user_satisfactions" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "satisfaction_category_id" char(36) NOT NULL,
  "user_id" char(36) NOT NULL,
  "partner_id" char(36) NOT NULL,
  "rating" int4,
  "comments" text,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "crm_stories" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "slug" varchar(255) NOT NULL,
  "name" varchar(50) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "crm_boards" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "crm_story_id" char(36) NOT NULL,
  "contact_id" char(36) NOT NULL,
  "opportunity" varchar(255),
  "profit_expectation" float4,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "tour_packages" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "odoo_package_id" int2 NOT NULL,
  "details" json
);

CREATE TABLE "user_tour_purchase_participants" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "tour_package_id" char(36) NOT NULL,
  "user_id" char(36) NOT NULL,
  "is_new_jamaah" bool DEFAULT true,
  "identity_type" identity_type_enum NOT NULL DEFAULT 'ktp',
  "identity_number" varchar(20) NOT NULL,
  "full_name" varchar(50) NOT NULL,
  "sex" sex_enum NOT NULL,
  "birth_date" date NOT NULL,
  "birth_place" varchar(100) NOT NULL,
  "phone_number" varchar(20) NOT NULL,
  "city_id" char(36) NOT NULL,
  "marital_status" marital_status_enum NOT NULL DEFAULT 'kawin',
  "address" text NOT NULL,
  "kk_number" varchar(20) NOT NULL,
  "passpor_number" varchar(20),
  "passpor_name" varchar(50),
  "imigration_office_id" char(36) NOT NULL,
  "passpor_validity_period" date,
  "national_id_file" char(36) NOT NULL,
  "kk_file" char(36) NOT NULL,
  "birth_certificate" char(36),
  "marriage_certificate" char(36),
  "photo_3x4" char(36) NOT NULL,
  "photho_4x6" char(36) NOT NULL,
  "meningitis_free_certicate" char(36),
  "passpor_file" char(36),
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "user_tour_purchases" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "tour_package_id" char(36) NOT NULL,
  "customer_name" varchar(50) NOT NULL,
  "customer_identity_type" char(3),
  "identity_number" varchar(20),
  "full_name" varchar(50),
  "sex" sex_enum,
  "birth_date" date,
  "birth_place" varchar(100),
  "phone_number" varchar(20),
  "city_id" char(36),
  "marital_status" marital_status_enum,
  "customer_address" text,
  "user_id" char(36) NOT NULL,
  "contact_id" char(36) NOT NULL,
  "airline_odoo_id" int2,
  "airlines_class" varchar(50),
  "airline_price" float4,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "user_tour_purchase_rooms" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "room_odoo_id" int2,
  "room_class" varchar(30),
  "price" float4,
  "quantity" int2,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "user_tour_purchase_transaction" (
  "transaction_id" char(36) NOT NULL,
  "user_tour_purchase_id" char(36) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "user_zakats" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "user_id" char(36) NOT NULL,
  "transaction_id" char(36) NOT NULL,
  "contact_id" char(36) NOT NULL,
  "payment_method_code" int4,
  "master_zakat_id" char(36) NOT NULL,
  "type_zakat" type_zakat_enum,
  "current_gold_price" int4,
  "gold_nishab" int4,
  "wealth" int4,
  "total" int4,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "menu_permissions" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("id");

ALTER TABLE menu_user_permissions ADD FOREIGN KEY (menu_id) REFERENCES "users" ("id");

ALTER TABLE menu_user_permissions ADD FOREIGN KEY ("menu_permission_id") REFERENCES "menu_permissions" ("id");

ALTER TABLE "global_info_category_settings" ADD FOREIGN KEY ("global_info_category_id") REFERENCES "global_info_categories" ("id");

ALTER TABLE "galleries" ADD FOREIGN KEY ("web_content_category_id") REFERENCES "web_comprof_categories" ("id");

ALTER TABLE "gallery_images" ADD FOREIGN KEY ("gallery_id") REFERENCES "galleries" ("id");

ALTER TABLE "gallery_images" ADD FOREIGN KEY ("file_id") REFERENCES "files" ("id");

ALTER TABLE "testimonials" ADD FOREIGN KEY ("web_content_category_id") REFERENCES "web_comprof_categories" ("id");

ALTER TABLE "testimonials" ADD FOREIGN KEY ("file_id") REFERENCES "files" ("id");

ALTER TABLE "faqs" ADD FOREIGN KEY ("web_content_category_id") REFERENCES "web_comprof_categories" ("id");

ALTER TABLE "faq_lists" ADD FOREIGN KEY ("faq_id") REFERENCES "faqs" ("id");

ALTER TABLE "promotions" ADD FOREIGN KEY ("promotion_package_id") REFERENCES master_promotions ("id");

ALTER TABLE "promotion_platforms" ADD FOREIGN KEY ("promotion_id") REFERENCES "promotions" ("id");

ALTER TABLE "promotion_positions" ADD FOREIGN KEY ("promotion_platform_id") REFERENCES "promotion_platforms" ("id");

ALTER TABLE "tour_package_promotions" ADD FOREIGN KEY ("promotion_id") REFERENCES "promotions" ("id");

ALTER TABLE "tour_package_promotions" ADD FOREIGN KEY ("tour_package_id") REFERENCES "tour_packages" ("id");

ALTER TABLE "tour_package_promotions" ADD FOREIGN KEY ("partner_id") REFERENCES "partners" ("id");

ALTER TABLE "tour_package_promotion_positions" ADD FOREIGN KEY ("tour_package_promotion_position_id") REFERENCES "tour_package_promotions" ("id");

ALTER TABLE "setting_products" ADD FOREIGN KEY ("product_id") REFERENCES "master_products" ("id");

ALTER TABLE "setting_product_periods" ADD FOREIGN KEY ("setting_product_id") REFERENCES "setting_products" ("id");

ALTER TABLE "setting_product_features" ADD FOREIGN KEY ("setting_product_id") REFERENCES "setting_products" ("id");

ALTER TABLE "contacts" ADD FOREIGN KEY ("logo") REFERENCES "files" ("id");

ALTER TABLE "partners" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "partners" ADD FOREIGN KEY ("contact_id") REFERENCES "contacts" ("id");

ALTER TABLE "partners" ADD FOREIGN KEY ("product_id") REFERENCES "master_products" ("id");

ALTER TABLE "partner_extra_products" ADD FOREIGN KEY ("partner_id") REFERENCES "partners" ("id");

ALTER TABLE "partner_extra_products" ADD FOREIGN KEY ("product_id") REFERENCES "master_products" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "disbursements" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "calendar_participants" ADD FOREIGN KEY ("calendar_id") REFERENCES "calendars" ("id");

ALTER TABLE "prays" ADD FOREIGN KEY ("file_id") REFERENCES "files" ("id");

ALTER TABLE "user_satisfactions" ADD FOREIGN KEY ("satisfaction_category_id") REFERENCES "satisfaction_categories" ("id");

ALTER TABLE "user_satisfactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_satisfactions" ADD FOREIGN KEY ("partner_id") REFERENCES "partners" ("id");

ALTER TABLE "crm_boards" ADD FOREIGN KEY ("crm_story_id") REFERENCES "crm_stories" ("id");

ALTER TABLE "crm_boards" ADD FOREIGN KEY ("contact_id") REFERENCES "contacts" ("id");

ALTER TABLE "user_tour_purchase_participants" ADD FOREIGN KEY (user_tour_purchase_id) REFERENCES "tour_packages" ("id");

ALTER TABLE "user_tour_purchases" ADD FOREIGN KEY ("tour_package_id") REFERENCES "tour_packages" ("id");

ALTER TABLE "user_tour_purchases" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_tour_purchases" ADD FOREIGN KEY ("contact_id") REFERENCES "contacts" ("id");

ALTER TABLE "user_tour_purchase_transaction" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "user_tour_purchase_transaction" ADD FOREIGN KEY ("user_tour_purchase_id") REFERENCES "user_tour_purchases" ("id");
