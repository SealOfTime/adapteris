import { Button } from '@mui/material';
import React, {useEffect, useRef, useState} from 'react';
import { VkLogo } from './vk-logo';
import {useNavigate, useSearchParams} from "react-router-dom";
import {useUser} from "hooks/useAuth";

export const VkLoginButton = () => {
    const [popup, setPopup] = useState<Window | null>(null);
    const user = useUser()
    const navigate = useNavigate();
    const [params] = useSearchParams({redirect: "/"});
    useEffect(()=> {
        if(user.data) {
            navigate(params.get("redirect"))
        }
    }, [user])
    useEffect(
        () => {
            if(popup) {
                const timer = setInterval(()=>{
                    if(popup?.closed) {
                        clearInterval(timer);
                        setPopup(null);
                    }
                }, 1000);
                return ()=>clearInterval(timer)
            }},
        [popup, setPopup]
    )
    const initiateLogin = () => {
        if (!popup) {
            const newPopup = popupCenter({ url: "/api/auth/login", title: "Окно авторизации", w: 660, h: 360 })
            newPopup.onmessage = (e) => {
                if (e.data === "success login") {
                    newPopup.close();
                    navigate(params.get("redirect"))
                }
            };
            setPopup(newPopup);
        } else {
            popup.focus();
        }
    };
    return (
        <>
            <Button variant="contained" sx={{
                padding: '0.5rem',
                margin: 'auto',
                display: 'flex',
                fontWeight: 'bold',
                justifyContent: 'space-around',
                color: 'black',
                backgroundColor: 'white',
                borderRadius: '20px',
                '&:hover': {
                    color: 'white',
                    backgroundColor: '#0077FF',
                }
            }} onClick={initiateLogin}><VkLogo/><div style={{margin: '1rem'}}>Войти Вконтакте</div></Button>
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
