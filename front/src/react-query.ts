import { QueryClient } from "react-query";

const queryFn = async ({ queryKey }) => {
    const res = await fetch(`/api/${queryKey.join('/')}`);
    if(res.status === 404) {
        return null;
    }

    const data = await res.json();
    return data;
};

export const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            queryFn: queryFn,
        }
    }
}
);
