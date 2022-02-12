import React, { createContext, FC, ReactChildren, useContext, useState } from "react";
import { useQuery } from "react-query";

export const useAuth = () => {
    return useContext(authContext);
}

export const ProvideAuth: FC = ({ children }) => {
    const auth = useProvideAuth();
    return (
        <authContext.Provider value={auth}> {children} </authContext.Provider>
    )
};

const authContext = createContext({});

type Authentication = {
    signIn: () => void
    signOut: () => void
}

type UserRole =
    | 'ADMIN'
    | 'USER'

type User = {
    roles: UserRole
}


const useProvideAuth = () => {
    const user;
    return {};
}