
import React, { FC, useEffect, useState } from "react";
import { BookmarkIcon, PencilAltIcon, SaveIcon } from "@heroicons/react/outline";
import { Box, Stack, Card, Typography, TextField, Icon, Container, IconButton, LinearProgressClassKey } from "@mui/material";
import { color } from "@mui/system";
import { useAuthCtx, useAuthenticatedUser } from "hooks/useAuth";
import { Spinner } from "components/organisms/spinner";
import { useNavigate } from "react-router-dom";

export type User = {
    fullName: string
    shortName?: string
    studentId: number
    group: string
    vkLink?: string
    phone: string
    tg?: string
    email: string
}

const iconStyle = { width: '2rem', color: 'black' };

export const ProfilePage: FC = () => {
    const navigate = useNavigate();
    const user = useAuthenticatedUser();
    const [edited, setEdited] = useState(false);

    if (!user) {
        return <Spinner />;
    }
    const editUser = () => setEdited(true);
    const saveUser = () => console.log();

    const actionButton = edited
        ? <IconButton onClick={saveUser}><BookmarkIcon style={iconStyle} /></IconButton>
        : <IconButton onClick={editUser}><PencilAltIcon style={iconStyle} /></IconButton >
    return (
        <>
            <Stack width="100%" pt='2.5rem' paddingLeft='0.5rem' direction="row" justifyContent="space-between">
                <Typography variant="h1" gutterBottom>Профиль</Typography>
                {actionButton}
            </Stack>
            <Box>
                <Card variant="outlined">
                    <Stack width='70%' pt='1.5rem' pb='1rem' paddingLeft='1rem' direction='column' justifyContent='space-between'>
                        <UserField label="Полное имя" value={user.fullname} edited={edited} />
                        <UserField label="Краткое имя" value={user.shortname} edited={edited} />
                        <UserField label="VK" value={user.vk} edited={edited} />
                        <UserField label="Telegram" value={user.tg} edited={edited} />
                        <UserField label="Email" value={user.email} edited={edited} />
                        <UserField label="Номер телефона" value={user.phone} edited={edited} />
                    </Stack>
                </Card>
            </Box>
        </>
    );
}

const UserField = ({ label, value, edited, onEdit }) => (
    <>
        <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>{label}</Typography>
        <TextField disabled={!edited} variant="standard" property="required" value={value} onChange={onEdit}></TextField>
    </>
)