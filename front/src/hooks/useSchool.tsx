import { MutationFunction, useMutation, useQuery, useQueryClient } from "react-query";

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
export type EventStatus = 'COMPLETED' | 'PENDING' | 'REGISTERED' | 'FAILED' | 'TODO';
export type SchoolEvent = {
    id: number
    name: string
    description: string
    status?: EventStatus,
}

export function useSchool(id: number | 'my') {
    if(id == 'my') {
        id = 1;
    }
    const { data, ...rest } = useQuery<{ school: School }>(['school', id]);
    return { data: data?.school, ...rest }
}

export function useAddStage() {
    const queryClient = useQueryClient();
    const addStage: MutationFunction<Stage, { schoolId: number, stage: Stage }> = async ({ schoolId, stage }) => {
        const resp = await fetch(`/api/school/${schoolId}/events`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(stage),
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
