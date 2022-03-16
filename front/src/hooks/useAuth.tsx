import { Spinner } from "components/organisms/spinner";
import React, {createContext, FC, ReactChildren, useCallback, useContext, useEffect, useState} from "react";
import { useQuery } from "react-query";
import {createSearchParams, useLocation, useNavigate} from "react-router-dom";

type UserRole =
    | 'ADMIN'
    | 'USER'

export type User = {
    id: number
    role: UserRole
    shortname: string
    vk: string
    email: string
    fullname?: string
    phone?: string
    tg?: string
}
export const useLogin = () => {
    const {pathname} = useLocation();
    const navigate = useNavigate();
    const login = useCallback(
        () => {
            navigate({
                pathname: "/login",
                search: createSearchParams({
                    redirect: pathname,
                }).toString(),
            })
        },
        [navigate, pathname],
    );
    return login
}

export const useAuthenticatedUser = () => {
    const login = useLogin();
    const user = useUser();
    useEffect(() => {
        if (!user.data && !user.isLoading) {
            login();
        }
    }, [user, login]);
    return user.data;
}

export const useCanEdit = () => {
    const {data: user} = useUser();
    return user && user.role === 'ADMIN';
}

export const useUser = () => {
    const { isLoading, error, data } = useQuery<User | null>(['profile', 'my'], {
        queryFn: async ({queryKey}) => {
            const resp = await fetch(`/api/${queryKey.join('/')}`)
            if(resp.status === 401) {
                return null;
            }
            if(resp.status === 200) {
                return resp.json();
            }
            throw `Unexpected error: ${await resp.text()}`
        },
        retry: 1,
    })
    return { data, isLoading };
}
