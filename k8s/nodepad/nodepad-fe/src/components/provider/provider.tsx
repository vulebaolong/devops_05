"use client";
import "@mantine/core/styles.css";

import { useHotkeys } from "@mantine/hooks";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import dayjs from "dayjs";
import customParseFormat from "dayjs/plugin/customParseFormat";
import duration from "dayjs/plugin/duration";
import relativeTime from "dayjs/plugin/relativeTime";
import timezone from "dayjs/plugin/timezone";
import utc from "dayjs/plugin/utc";
import { ReactNode } from "react";
import { Toaster } from "sonner";
import MantineProvider from "./mantine/mantine.provider";

dayjs.extend(relativeTime);
dayjs.extend(duration);
dayjs.extend(utc);
dayjs.extend(timezone);
dayjs.extend(customParseFormat);

const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            refetchOnWindowFocus: false,
            retry: false,
            gcTime: 5 * 60 * 1000,
            staleTime: 30 * 1000,
        },
    },
});

export default function Provider({ children }: { children: ReactNode }) {
    useHotkeys([["mod+.", () => process.env.NEXT_PUBLIC_IS_PRODUCTION !== `true` && window.open("/test", "_blank")]]);

    return (
        <QueryClientProvider client={queryClient}>
            <MantineProvider>
                {children}
                <Toaster />
            </MantineProvider>
        </QueryClientProvider>
    );
}
