import React, {FC, useState} from "react";
import {Step, useAddEvent, useAddStage} from "hooks/useSchool";
import {
    Box,
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    Stack, TextField,
    Tooltip,
    Typography
} from "@mui/material";
import {SchoolEventBubble} from "components/organisms/school-tree/event";
import {useCanEdit} from "hooks/useAuth";
import {AddButton} from "components/organisms/school-tree/add-button";
import {useForm} from "react-hook-form";
import {HookFormDatePicker} from "components/organisms/hook-form-date-picker";

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
    const [dialogOpen, setDialogOpen] = useState(false);
    const sendAddStageRequest = useAddEvent();
    const {register, handleSubmit, control} = useForm({
        defaultValues: {
            name: "Новое событие",
            description: null,
        },
    });
    const addStage = (event) => {
        sendAddStageRequest({stepId: step.id, event: event}, {
            onSuccess: () => {
                setDialogOpen(false)
            }
        })
    }
    return (
        <>
            <Tooltip title="Добавить событие">
                <AddButton onClick={() => setDialogOpen(true)}/>
            </Tooltip>
            <Dialog open={dialogOpen}>
                <Stack sx={{padding: '2rem'}} spacing="2rem">
                    <Typography variant="h2">Добавить событие</Typography>
                    <form onSubmit={handleSubmit(addStage)}>
                        <Stack spacing="3rem">
                            <TextField label="Название события" {...register("name")}/>
                            <TextField label="Описание события" multiline rows={5} {...register("description")}/>
                            <Stack direction="row" mt='auto' justifyContent="space-around">
                                <Button type="submit">Сохранить</Button>
                                <Button onClick={() => setDialogOpen(false)}>Закрыть</Button>
                            </Stack>
                        </Stack>
                    </form>
                </Stack>
            </Dialog>
        </>
    )
}
