import { CheckCircleIcon, PlusCircleIcon } from "@heroicons/react/outline";
import { Box, Button, Dialog, DialogActions, DialogContent, DialogTitle, IconButton, Stack, Tooltip, Typography } from "@mui/material";
import { Spinner } from "components/organisms/spinner";
import { useAuthenticatedUser } from "hooks/useAuth";
import { School, SchoolEvent, Stage, Step, useAddStage, useSchool } from "hooks/useSchool";
import React, { FC, useState } from "react";
import { useQuery } from "react-query";
import {SchoolTree} from "components/organisms/school-tree/school-tree";
import {useParams} from "react-router-dom";

export const SchoolPage = () => {
    const {schoolId} = useParams();
    const { data: school, isLoading } = useSchool(schoolId ? parseInt(schoolId) : "my");
    if (isLoading) {
        return <Spinner />
    }

    if (!school) {
        return <>Проблема</>;
    }

    return (
        <Box pt='1.25rem'>
            <Typography variant="h1" gutterBottom>План школы {schoolId}:</Typography>
            <SchoolTree school={school} />
        </Box>
    )
}
