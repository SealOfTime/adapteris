--User
CREATE TABLE "User"
(
    "id"            serial PRIMARY KEY,
    "full_name"     text,
    "short_name"    varchar(255) NOT NULL CHECK ("short_name" <> ''),
    "registered_at" date         NOT NULL,
    "is_admin"      boolean      NOT NULL
);

CREATE TABLE "External_Account"
(
    "id"          serial PRIMARY KEY,
    "user_id"     integer NOT NULL references "User"
        ON DELETE CASCADE,
    "service"     varchar(64),
    "external_id" text
);

ALTER TABLE "External_Account"
    ADD CONSTRAINT "External_Account_fk0" FOREIGN KEY ("user_id") REFERENCES "User" ("id");

--School
CREATE TABLE "Adapter_School"
(
    "id"         serial PRIMARY KEY,
    "created_by" integer references "User" ("id")
        ON DELETE SET NULL,
    "name"       varchar(255) NOT NULL
        CHECK ("name" <> ''),
    "start_date" date,
    "end_date"   date
);

CREATE TABLE "School_Participant"
(
    "id"        serial PRIMARY KEY,
    "school_id" integer      NOT NULL references "Adapter_School" ("id")
        ON DELETE RESTRICT,
    "user_id"   integer      NOT NULL references "User" ("id")
        ON DELETE RESTRICT,
    role        varchar(255) NOT NULL CHECK (role <> '')
);

CREATE FUNCTION isOrganizer(org_id integer)
    RETURNS boolean
AS $$
    BEGIN
        IF EXISTS (SELECT 1 FROM "School_Participant" WHERE "id" = org_id AND "role" = 'ORGANIZER') THEN
            return true;
        END IF;
        return false;
    END;
$$ LANGUAGE plpgsql;

CREATE TABLE "School_Stage"
(
    "id"          serial PRIMARY KEY,
    "school_id"   integer      NOT NULL references "Adapter_School" ("id")
        ON DELETE CASCADE,
    "name"        varchar(255) NOT NULL CHECK ("name" <> ''),
    "description" text
);

--Event
CREATE TABLE "School_Event"
(
    "id"          serial PRIMARY KEY,
    "stage_id"    integer      NOT NULL references "School_Stage"
        ON DELETE CASCADE,
    "name"        varchar(255) NOT NULL CHECK ("name" <> ''),
    "description" text,
    "type"        varchar(255) NOT NULL CHECK ("type" <> '')
);

CREATE TABLE "Event_Participation_Policy"(
    "participant_id" integer
        REFERENCES "School_Participant"("id"),
    "event_id" integer
        REFERENCES "School_Event"("id"),
    "policy" varchar(64),
    CONSTRAINT "Event_Participation_Policy_pk" PRIMARY KEY ("participant_id", "event_id")
);

CREATE TABLE "School_Event_Session"
(
    "id"               serial PRIMARY KEY,
    "event_id"         integer      NOT NULL
        REFERENCES "School_Event" ("id")
        ON DELETE CASCADE,
    "name"             varchar(255) NOT NULL CHECK ("name" <> ''),
    "description"      text,
    "place"            text         NOT NULL CHECK ("place" <> ''),
    "starts_at"        timestamp    NOT NULL,
    "ends_at"          timestamp    NOT NULL,
    "max_participants" integer      NOT NULL
);

CREATE TABLE "School_Event_Session_Organizer"
(
    "organizer_id"     integer NOT NULL
        REFERENCES "School_Participant" ("id") ON DELETE CASCADE
        CHECK (isOrganizer("organizer_id")),
    "event_session_id" integer NOT NULL
        REFERENCES "School_Event_Session" ("id") ON DELETE CASCADE,
    CONSTRAINT "School_Event_Session_Organizer_pk" PRIMARY KEY ("organizer_id", "event_session_id")
);

CREATE TABLE "Event_Participation"
(
    "id"                serial PRIMARY KEY,
    "event_session_id"  integer NOT NULL
        REFERENCES "School_Event_Session" ("id") ON DELETE CASCADE,
    "student_id"        integer NOT NULL
        REFERENCES "School_Participant" ("id") ON DELETE CASCADE,
    "organizer_comment" text,
    "is_credited"       boolean NOT NULL
        DEFAULT false
);

--Event Grading
CREATE TABLE "Event_Grade_Criteria"
(
    "id"          serial PRIMARY KEY,
    "event_id"    integer       NOT NULL
        REFERENCES "School_Event" ("id") ON DELETE CASCADE,
    "name"        varchar(255)  NOT NULL,
    "description" text,
    "min"         decimal(5, 3),
    "max"         decimal(5, 3)
);

CREATE TABLE "Participation_Result"
(
    "id"                serial PRIMARY KEY,
    "participant_id"    integer NOT NULL
        REFERENCES "Event_Participation" ("id") ON DELETE CASCADE,
    "criteria_id"       integer NOT NULL
        REFERENCES "Event_Grade_Criteria" ("id") ON DELETE CASCADE,
    "value" decimal(5,3) NOT NULL DEFAULT(0.0),
    "organizer_comment" text
);