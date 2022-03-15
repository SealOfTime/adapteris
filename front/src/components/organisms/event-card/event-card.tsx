import React, { FC, useMemo, useState } from "react";
import { Card, Stack, CardContent, Typography, Button, IconButton } from "@mui/material";
import { ClockIcon, OfficeBuildingIcon, UsersIcon, ArrowCircleRightIcon } from "@heroicons/react/outline";
import { FullsizeEventCard } from "../../pages/fullsize-event-card";
import { Link } from 'react-router-dom'
import Dialog from '@mui/material/Dialog';



export type SchoolEvent = {
    id: number
    name: string
    description?: string
    datetime: Date
    place: string
    organizers: string[]
}

type EventCardProps = {
    event: SchoolEvent
}

export type EventRegistration = {
    id: number
    userId: number
    isRegistered: boolean,
    event: SchoolEvent
}

export const EventCard: FC<EventCardProps> = ({ event }) => {
    const { day, month, hour, minute } = useMemo(() => formatDate(event.datetime), [event.datetime])
    const iconStyle = { width: '1rem', marginRight: '0.5rem' };
    const buttonStyle = { width: '2rem' }
    const [registration, setRegistration] = useState('')
    function openFullsizeCard({stubEventRegistration}) {
    setRegistration(stubEventRegistration)
    return (
        <FullsizeEventCard registration={useState.arguments[0]}/>
    )
}
    return (
        <>
        <Card variant="outlined">
            <CardContent>
                <Stack direction="row" justifyContent="space-between">
                    <Stack direction="column">
                        <Typography variant="h3" gutterBottom>{event.name}</Typography>
                        <Typography variant="subtitle1" display="flex">
                            <ClockIcon style={iconStyle} />
                            <span>{day} {month} в {hour}:{minute}</span>
                        </Typography>
                        <Typography variant="subtitle1" display="flex">
                            <OfficeBuildingIcon style={iconStyle} />
                            <span>{event.place}</span>
                        </Typography>
                        <Typography variant="subtitle1" display="flex">
                            <UsersIcon style={iconStyle} />
                            <span>{event.organizers.join(", ")}</span>
                        </Typography>
                    </Stack>
                    <Link to="/event">
                        <IconButton onClick={() => (
                        openFullsizeCard({stubEventRegistration}))}><ArrowCircleRightIcon style={buttonStyle} /></IconButton>
                    </Link>
                </Stack>
            </CardContent>
        </Card>
        </>
    )
}

function formatDate(date: Date) {
    const day = date.getDate();
    const month = monthName[date.getMonth() - 1];
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

