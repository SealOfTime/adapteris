import { useMemo } from "react";

export const useDMYDate = (date: Date) => {
    const {day, month, year} = useFormatDate(date);
    return `${day} ${month} ${year}`
}

export function formatDate(date:Date|string) {
    if(typeof(date) === "string") {
        date = new Date(date);
    }
    const day = date?.getDate();
    const month = monthName[date?.getMonth()];
    const year = date?.getFullYear();
    const hour = date?.getHours();
    const minute = date?.getMinutes();
    return {
        day,
        month,
        year,
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
