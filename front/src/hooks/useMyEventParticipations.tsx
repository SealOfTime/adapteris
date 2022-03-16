import {useMutation, useQuery, useQueryClient, UseQueryResult} from "react-query";
import {SchoolEvent} from "hooks/useEvent";
import {useAuthenticatedUser, useLogin} from "hooks/useAuth";
import {createSearchParams} from "react-router-dom";

export type ParticipationEventSession = {
    id: number
    place: string
    datetime: Date,
    event: SchoolEvent,
}

export type EventParticipation = {
    id: number
    session: ParticipationEventSession
    passed: boolean | null
}

export type ParticipationSchool = {
    id: number
    name: string
}

export type SchoolParticipation = {
    id: number
    result: 'ACCEPT' | 'DECLINE' | ''
    school: ParticipationSchool
}

export function useMySchoolResult(schoolId?: number) {
    const {data, ...rest} = useQuery<{result: SchoolParticipation}>(['participations', 'school', 'my', schoolId], {
        enabled: schoolId != undefined,
    });
    return {data: data?.result, ...rest};
}

export function useRegisterForSchool() {
    const queryClient = useQueryClient();
    const login = useLogin();
    const register = async ({schoolId}) => {
        const resp = await fetch(`/api/participations/school/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({schoolId})
        })
        if(resp.status === 401) {
            login();
        }
        if(resp.status !== 201) {
            throw "nope";
        }
    }
    const mutation = useMutation(register, {
        onSuccess: () => {
            queryClient.invalidateQueries(['participations', 'school'])
        }
    })
    return mutation.mutate
}

type MyEventParticipationsOptions = {
    schoolId?: number
    eventId?: number
}
export function useMyEventParticipations(opts?: MyEventParticipationsOptions): UseQueryResult<EventParticipation[]> {
    const user = useAuthenticatedUser()
    const participations = useQuery(
        ['participations', 'event', 'my', opts?.schoolId, opts?.eventId], {
        queryFn: async ({queryKey})=>{
            const searchParams = createSearchParams({
                ...(queryKey[3] && {schoolId: queryKey[3] as string}),
                ...(queryKey[4] && {eventId: queryKey[4] as string}),
            }).toString()
            const res = await fetch(`/api/participations/event/my?${searchParams}`, {
            });
            if(res.status === 404) {
                return null;
            }
            if(res.status !== 200) {
                throw "couldn't find my participations";
            }
            const response = await res.json() as {participations: EventParticipation[]}
            return response.participations;
        },
        enabled: user?.role != 'ADMIN',
    });
    return participations;
}

export function useRegisterForSession() {
    const queryClient = useQueryClient()
    const register = async ({sessionId}) => {
        const resp = await fetch(`/api/participations/event/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({sessionId: sessionId})
        })
        if(resp.status !== 200) {
            throw "nope";
        }
    }
    const mutation = useMutation(register, {
        onSuccess: () => {
            queryClient.invalidateQueries(['participations', 'event'])
        }
    })
    return mutation.mutate
}

export function useCancelSessionParticipation() {
    const queryClient = useQueryClient()
    const register = async ({participantId}) => {
        const resp = await fetch(`/api/participations/event/${participantId}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
        })
        if(resp.status !== 200) {
            throw "nope";
        }
    }
    const mutation = useMutation(register, {
        onSuccess: () => {
            queryClient.invalidateQueries(['participations', 'event'])
        }
    })
    return mutation.mutate
}
