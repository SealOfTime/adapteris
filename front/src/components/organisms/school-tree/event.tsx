import React, {FC} from "react";
import {SchoolEvent} from "hooks/useSchool";
import {IconButton, Stack} from "@mui/material";
import {
    CheckCircleIcon,
    ClockIcon,
    DotsCircleHorizontalIcon,
    ExclamationCircleIcon,
    QuestionMarkCircleIcon, XCircleIcon
} from "@heroicons/react/outline";
import {useNavigate} from "react-router-dom";

export const SchoolEventBubble: FC<{ event: SchoolEvent }> = ({ event }) => {
    const navigate = useNavigate();
    let statusIcon: JSX.Element;
    switch(event.status) {
        case "COMPLETED": statusIcon = <CheckCircleIcon />; break;
        case "PENDING": statusIcon = <QuestionMarkCircleIcon/>; break;
        case "REGISTERED": statusIcon = <ExclamationCircleIcon/>; break;
        case "FAILED": statusIcon=<XCircleIcon/>;break;
        case undefined:
        case "TODO": statusIcon = <DotsCircleHorizontalIcon/>; break;
    }
    return (
        <>
            <Stack style={{ width: 64, height: 64 }}>
                <IconButton onClick={()=>navigate(`/event/${event.id}`)}>
                    {statusIcon}
                </IconButton>
                <span style={{
                    //https://stackoverflow.com/questions/6618648/can-overflow-text-be-centered
                    whiteSpace: 'nowrap',
                    marginLeft: '-100%',
                    marginRight: '-100%',
                    textAlign: 'center',
                }}>{event.name}</span>
            </Stack>
        </>
    )
}
