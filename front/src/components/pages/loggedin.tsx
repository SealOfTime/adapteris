import {Typography} from "@mui/material";
import {useEffect} from "react";

export const LoggedInPage = () => {
    useEffect(
        () => {
            window.postMessage("success login", window.opener);
        }
    )
    return (
        <Typography variant="h3">
            Через минуту вы отправитесь назад
        </Typography>
    );
}
