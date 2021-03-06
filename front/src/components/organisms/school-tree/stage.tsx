import React, {FC, useState} from "react";
import {Stage, useAddStep} from "hooks/useSchool";
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
import {SchoolStep} from "components/organisms/school-tree/step";
import {AddButton} from "components/organisms/school-tree/add-button";
import {useCanEdit} from "hooks/useAuth";
import {useDMYDate} from "hooks/useFormatDate";

export const SchoolStage: FC<{ stage: Stage }> = ({stage}) => {
    const canEdit = useCanEdit();
    const startDate = useDMYDate(stage.start)
    const endDate = useDMYDate(stage.end)
    return (
        <Box sx={{
            borderBottom: t => `1px solid ${t.palette.grey.A200}`,
        }}>
            <Typography variant="h2" gutterBottom>{stage.name}</Typography>
            <Typography fontSize="1rem" gutterBottom>{startDate} - {endDate}</Typography>
            {stage.steps.map(step => <SchoolStep key={step.id} step={step}/>)}
            <Stack direction="row" justifyContent="center">
                {canEdit && <AddStepButton stage={stage}/>}
            </Stack>
        </Box>
    )
}

const AddStepButton: FC<{ stage: Stage }> = ({stage}) => {
    const addStep = useAddStep();
    return (
        <>
            <Tooltip title="Добавить шаг">
                <AddButton onClick={() => addStep({stageId: stage.id})}/>
            </Tooltip>
        </>
    )
}
