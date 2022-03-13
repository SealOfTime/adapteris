import { FC } from "react";
import { ButtonBase, Card, Link, Typography, Stack, Box } from "@mui/material";
import { CheckIcon, XCircleIcon } from "@heroicons/react/outline";
import { SchoolEvent } from "../event-card/event-card";

export type UserEventParticipation = {
    name: string
    results?: number
    datetime?: Date
    place?: string
    organizers?: string[]
}

type ResultCardProps = {
    event: UserEventParticipation
}

export const ResultCard: FC<ResultCardProps> = ({ event }) => {
    const iconStyle = { width: '4rem', marginLeft: "15rem" };
    return (
        <>
            <Card variant="outlined">
                <Stack direction="row" justifyContent="left">
                    <Typography variant="h2" pt='15px' pl='1rem' gutterBottom>
                        {event.name}</Typography>
                    <Box marginRight="20rem">
                        {event.results >= 60 ? <CheckIcon style={iconStyle} /> :
                            <Box pt="0.5rem"> <XCircleIcon style={iconStyle} /></Box>}
                    </Box>
                </Stack>
            </Card>
        </>
    )
}

