import {useMutation, useQuery, useQueryClient} from "react-query";
import {User} from "hooks/useAuth";

export type EventStatus = 'COMPLETED' | 'PENDING' | 'REGISTERED' | 'FAILED' | 'TODO';
export type SchoolEvent = {
    id: number
    name: string
    description: string
    sessions: EventSession[]
    status?: EventStatus,
}

export type EventSession = {
    id: number
    place: string
    datetime: Date
}

export function useEvent(eventId: number) {
    return useQuery<SchoolEvent>(['event', eventId])
}

export function useAddSession() {
    const qc = useQueryClient();
    const addSession = async ({eventId, session}) => {
        const res = await fetch(`/api/event/${eventId}/sessions`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({session}),
        })
        if (res.status !== 201) {
            throw "nope";
        }
    }
    const mutation = useMutation(addSession, {
        onSuccess: () => {
            qc.invalidateQueries(['event'])
        }
    })
    return mutation.mutate
}
