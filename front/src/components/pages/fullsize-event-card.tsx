import { FC, useMemo, useState} from "react";
import { Link } from 'react-router-dom'
import { Card, Stack, CardContent, Button, Typography, Box, stepButtonClasses} from "@mui/material"
import { SchoolEvent } from "components/organisms/event-card/event-card";
import { ClockIcon, UsersIcon, OfficeBuildingIcon } from "@heroicons/react/outline"
import { EventTimetable } from "../organisms/event-card/event-timetable"
import { useFormatDate } from "hooks/useFormatDate"
import { RegsTimetable } from "./organizers/org-event-timetable";

type EventRegistration = {
    id: number
    userId: number
    isRegistered: boolean,
    event: SchoolEvent
}

type FullsizeEventCardProps = {
    registration: EventRegistration
}

export const FullsizeEventCard: FC<FullsizeEventCardProps> = ({ registration }) => {
    const [unregistered, registered] = useState(false);
    registration = stubEventRegistration
    const { day, month, hour, minute } = useFormatDate(registration.event.datetime)
    const iconStyle = { width: '1rem', marginRight: '0.5rem' }
    /* потом удалить */
    const [organizer, notOrganizer] = useState(false);
    return (
        <>
            <Box pt='1.25rem'/>
            <Stack width="100%" pt='1.25rem' paddingLeft='0.5rem' direction="row" justifyContent="space-between">
                <Typography variant="h1" gutterBottom>Мероприятие
                </Typography>        
        <Button variant="outlined" onClick={() => notOrganizer(true)}>Для организаторов типа</Button>
        {(organizer==true? <Typography> Вы орагнизатор
                </Typography> : <Typography>Вы не орагнутан
                </Typography> )}
            </Stack>
            <Card variant="outlined">
                <CardContent>
                    <Stack>
                        <Typography variant="h3" gutterBottom>{registration.event.name}</Typography>
                        <Typography variant="subtitle1" display="flex">
                            <ClockIcon style={iconStyle} />
                            <span>{day} {month} в {hour}:{minute}</span>
                        </Typography>
                        <Typography variant="subtitle1" display="flex">
                            <OfficeBuildingIcon style={iconStyle} />
                            <span>{registration.event.place}</span>
                        </Typography>
                        <Typography variant="subtitle1" display="flex">
                            <UsersIcon style={iconStyle} />
                            <span>{registration.event.organizers.join(", ")}</span>
                        </Typography>
                    </Stack>
                    <Box marginTop="1rem">
                        <Stack direction="column" justifyContent="center">
                            {unregistered == false ?
                                <Stack justifyContent="center">
                                    <Typography align="center" marginBottom="1rem">Вы зарегистрированы на это мероприятие</Typography>
                                    <Button variant="outlined" onClick={() => (registered(true))}>Отменить запись</Button>
                                </Stack>
                                : <Stack direction='column' space-between="10px" pb="5rem">
                                    <Typography align="center" marginBottom="1rem">Вы не зарегистрированы на это мероприятие </Typography>
                                    <Typography align="center" variant="h4" gutterBottom> Доступные слоты </Typography>
                                    <EventTimetable events={stubEvents}></EventTimetable>
                                    <Button variant="outlined" onClick={() => (registered(false))}>Зарегистрироваться</Button>
                                </Stack>}
                        </Stack>
                    </Box>
                </CardContent>
            </Card>
        </>
    )
}

const stubEventRegistration: EventRegistration =
{
    id: 11,
    userId: 123,
    isRegistered: true,
    event: {
        id: 11,
        name: "Игротехника",
        datetime: new Date(),
        place: "Ломоносова, 9. ауд. 1220",
        organizers: ['Вдовицын М.В.', 'Суязова И.М.']
    }
}

const stubEvents: SchoolEvent[] = [
    {
        id: 11,
        name: "Игротехника",
        datetime: new Date(),
        place: "Ломоносова, 9. ауд. 1220",
        organizers: ['Вдовицын М.В.', 'Суязова И.М.']
    },
    {
        id: 22,
        name: "Игротехника",
        datetime: new Date(),
        place: "Ломоносова, 9. ауд. 1220",
        organizers: ['Вдовицын М.В.', 'Суязова И.М.']
    },
    {
        id: 33,
        name: "Игротехника",
        datetime: new Date(),
        place: "Ломоносова, 9. ауд. 1220",
        organizers: ['Вдовицын М.В.', 'Суязова И.М.']
    },
    {
        id: 44,
        name: "Игротехника",
        datetime: new Date(),
        place: "Ломоносова, 9. ауд. 1220",
        organizers: ['Вдовицын М.В.', 'Суязова И.М.']
    },
    {
        id: 55,
        name: "Игротехника",
        datetime: new Date(),
        place: "Ломоносова, 9. ауд. 1220",
        organizers: ['Вдовицын М.В.', 'Суязова И.М.']
    },
];
