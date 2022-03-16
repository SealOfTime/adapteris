import { MutationFunction, useMutation, useQuery, useQueryClient } from "react-query";
import {SchoolEvent} from "hooks/useEvent";

export type School = {
    id: number
    name: string
    visible: boolean
    start: Date
    end: Date
    registerStart: Date
    registerEnd: Date
    stages: Stage[]
};

export type Stage = {
    id: number
    name: string
    description: string
    start: Date
    end: Date
    steps: Step[]
};

export type Step = {
    id: number
    mustComplete: number
    events: SchoolEvent[]
}

export function useSchool(id: number) {
    const { data, ...rest } = useQuery<{ school: School }>(['school', id]);
    return { data: data?.school, ...rest }
}
export function useDupeSchool() {
    const queryClient = useQueryClient();
    const dupeSchool: MutationFunction<void, {schoolId: number}> = async ({schoolId}) => {
        const resp = await fetch(`/api/school/${schoolId}/copy`, {
            method: 'POST',
        })
        if(resp.status !== 201) {
            throw "couldn't copy school";
        }
    }
    const mutation = useMutation(dupeSchool, {
        onSuccess: () => queryClient.invalidateQueries(['school', (new Date()).getFullYear()])
    })
    return mutation.mutate
}

export function useAddSchool() {
    const queryClient = useQueryClient();
    const addSchool: MutationFunction<void, {school: Partial<School>}> = async ({school}) => {
        const resp = await fetch(`/api/school`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(school),
        })
        if(resp.status !== 200){
            throw "couldn't create school";
        }
    }
    const mutation = useMutation(addSchool, {
        onSuccess: () => {
            queryClient.invalidateQueries(['school'])
        }
    })
    return mutation.mutate
}
export function useAddStage() {
    const queryClient = useQueryClient();
    const addStage: MutationFunction<Stage, { schoolId: number, stage: Stage }> = async ({ schoolId, stage }) => {
        const resp = await fetch(`/api/school/${schoolId}/stages`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({stage}),
        })
        return await resp.json();
    }
    const mutation = useMutation(addStage, {
        onSuccess: () => {
            queryClient.invalidateQueries(['school'])
        }
    });
    return mutation.mutate;
}

export function useAddStep() {
    const queryClient = useQueryClient();
    const addStep: MutationFunction<Step, {stageId: number}> = async ({stageId}) => {
        const resp = await fetch(`/api/stage/${stageId}/steps`, {
            method: 'POST',
        })
        return await resp.json();
    }
    const mutation = useMutation(addStep, {
        onSuccess: () => {
            queryClient.invalidateQueries(['school'])
        },
    })
    return mutation.mutate;
}

export function useAddEvent() {
    const queryClient = useQueryClient();
    const addEvent: MutationFunction<Event, {stepId: number, event: Partial<SchoolEvent>}> = async ({stepId, event}) => {
        const resp = await fetch(`/api/step/${stepId}/events`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(event),
        })
        if(resp.status != 201) {
            throw "couldn't create event";
        }

        return await resp.json();
    }
    const mutation = useMutation(addEvent, {
        onSuccess: () => {
            queryClient.invalidateQueries(['school'])
        },
    })
    return mutation.mutate;
}
