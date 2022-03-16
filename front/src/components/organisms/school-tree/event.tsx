import React, {FC, useMemo} from "react";
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
import {useMyEventParticipations} from "hooks/useMyEventParticipations";

export const SchoolEventBubble: FC<{ event: SchoolEvent }> = ({ event }) => {
    const navigate = useNavigate();
    const participations = useMyEventParticipations();
    const status = useMemo(()=>{
        const now = new Date();
        if(participations.data) {
            for (const p of participations.data) {
                if (p.session.event.id === event.id) {
                    if (p.passed) {
                        return "COMPLETED"
                    }
                    if (p.passed === null) {
                        if (p.session.datetime < now) {
                            return "PENDING"
                        }
                        return "REGISTERED"
                    }
                    return "FAILED"
                }
            }
        }
        return "TODO"
    }, [participations.data])
    let statusIcon: JSX.Element;
    switch(status) {
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
