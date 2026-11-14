export const API_ENDPOINTS = {
    auth: {
        me: "/auth/me",
        login: "/auth/login",
        register: "/auth/register",
        logout: "/auth/logout",
    },
    notes: {
        root: "/notes",
        detail: (id: string) => `/notes/${encodeURIComponent(id)}`,
        public: (shareKey: string) => `/notes/public/${encodeURIComponent(shareKey)}`,
    },
} as const;
