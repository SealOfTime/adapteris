import { useMemo } from "react";

export function formatDate(date:Date) {
    const day = date.getDate();
    const month = monthName[date.getMonth()];
    const hour = date.getHours();
    const minute = date.getMinutes();
    return {
        day,
        month,
        hour: hour < 10 ? `0${hour}` : hour,
        minute: minute < 10 ? `0${minute}` : minute
    }
}

const monthName = [
    'января',
    'февраля',
    'марта',
    'апреля',
    'мая',
    'июня',
    'июля',
    'сентября',
    'октября',
    'ноября',
    'декабря'
];

export function useFormatDate(date: Date) {
    return useMemo(() => formatDate(date), [date]);
}