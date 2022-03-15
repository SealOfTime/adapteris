import { Spinner } from "components/organisms/spinner";
import React, { createContext, FC, ReactChildren, useContext, useEffect, useState } from "react";
import { useQuery } from "react-query";
import { useNavigate } from "react-router-dom";

export const useAuthenticatedUser = () => {
    const navigate = useNavigate();
    const user = useAuth();
    useEffect(() => (!user.data && !user.isLoading) && navigate("/login"), [user, navigate]);
    return user.data;
}

export const useCanEdit = () => {
    const user = useAuthenticatedUser();
    return user && user.role === 'ADMIN';
}

export const useAuth = () => {
    return useContext(authContext);
}

export const ProvideAuth: FC = ({ children }) => {
    const auth = useProvideAuth();
    return (
        <authContext.Provider value={auth}> {children} </authContext.Provider>
    )
};

const authContext = createContext<
    {
        data: User | null,
        isLoading: boolean
    }>({ data: null, isLoading: false });

type UserRole =
    | 'ADMIN'
    | 'USER'

type User = {
    fullname: string
    shortname: string
    isu: number
    group: string
    vk: string
    phone: string
    tg: string
    email: string
    role: UserRole
}


const useProvideAuth = () => {
    const { isLoading, error, data } = useQuery<User>(['profile', 'my'], {
        retry: 1,
    })
    return { data, isLoading };
}
