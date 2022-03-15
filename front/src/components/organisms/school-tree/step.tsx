import React, {FC, useState} from "react";
import {Step} from "hooks/useSchool";
import {
    Box,
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    Stack,
    Tooltip,
    Typography
} from "@mui/material";
import {SchoolEventBubble} from "components/organisms/school-tree/event";
import {useCanEdit} from "hooks/useAuth";
import {AddButton} from "components/organisms/school-tree/add-button";

export const SchoolStep: FC<{ step: Step }> = ({ step }) => {
    const canEdit = useCanEdit();
    const text = step.mustComplete == 0
        ? "Поучаствуй во всех мероприятиях:"
        : `Поучаствуй в ${step.mustComplete} из ${step.events.length} мероприятий:`
    return (
        <>
            <Box sx={{
                boxSizing: 'border-box',
                marginBottom: '2rem',
                paddingTop: '0.5rem',
                paddingBottom: '1.5rem',
                '&:hover': {
                    backgroundColor: 'grey.100',
                },
            }}>
                <Typography variant="h6" align="center" mt="10px">{text}</Typography>
                <Stack direction="row" justifyContent="space-evenly">
                    {step.events.map(e => <SchoolEventBubble key={e.id} event={e} />)}
                    {canEdit && <AddEventButton step={step}/>}
                </Stack>
            </Box>
        </>
    )
}

const AddEventButton: FC<{step: Step}> = ({step})=>{
    const [dialog, setDialog] = useState<boolean>(false);
    return (
        <>
            <Tooltip title="Добавить мероприятие"><AddButton onClick={()=>setDialog(true)}/></Tooltip>
            <Dialog open={dialog}>
                <DialogTitle><Typography variant="h2">Добавить мероприятие</Typography></DialogTitle>
                <DialogContent></DialogContent>
                <DialogActions>
                    <Button onClick={()=>setDialog(false)}>Сохранить</Button>
                    <Button onClick={()=>setDialog(false)}>Закрыть</Button>
                </DialogActions>
            </Dialog>
        </>
    )
}
