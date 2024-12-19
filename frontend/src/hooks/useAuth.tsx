/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/
import instance from "./base";
import { useState } from "react";
import useStore from "@/store/useStore";
import { AuthResponse } from "@/type/type";
import { useToast } from "./use-toast";
import { v4 as uuidv4 } from 'uuid';
import  GenerateSignature  from "@/utils/signature";

export const useAuth = () => {
    const { setToken, setRole } = useStore();
    const { toast } = useToast();

    const [state, setState] = useState<{ loading: boolean; error: boolean; data: AuthResponse | null }>({
        loading: true,
        error: false,
        data: null,
    });

    const fetchAuth = (username: string, password: string) => {

        const timestamp = new Date().toISOString();
        const key = import.meta.env.VITE_SECRET_KEY_HMAC || "key";
        const signature = GenerateSignature(key, timestamp);

        instance
            .post("/signin", { username, password }, {
                headers: {
                    "X-RequestId": uuidv4(),
                    "X-Signature": signature,
                    "X-TimeStamp": timestamp,
                },
             })
            .then((response) => {
                const authResponse: AuthResponse = response.data.data;
                setToken(authResponse.token);
                setRole(authResponse.role);
                setState({
                    loading: false,
                    error: false,
                    data: authResponse,
                });
                toast({ description: "Login success" });
            })
            .catch((e) => {
                setState({
                    loading: false,
                    error: true,
                    data: null,
                });
                console.log(e);
                toast({ description: "Login failed" });
            });
    }

    return {
        fetchAuth,
        loading: state.loading,
        error: state.error,
        data: state.data,
    };
};