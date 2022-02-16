import { FC, useMemo } from "react";
import { Card, Stack, CardContent, Button, Typography } from "@mui/material"
import { EventCard, SchoolEvent } from "components/organisms/event-card/event-card";
import { ClockIcon, UsersIcon, OfficeBuildingIcon } from "@heroicons/react/outline"

type EventRegistration = {
    name: string
    place: string
    description?: string
    datetime: Date
    organizers: string[]
    isRegistered: boolean
}

type FullsizeEventCardProps = {
    registration: EventRegistration
}

export const FullsizeEventCard: FC<SchoolEvent> = (event) => {
    const {day, month, hour, minute} = useMemo(()=>formatDate(event.datetime), [event.datetime])
    const iconStyle = { width: '1rem', marginRight: '0.5rem' };
    return (
        <>
            <Card variant="outlined" >
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
                    </Stack>
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


