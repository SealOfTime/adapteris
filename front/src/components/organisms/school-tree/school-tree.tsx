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
import {useCanEdit} from "hooks/useAuth";
import React, {FC, useState} from "react";
import {SchoolStage} from "components/organisms/school-tree/stage";
import {School, useAddStage} from "hooks/useSchool";
import {AddButton} from "components/organisms/school-tree/add-button";
import {useForm} from "react-hook-form";
import {HookFormDatePicker} from "components/organisms/hook-form-date-picker";

export const SchoolTree: FC<{ school: School }> = ({ school }) => {
    const canEdit = useCanEdit();
    return (
        <>
            {school.stages.map(stage => <SchoolStage key={stage.id} stage={stage} />)}
            <Stack direction="row" justifyContent="center">
                {canEdit && <AddStageButton school={school} />}
            </Stack>
        </>
    );
}

const AddStageButton: FC<{ school: School }> = ({ school }) => {
    const [dialogOpen, setDialogOpen] = useState(false);
    const sendAddStageRequest = useAddStage();
    const {register, handleSubmit, control} = useForm({
        defaultValues: {
            name: "Новый блок",
            description: null,
            start: new Date(),
            end: new Date(),
        },
    });
    const addStage = (stage) => {
        sendAddStageRequest({schoolId: school.id, stage: stage}, {
            onSuccess: () => {
                setDialogOpen(false)
            }
        })
    }
    return (
        <>
            <Tooltip title="Добавить блок">
                <AddButton onClick={() => setDialogOpen(true)}/>
            </Tooltip>
            <Dialog open={dialogOpen}>
                <Stack sx={{padding: '2rem'}} spacing="2rem">
                    <Typography variant="h2">Добавить блок</Typography>
                    <form onSubmit={handleSubmit(addStage)}>
                        <Stack spacing="3rem">
                            <TextField label="Название блока" {...register("name")}/>
                            <TextField label="Описание блока" multiline rows={5} {...register("description")}/>
                            <Stack direction="row" spacing="1rem">
                                <HookFormDatePicker label="Начало" name="start" control={control}/>
                                <HookFormDatePicker label="Конец" name="end" control={control}/>
                            </Stack>
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
};
