import { create } from "zustand";
import { persist } from "zustand/middleware";

interface AuthState {
    token: string | null;

    isAuthenticated: () => boolean;

    login: (token: string) => void;

    logout: () => void;
}

export const useAuthStore = create<AuthState>()(
    persist(
        (set, get) => ({
            token: null,

            isAuthenticated: () => {
                return get().token !== null;
            },

            login: (token) => {
                set({ token });
            },

            logout: () => {
                set({ token: null });
            },
        }),
        {
            name: "auth-storage",
        },
    ),
);