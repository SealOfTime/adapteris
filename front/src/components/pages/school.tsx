import {Box, Typography} from "@mui/material";
import {Spinner} from "components/organisms/spinner";
import {SchoolEvent, useSchool} from "hooks/useSchool";
import React, {useEffect} from "react";
import {SchoolTree} from "components/organisms/school-tree/school-tree";
import {useNavigate, useParams} from "react-router-dom";
import {useDMYDate} from "hooks/useFormatDate";
import {useMySchoolResult} from "hooks/useMyEventParticipations";
import {useCanEdit} from "hooks/useAuth";

export const SchoolPage = () => {
    const params = useParams();
    const schoolId =  params["schoolId"] ? parseInt(params["schoolId"]) : 0
    const canEdit = useCanEdit();

    const navigate = useNavigate();
    const participation = useMySchoolResult(schoolId)
    useEffect(()=> {
        if(!canEdit && !participation.isLoading && !participation.data) {
            navigate("/")
        }
    }, [participation.isLoading, participation.data, navigate])

    const school = useSchool(schoolId);

    const startDate = useDMYDate(school.data?.start)
    const endDate = useDMYDate(school.data?.end)

    if (school.isLoading) {
        return <Spinner />
    }

    if (!school) {
        return <>Проблема</>;
    }

    return (
        <Box pt='1.25rem'>
            <Typography variant="h1" gutterBottom>План Школы {schoolId} года:</Typography>
            <Typography fontSize="1.5rem" textAlign="center" gutterBottom>{startDate} - {endDate}</Typography>
            <SchoolTree school={school.data} />
        </Box>
    )
}
