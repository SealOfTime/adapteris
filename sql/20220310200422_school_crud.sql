-- +goose Up
-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
CREATE TABLE school (
    "id" serial PRIMARY KEY,
    "name" text NOT NULL UNIQUE CHECK ("name" <> ''),
    "visible" boolean NOT NULL DEFAULT FALSE,
    "start_date" date,
    "end_date" date,
    "register_start" date,
    "register_end" date
);

CREATE TABLE stage (
    "id" serial PRIMARY KEY,
    "name" text NOT NULL UNIQUE CHECK ("name" <> ''),
    "description" text,
    "start_date" date,
    "end_date" date,
    "school_id" int NOT NULL REFERENCES school ON DELETE CASCADE
);

CREATE TABLE step (
    "id" bigserial PRIMARY KEY,
    "must_complete" int NOT NULL DEFAULT -1,
    "stage_id" bigint NOT NULL REFERENCES stage ON DELETE CASCADE
);

CREATE TABLE "event" (
    "id" bigserial PRIMARY KEY,
    "name" text NOT NULL CHECK ("name" <> ''),
    "description" text,
    "step_id" bigint NOT NULL REFERENCES step ON DELETE CASCADE
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS event;

DROP TABLE IF EXISTS step;

DROP TABLE IF EXISTS stage;

DROP TABLE IF EXISTS school;

-- +goose StatementEnd
