import React, { FC } from "react";
import { EventCard, SchoolEvent } from "components/organisms/event-card/event-card";
import { Box, Stack, Typography } from "@mui/material";

export const FeedPage: FC = () => {
    const events = stubEvents;
    return (
        <>
            <Box pt='1.25rem' pb='64px'>
                <Typography variant="h1" gutterBottom>Предстоящие мероприятия:</Typography>
                <Stack spacing={2}>
                    {events.map(e => <EventCard key={e.id} event={e} />)}
                </Stack>
            </Box>
        </>
    );
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
