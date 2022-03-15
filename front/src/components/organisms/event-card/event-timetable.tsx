import { FC, useMemo, useState } from 'react';
import { EventRegistration, SchoolEvent } from "./event-card";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell, { tableCellClasses } from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { formatDate } from 'hooks/useFormatDate';
import { Button, Dialog, DialogTitle, List, ListItem, ListItemText } from '@mui/material';
import { UserCard } from '../user-card';

type EventTimetableProps = {
    events: SchoolEvent[]
}

export const EventTimetable: FC<EventTimetableProps> = ({ events }) => {
    let registrations = stubEventRegistrations
    const [open, setOpen] = useState(false)
    const eventsRows = events.map(e => {
        const {hour, minute, month, day} = formatDate(e.datetime);
        const organizers = (e.organizers).join(' ')
        return (
            <TableRow onClick={() => setOpen(true)}>
                <TableCell>{day} {month}</TableCell>
                <TableCell align="left">{hour}:{minute}</TableCell>
                <TableCell align="left">{e.place}</TableCell>
                <TableCell align="left">{organizers}</TableCell>
            </TableRow>)});
    return (
        <>
        <Paper>
                <TableContainer component={Paper}>
                            <Table sx={{ minWidth: 500 }}>
                                <TableHead>
                                <TableRow>
                                <Dialog onClose={() => {setOpen(false)}} open={open}>
                                <DialogTitle>Список участников</DialogTitle>
                                <List sx={{ minWidth: 400 }}>
                                {registrations.map(r => 
                                     <ListItem
                                     key={r.id}>
                                     <ListItemText primary={r.id}></ListItemText>
                                   </ListItem>
                                )}
                                </List>
                                </Dialog>
                            <TableCell align="left">Дата</TableCell>
                            <TableCell align="left">Время</TableCell>
                            <TableCell align="left">Место</TableCell>
                            <TableCell align="left">Ведущие</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {eventsRows}
                    </TableBody>
                </Table>
            </TableContainer>
            </Paper>
        </>
    );
}
const stubEventRegistrations: EventRegistration[] =
[
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
},
{
    id: 22,
    userId: 222,
    isRegistered: true,
    event:   {
        id: 22,
        name: "Игротехника",
        datetime: new Date(),
        place: "Ломоносова, 9. ауд. 1220",
        organizers: ['Вдовицын М.В.', 'Суязова И.М.']
    },
}
]