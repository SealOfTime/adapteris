import React, {FC, useState} from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import {formatDate} from 'hooks/useFormatDate';
import {EventSession} from "hooks/useEvent";
import {Box, Button, Dialog, DialogActions, DialogContent, DialogTitle, Typography} from "@mui/material";
import {useRegisterForSession} from "hooks/useMyEventParticipations";
import {useCanEdit} from "hooks/useAuth";

type EventTimetableProps = {
    sessions: EventSession[]
}

export const SessionsTable: FC<EventTimetableProps> = ({sessions}) => {
    const [selectedSession, selectSession] = useState<EventSession | null>(null)
    const canEdit = useCanEdit()
    const now = new Date();

    const eventsRows = sessions.map(session => {
        const {hour, minute, month, day} = formatDate(session.datetime);
        const clickable = canEdit || new Date(session.datetime) > now;
        return (
            <TableRow key={session.id} onClick={() => clickable && selectSession(session)} sx={{
            ...(clickable
                ? {'&:hover': {
                        backgroundColor: 'grey.200',
                        cursor: 'pointer',
                    }}
                : {}),
            }}>
                <TableCell>{day} {month}</TableCell>
                <TableCell align="left">{hour}:{minute}</TableCell>
                <TableCell align="left">{session.place}</TableCell>
            </TableRow>
        )
    });
    return (
        <>
            <TableContainer component={Paper} sx={{width: '100%'}}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell align="left">Дата</TableCell>
                            <TableCell align="left">Время</TableCell>
                            <TableCell align="left">Место</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {eventsRows}
                    </TableBody>
                </Table>
            </TableContainer>
            {canEdit
                ? <ParticipantsListDialog session={selectedSession} close={()=>selectSession(null)} />
                : <ConfirmRegisterDialog session={selectedSession} close={()=>selectSession(null)}/>
            }
        </>
    );
}
const ParticipantsListDialog: FC<{session: EventSession | null, close: ()=>void,}> = ({session, close}) => {
    return (
        <Dialog open={session != null}>
            <Box sx={{padding: '2rem'}}>
                <Typography variant="h3">Список участников</Typography>
            </Box>
        </Dialog>
    );
}

const ConfirmRegisterDialog: FC<{session: EventSession | null, close: ()=>void,}> = ({session, close}) => {
    const registerForSession = useRegisterForSession();
    const confirm = () => {
        registerForSession({sessionId: session?.id}, {
            onSuccess: close,
        })
    }
    return (
        <Dialog open={session != null}>
            <DialogTitle><Typography variant="h3">Запись на мероприятие</Typography></DialogTitle>
            <DialogContent>
                Вы уверены, что хотите записаться на это проведение мероприятия?
            </DialogContent>
            <DialogActions>
                <Button onClick={confirm}>Да</Button>
                <Button onClick={close}>Я ещё подумаю</Button>
            </DialogActions>
        </Dialog>
    )
}
