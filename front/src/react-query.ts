import { QueryClient } from "react-query";

const queryFn = async ({ queryKey }) => {
    const raw = await fetch(`/api/${queryKey.join('/')}`);
    const data = await raw.json();
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