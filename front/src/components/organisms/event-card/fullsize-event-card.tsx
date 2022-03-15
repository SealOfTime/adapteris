import React, { FC, useMemo, useState} from "react";
import { Card, Stack, CardContent, Button, Typography, Box, stepButtonClasses } from "@mui/material"
import { SchoolEvent, EventRegistration } from "components/organisms/event-card/event-card";
import { ClockIcon, UsersIcon, OfficeBuildingIcon } from "@heroicons/react/outline"

type FullsizeEventCardProps = {
        registration : EventRegistration
}

export const FullsizeEventCard: FC<FullsizeEventCardProps> = ({registration}) => {
    const[unregistered, registered] = useState(false);
    registration = stubEventRegistration
    const {day, month, hour, minute} = useMemo(()=>formatDate(registration.event.datetime), [registration.event.datetime])
    const iconStyle = { width: '1rem', marginRight: '0.5rem' }

    return (
        <>
        <Box pt='1.25rem'/>
                <Stack width="100%" pt='1.25rem' paddingLeft='0.5rem' direction="row" justifyContent="space-between">
                    <Typography variant="h1" gutterBottom>Мероприятие
                    </Typography>
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
                            {unregistered == false?
                            <Stack justifyContent="center">
                            <Typography align="center">Вы зарегистрированы на мероприятие:</Typography>
                             <Button variant="text" onClick={() => (registered(true))}>Отменить запись</Button>
                             </Stack>
                             : <Stack direction='column'>
                                 <Typography align="center">Вы не зарегистрированы на это мероприятие </Typography>
                                 <Button variant="text" onClick={() => (registered(false))}>Зарегистрироваться</Button>
                             </Stack> }
                        </Stack>
                        </Box>
                </CardContent>
            </Card>
        </>
    )
}

function formatDate(date: Date) {
    const day = date.getDate();
    const month = monthName[date.getMonth()-1];
    const hour = date.getHours();
    const minute = date.getMinutes();
    return {
        day,
        month,
        hour: hour < 10 ? `0${hour}` : hour,
        minute: minute < 10 ? `0${minute}` : minute
    }
}

const monthName = [
    'января',
    'февраля',
    'марта',
    'апреля',
    'мая',
    'июня',
    'июля',
    'сентября',
    'октября',
    'ноября',
    'декабря'
]

const stubEventRegistration: EventRegistration =
{
    id: 11,
    userId: 123,
    isRegistered: true,
    event:   {
        id: 11,
        name: "Игротехника",
        datetime: new Date(),
        place: "Ломоносова, 9. ауд. 1220",
        organizers: ['Вдовицын М.В.', 'Суязова И.М.']
    }
}
