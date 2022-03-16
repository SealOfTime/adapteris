import React, {FC} from "react";
import {Box, Card, Stack, Typography} from "@mui/material";
import {useAuthenticatedUser} from "hooks/useAuth";
import {Spinner} from "components/organisms/spinner";
import {ParticipationEventSession, useMyEventParticipations} from "hooks/useMyEventParticipations";
import {ClockIcon, OfficeBuildingIcon} from "@heroicons/react/outline";
import {useFormatDate} from "hooks/useFormatDate";
import {useNavigate} from "react-router-dom";

export const FeedPage: FC = () => {
    const user = useAuthenticatedUser();
    const participations = useMyEventParticipations()
    if (!user) {
        return <Spinner />
    }

    return (
        <>
            <Box pt='1.25rem'>
                <Typography variant="h1" gutterBottom>Предстоящие мероприятия:</Typography>
                <Stack spacing={2} sx={{height: '100%'}}>
                    {!participations.data?.length &&
                        <Typography fontSize="1.2rem" textAlign="center" mt='50%'>Пока здесь пусто, выбирай мероприятия школы и записывайся!</Typography>
                    }
                    {participations.data?.map(p =>
                        <EventSessionCard key={p.id} session={p.session}/>
                    )}
                </Stack>
            </Box>
        </>
    );
}

const iconStyle = { width: '1rem', marginRight: '0.5rem' };
const EventSessionCard: FC<{session: ParticipationEventSession}> = ({session}) => {
    const {day, month, hour, minute} = useFormatDate(session.datetime)
    const navigate = useNavigate();
    const openFullPage = () => navigate(`/event/${session.event.id}`)

    return (
        <Card variant="outlined" sx={{
            padding: '2rem',
            transition: 'all 0.1s ease-out',
            '&:hover': {
                backgroundColor: 'grey.200',
                cursor: 'pointer',
            },
        }} onClick={openFullPage}>
            <Typography variant="h2" gutterBottom>{session.event.name}</Typography>
            <Typography variant="subtitle1" display="flex">
                <ClockIcon style={iconStyle} />
                <span>{day} {month} в {hour}:{minute}</span>
            </Typography>
            <Typography variant="subtitle1" display="flex">
                <OfficeBuildingIcon style={iconStyle} />
                <span>{session.place}</span>
            </Typography>
        </Card>
    )
}
