-- +goose Up
-- +goose StatementBegin
CREATE TABLE "event_session" (
    id bigserial PRIMARY KEY,
    event_id bigint NOT NULL REFERENCES "event" ON DELETE CASCADE,
    max_participants int NOT NULL,
    place text NOT NULL,
    time timestamp NOT NULL
);

CREATE TYPE "school_result" AS ENUM (
    'DECLINE',
    'ACCEPT'
);

CREATE TABLE "school_participant" (
    id bigserial PRIMARY KEY,
    school_id int NOT NULL REFERENCES "school" ON DELETE CASCADE,
    account_id bigint NOT NULL REFERENCES "account" ON DELETE CASCADE,
    result school_result,
    result_comment text NOT NULL
);

CREATE TABLE "event_participant" (
    id bigserial PRIMARY KEY,
    participant_id bigint NOT NULL REFERENCES "school_participant",
    event_session_id bigint NOT NULL REFERENCES "event_session",
    passed boolean
);

CREATE OR REPLACE FUNCTION register_participant (school_participant_id bigint, session_id bigint)
    RETURNS boolean
    AS $$
BEGIN
    IF EXISTS ( WITH event_id AS ( SELECT event_id FROM "event_session" WHERE id = session_id)
                SELECT
                    1 
                FROM "event_participant"
                    INNER JOIN "event_session" ON "event_session".id = "event_participant".event_session_id
                WHERE
                    "event_session".event_id = event_id 
                    AND "event_participant".participant_id = school_participant_id
                    AND "event_participant".passed <> false
        ) THEN RETURN false; END IF;
    RETURN true;
END;
$$
LANGUAGE plpgsql;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "event_session";

DROP TYPE "school_result";

DROP TABLE "school_participant";

DROP TABLE "event_participant";

-- +goose StatementEnd
