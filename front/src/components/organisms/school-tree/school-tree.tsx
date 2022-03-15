import {PlusCircleIcon} from "@heroicons/react/outline";
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    IconButton,
    Stack,
    Tooltip,
    Typography
} from "@mui/material";
import {useAuthenticatedUser, useCanEdit} from "hooks/useAuth";
import React, {FC, useState} from "react";
import {SchoolStage} from "components/organisms/school-tree/stage";
import {School, useAddStage} from "hooks/useSchool";
import {AddButton} from "components/organisms/school-tree/add-button";

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
    return (
        <>
            <Tooltip title="Добавить блок">
                <AddButton onClick={() => setDialogOpen(true)}/>
            </Tooltip>
            <Dialog open={dialogOpen}>
                <DialogTitle><Typography variant="h2">Добавить блок</Typography></DialogTitle>
                <DialogContent>
                    <form>

                    </form>
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => setDialogOpen(false)}>Сохранить</Button>
                    <Button onClick={() => setDialogOpen(false)}>Закрыть</Button>
                </DialogActions>
            </Dialog>
        </>
    )
};
