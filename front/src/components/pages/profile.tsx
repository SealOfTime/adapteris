import {FC, useState} from "react";
import {BookmarkIcon, PencilAltIcon, SaveIcon} from "@heroicons/react/outline";
import { Box, Stack, Card, Typography, TextField, Icon, Container, IconButton, LinearProgressClassKey } from "@mui/material";
import { color } from "@mui/system";

export type user = {
    fullName: string
    shortName?: string
    studentId: number
    group: string
    vkLink?: string
    phone: string
    tg?: string
    email: string
}

export const ProfilePage: FC<> = () => {
    const iconStyle={width: '2rem', color: 'black'};
    const user = stubUser;
    const [saved, edit] = useState(true);
    return (
        <>
        <Box pt='1.25rem'>
        <Stack width="100%" pt='1.25rem' paddingLeft='0.5rem' direction="row" justifyContent="space-between">
            <Typography variant="h1" gutterBottom>Профиль
            </Typography>
                {saved==true? <IconButton><PencilAltIcon style={iconStyle} onClick={() => edit(false)} /></IconButton> :
                                <IconButton><BookmarkIcon style={iconStyle} onClick={() => edit(true)} /></IconButton>}
        </Stack>
        </Box>
        <Box pb='64px'>
        <Card variant="outlined">
        <Stack width='70%' pt='1.5rem' pb='1rem' paddingLeft='1rem' direction='column' justifyContent='space-between'> 
            <Typography variant="h3" textAlign='match-parent'>Полное имя</Typography>
            <TextField disabled={saved} variant="standard" property="required" value={user.fullName}></TextField>
            <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Краткое имя</Typography>
            <TextField disabled={saved} variant="standard" property="required" value={user.shortName}></TextField>
            <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>ID в ИСУ</Typography>
            <TextField disabled={saved} variant="standard" property="required" value={user.studentId}></TextField>
            <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Учебная группа</Typography>
            <TextField disabled={saved} variant="standard"  property="required" value={user.group}></TextField>
            <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Ссылка в VK</Typography>
            <TextField disabled={saved} variant="standard" property="required"value={user.vkLink}></TextField>
            <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Номер телефона</Typography>
            <TextField disabled={saved} variant="standard" property="required" value={user.phone}></TextField>
            <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Telegram</Typography>
            <TextField disabled={saved} variant="standard" value={user.tg}></TextField>
            <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Mail</Typography>
            <TextField disabled={saved} variant="standard" property="required" value={user.email}></TextField>
        </Stack>
        </Card>
        </Box>
        </>
    );
}

const stubUser = {
    fullName: "Ekaterina",
    shortName: "GGGGGG",
    studentId: 213784,
    group: "P33102",
    vkLink: 'vk.com',
    phone: "89054758396",
    tg: "@Smaash",
    email: "@gmail.com",
}

