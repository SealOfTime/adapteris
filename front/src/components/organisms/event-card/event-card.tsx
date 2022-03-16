import React, { FC, useMemo, useState } from "react";
import { Card, Stack, CardContent, Typography, Button, IconButton } from "@mui/material";
import { ClockIcon, OfficeBuildingIcon, UsersIcon, ArrowCircleRightIcon } from "@heroicons/react/outline";
import { EventPage } from "../../pages/event-page";
import {Link, useParams} from 'react-router-dom'
import Dialog from '@mui/material/Dialog';
import {useEvent} from "hooks/useEvent";

export const EventCard: FC = () => {
    const {eventId} = useParams();
    const event = useEvent(parseInt(eventId));
    return (
        <>
        </>
    )
}
