import React, {FC, useMemo, useState} from "react";
import {useParams} from 'react-router-dom'
import {
    Box, Button,
    Card,
    CardContent,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    Stack, TextField,
    Tooltip,
    Typography
} from "@mui/material"
import {ClockIcon, OfficeBuildingIcon} from "@heroicons/react/outline"
import {SessionsTable} from "../organisms/event-card/sessions-table"
import {useAuthenticatedUser, useCanEdit} from "hooks/useAuth";
import {
    EventParticipation,
    ParticipationEventSession,
    useCancelSessionParticipation,
    useMyEventParticipations
} from "hooks/useMyEventParticipations";
import {SchoolEvent, useAddSession, useEvent} from "hooks/useEvent";
import {Spinner} from "components/organisms/spinner";
import {useFormatDate} from "hooks/useFormatDate";
import {AddButton} from "components/organisms/school-tree/add-button";
import {DateTimePicker} from "@mui/lab";

export const EventPage: FC = () => {
    const user = useAuthenticatedUser();
    const canEdit = useCanEdit();

    const {eventId: eventIdRaw} = useParams();
    const eventId = parseInt(eventIdRaw);
    const event = useEvent(eventId);

    const participations = useMyEventParticipations();
    console.log(participations?.data)
    const registeredFor = useMemo(()=> {
            return participations.data?.find(p => p.session.event.id === eventId && p.passed != false)
    }, [participations, eventId])
    if(event.isLoading || !event.data) {
        return <Spinner/>
    }

    return (
        <>
            <Box pt='1.25rem'/>
            <Typography variant="h1" gutterBottom>Мероприятие</Typography>
            <Card variant="outlined">
                <CardContent>
                    <Stack>
                        <Typography variant="h3" gutterBottom>{event.data.name}</Typography>
                        {event.data.description &&
                            <Typography gutterBottom>{event.data.description}</Typography>
                        }
                        {registeredFor && <SessionInfo participation={registeredFor}/>}
                    </Stack>
                    <Stack alignItems="center" alignContent="stretch">
                    {canEdit && "Выберите проведение, чтобы посмотреть список участников"}
                    {!canEdit && !registeredFor && "Выберите удобное для вас время и место"}
                    {!registeredFor && <SessionsTable sessions={event.data.sessions}/>}
                    {canEdit && <AddEventSessionButton event={event.data}/>}
                    </Stack>
                </CardContent>
            </Card>
        </>
    )
}

const AddEventSessionButton: FC<{ event: SchoolEvent}> = ({event}) => {
    const [dialog, setDialog] = useState<boolean>(false);
    const closeDialog = () => setDialog(false);
    const openDialog = ()=>setDialog(true);

    const [place, setPlace] = useState<string>("");
    const onPlaceChange = (e) => setPlace(e.target.value)

    const [datetime, setDatetime] = useState<Date>(new Date());

    const addSessionReq = useAddSession();
    const submit = () => {
        addSessionReq({eventId: event.id, session: {place: place, datetime: datetime}})
    }
    return (
        <>
        <Tooltip title="Добавить проведение" sx={{margin: 'auto'}}><AddButton onClick={openDialog}/></Tooltip>
            <Dialog open={dialog}>
                <DialogTitle><Typography variant="h3">Добавить проведение</Typography></DialogTitle>
                <DialogContent>
                    <Stack spacing="3rem" pt="1rem">
                    <TextField label="Место проведения" value={place} onChange={onPlaceChange}/>
                    <DateTimePicker
                        label="Время и дата проведения"
                        value={datetime}
                        onChange={setDatetime}
                        renderInput={(params) => <TextField {...params} />}
                    />
                    </Stack>
                </DialogContent>
                <DialogActions>
                    <Button onClick={submit}>Сохранить</Button>
                    <Button onClick={closeDialog}>Закрыть</Button>
                </DialogActions>
            </Dialog>
        </>
    )
}
const iconStyle = { width: '1rem', marginRight: '0.5rem' };
export const SessionInfo: FC<{participation: EventParticipation}> = ({participation}) => {
    const {day, month, hour, minute} = useFormatDate(participation.session.datetime)
    const sendCancelRegistrationReq = useCancelSessionParticipation();
    const cancelRegistration = () => sendCancelRegistrationReq(
        {participantId: participation.id},
        )
    return (
        <>
            <Typography variant="subtitle1" display="flex">
                <ClockIcon style={iconStyle} />
                <span>{day} {month} в {hour}:{minute}</span>
            </Typography>
            <Typography variant="subtitle1" display="flex">
                <OfficeBuildingIcon style={iconStyle} />
                <span>{participation.session.place}</span>
            </Typography>
            {participation.passed === null &&
                <Button variant="contained" onClick={cancelRegistration}>Отменить запись</Button>
            }
        </>
    )
}
