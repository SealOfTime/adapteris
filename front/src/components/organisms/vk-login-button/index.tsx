import { Button } from '@mui/material';
import React, { useRef, useState } from 'react';
import { VkLogo } from './vk-logo';

export const VkLoginButton = () => {
    const [popup, setPopup] = useState<Window | null>(null);
    const initiateLogin = () => {
        if (!popup) {
            const newPopup = popupCenter({ url: "/api/auth/login", title: "Окно авторизации", w: 660, h: 360 })
            newPopup.onmessage = (e) => {
                console.log(e);
                if (e === "login success") {
                    newPopup.close();
                    setPopup(null);
                }
            };
            const timer = setInterval(() => {
                if (newPopup.closed) {
                    clearInterval(timer)
                    setPopup(null);
                }
            }, 1000)
            setPopup(newPopup);
        } else {
            popup.focus();
        }
    };
    return (
        <>
            <Button variant="contained" style={{
                padding: 0,
                margin: 'auto',
                width: 120,
                display: 'flex',
                justifyContent: 'space-around',
            }} onClick={initiateLogin}><VkLogo /><div>Войти</div></Button>
        </>
    )
}

function popupCenter({ url, title, w, h }) {
    // Fixes dual-screen position                             Most browsers      Firefox
    const dualScreenLeft = window.screenLeft !== undefined ? window.screenLeft : window.screenX;
    const dualScreenTop = window.screenTop !== undefined ? window.screenTop : window.screenY;

    const width = window.innerWidth ? window.innerWidth : document.documentElement.clientWidth ? document.documentElement.clientWidth : screen.width;
    const height = window.innerHeight ? window.innerHeight : document.documentElement.clientHeight ? document.documentElement.clientHeight : screen.height;

    const systemZoom = width / window.screen.availWidth;
    const left = (width - w) / 2 / systemZoom + dualScreenLeft
    const top = (height - h) / 2 / systemZoom + dualScreenTop
    const newWindow = window.open(url, title, `
      popup,
      width=${w / systemZoom}, 
      height=${h / systemZoom}, 
      top=${top}, 
      left=${left}`
    )

    if (window.focus) newWindow.focus();
    return newWindow;
}