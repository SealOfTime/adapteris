import {School, useDupeSchool, useSchool} from "hooks/useSchool";
import React, {FC} from "react";
import {Button, Card, Stack, Typography} from "@mui/material";
import {useDMYDate} from "hooks/useFormatDate";
import {useMySchoolResult, useRegisterForSchool} from "hooks/useMyEventParticipations";
import {useCanEdit} from "hooks/useAuth";
import {useNavigate} from "react-router-dom";

export const SchoolCard: FC<{school: School}> = ({school}) => {
    const currentYear = (new Date()).getFullYear();
    const currentSchool = useSchool((new Date()).getFullYear())
    const isCurrentSchool = currentYear === school.id;

    const canEdit = useCanEdit();

    const myParticipation = useMySchoolResult(school.id)
    const sendRegisterForSchoolReq = useRegisterForSchool();

    const navigate = useNavigate();
    const openSchool = () => {
        navigate(`/school/${school.id}`)
    }
    const register = () => {
        sendRegisterForSchoolReq({schoolId: school.id})
    }

    const sendDupeReq = useDupeSchool()
    const duplicate = () => {
        sendDupeReq({schoolId: school.id})
    }

    const startDate = useDMYDate(school.start)
    const endDate = useDMYDate(school.end)
    const registerStart = useDMYDate(school.registerStart)
    const registerEnd = useDMYDate(school.registerEnd)
    return (
        <Card sx={{
            flexGrow: 0.5,
        }} variant="outlined">
            <Stack sx={{padding: '2rem', height: '100%'}}>
                <Typography variant="h1" textAlign="center" gutterBottom>{school.name}</Typography>
                <Typography fontSize="1.5rem" mb='0.5rem'><b>Начало:</b> {startDate}</Typography>
                <Typography fontSize="1.5rem" mb='0.5rem'><b>Конец:</b> {endDate}</Typography>
                <Typography variant="h2" gutterBottom>Регистрация: </Typography>
                <Typography fontSize="1.25rem" mb='0.5rem'><b>Начало:</b> {registerStart}</Typography>
                <Typography fontSize="1.25rem" mb='0.5rem'><b>Конец:</b> {registerEnd}</Typography>
                {canEdit || myParticipation.data || !isCurrentSchool
                    ? <Button variant="contained" sx={{marginTop: 'auto', marginBottom: '2rem'}} onClick={openSchool}>Перейти</Button>
                    : <Button variant="contained" sx={{marginTop: 'auto'}} onClick={register}>Принять участие</Button>
                }
                {canEdit && !isCurrentSchool && !currentSchool.data &&
                    <Button variant="contained" onClick={duplicate}>Создать школу этого года по образу и подобию</Button>
                }
            </Stack>
        </Card>
    )
}
