--User
INSERT INTO "User" (id, full_name, short_name, is_admin)
VALUES (0, 'Вдовицын Матвей Валентинович', 'Матвей', false);

INSERT INTO "User" (id, full_name, short_name, is_admin)
VALUES (2, 'Горбунова Екатерина Дмитриевна', 'Катя', false);

INSERT INTO "User" (id, full_name, short_name, is_admin)
VALUES (1, 'Шведова Мария Сергеевна', 'Маша', true);


--External Accounts
INSERT INTO "External_Account" (user_id, service, external_id)
VALUES (0, 'vk', '159087468');

INSERT INTO "External_Account" (user_id, service, external_id)
VALUES (1, 'vk', '159843705');

INSERT INTO "External_Account"(user_id, service, external_id)
VALUES (2, 'vk', '315112267');

--School
INSERT INTO "Adapter_School" (id, created_by, name, start_date, end_date)
VALUES (0, 1, 'Школа Адаптеров 2021', date '01.04.2021', date '04.07.2021');

INSERT INTO "School_Participant" (id, school_id, user_id, role)
VALUES (0, 0, 0, 'STUDENT');

INSERT INTO "School_Participant" (id, school_id, user_id, role)
VALUES (1, 0, 1, 'ORGANIZER');

INSERT INTO "School_Participant" (id, school_id, user_id, role)
VALUES (2, 0, 1, 'STUDENT');

INSERT INTO "School_Stage" (id, school_id, name, description)
VALUES (0, 0, 'Игротехника', 'Первый этап школы адаптеров');

--Event
INSERT INTO "School_Event" (id, stage_id, name, description, type)
VALUES (0, 0, 'Учимся выгонять людей', 'На этом занятии вы научитесь выгонять людей, которые вам мешают', 'TRAINING');

INSERT INTO "School_Event_Session" (event_id, name, description, place, starts_at, ends_at)
VALUES (0, 'DEFAULT', '', 'Ул. Ломоносова, 9', timestamp '10.04.2021 12:00', timestamp '10.04.2021 14:30');

INSERT INTO "School_Event_Session" (event_id, name, description, place, starts_at, ends_at)
VALUES (1, 'DEFAULT', '', 'Кронверкский пр. 49, ауд. 228', timestamp '12.04.2021 12:00', timestamp '12.04.2021 14:30');


INSERT INTO "School_Event_Session" (event_id, name, description, place, starts_at, ends_at)
VALUES (2, 'DEFAULT', '', 'Кронверкский пр. 49, ауд. 228', timestamp '05.05.2021 12:00', timestamp '05.05.2021 14:30');

INSERT INTO "School_Event_Session" (event_id, name, description, place, starts_at, ends_at)
VALUES (3, 'DEFAULT', '', 'Кронверкский пр. 49, ауд. 228', timestamp '05.05.2021 15:00', timestamp '05.05.2021 17:30');

