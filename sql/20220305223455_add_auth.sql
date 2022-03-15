-- +goose Up
-- +goose StatementBegin
CREATE TYPE site_role AS ENUM (
    'USER',
    'ADMIN'
);

CREATE DOMAIN email AS varchar(320)
CHECK (value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$');

CREATE TABLE "account" (
    "id" serial PRIMARY KEY,
    "registered_at" date NOT NULL DEFAULT now(),
    "role" site_role NOT NULL DEFAULT 'USER',
    "full_name" text,
    "short_name" varchar(255) NOT NULL CHECK ("short_name" <> ''),
    "email" email NOT NULL UNIQUE,
    "telegram" varchar(32),
    "vk" varchar(32),
    "phone_number" varchar(32)
);

CREATE TABLE "external_account" (
    "id" serial PRIMARY KEY,
    "account_id" integer NOT NULL REFERENCES "account" ON DELETE CASCADE,
    "service" varchar(64),
    "external_id" text
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "external_account";

DROP TABLE "account";

DROP DOMAIN email;

DROP TYPE site_role;

-- +goose StatementEnd
