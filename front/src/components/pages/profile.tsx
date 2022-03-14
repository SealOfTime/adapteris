import React, { FC } from "react";
import { PencilAltIcon } from "@heroicons/react/outline";
import { Box, Stack, Card, Typography, TextField, Icon, Container, IconButton, LinearProgressClassKey } from "@mui/material";
import { color } from "@mui/system";

export type User = {
    FullName: string;
    ShortName: string;
    StudentId: bigint;
    group: string;
    VkLink: string;
    phone: string;
    tg: string;
    email: string;
}

export const ProfilePage: FC = (User) => {
    const iconStyle = { width: '2rem', color: 'black' };
    return (
        <>
            <Box pt='1.25rem'>
                <Stack width="100%" pt='1.25rem' paddingLeft='0.5rem' direction="row" justifyContent="space-between">
                    <Typography variant="h1" gutterBottom>Профиль
                    </Typography>
                    <IconButton><PencilAltIcon style={iconStyle} /></IconButton>
                </Stack>
            </Box>
            <Box pb='64px'>
                <Card variant="outlined">
                    <Stack width='70%' pt='1.5rem' pb='1rem' paddingLeft='1rem' direction='column' justifyContent='space-between'>
                        <Typography variant="h3" textAlign='match-parent'>Полное имя</Typography>
                        <TextField variant="standard" property="required"></TextField>
                        <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Краткое имя</Typography>
                        <TextField variant="standard" property="required"></TextField>
                        <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>ID в ИСУ</Typography>
                        <TextField variant="standard" property="required"></TextField>
                        <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Учебная группа</Typography>
                        <TextField variant="standard" property="required"></TextField>
                        <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Ссылка в VK</Typography>
                        <TextField variant="standard" property="required"></TextField>
                        <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Номер телефона</Typography>
                        <TextField variant="standard" property="required"></TextField>
                        <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Telegram</Typography>
                        <TextField variant="standard"></TextField>
                        <Typography variant="h3" textAlign='match-parent' paddingTop='1.5rem'>Mail</Typography>
                        <TextField variant="standard" property="required"></TextField>
                    </Stack>
                </Card>
            </Box>
        </>
    );
}
