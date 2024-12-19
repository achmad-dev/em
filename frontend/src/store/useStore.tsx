/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

interface UserState {
    token: string;
    role?: string;
}

interface UserActions {
    setToken: (token: string) => void;
    setRole: (role: string) => void;
    reset: () => void;
}

const useStore = create<UserState & UserActions>()(
    devtools(
        persist(
            (set) => ({
                token: "",
                role: "",
                setToken: (token) => set({ token }),
                setRole: (role) => set({ role }),
                reset: () => set({ token: "", role: "" }),
            }),
            {
                name: "user-storage", // unique name
            }
        )
    )
);

export default useStore;