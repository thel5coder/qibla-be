
-- +migrate Up
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
    'view',
    'add',
    'edit',
    'delete',
    'approve'
    );

CREATE TYPE "web_content_category_type_enum" AS ENUM (
    'galery',
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

CREATE TYPE "status_transaction_enum" AS ENUM (
    'open',
    'close',
    'overdue'
    );

CREATE TYPE "payment_status_enum" AS ENUM (
    'sukses',
    'gagal'
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

CREATE TYPE "sex_enum" AS ENUM (
    'female',
    'male'
    );
-- +migrate Down
