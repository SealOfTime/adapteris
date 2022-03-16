import React from 'react';
import { Button, IconButton, Stack, TextField } from "@mui/material"
import { VkLoginButton } from 'components/organisms/vk-login-button';

export const LoginPage = () => {
    return (
        <>
            <Stack sx={{height: '100%'}}justifyContent="center">
                <VkLoginButton />
            </Stack>
        </>
    )
}
