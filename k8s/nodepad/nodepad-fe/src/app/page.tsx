import NodepadApp from "@/components/nodepad-app";
import type { Metadata } from "next";

export const metadata: Metadata = {
    alternates: {
        canonical: "/",
        languages: {
            "vi-VN": "/",
        },
    },
};

export default function Home() {
    return <NodepadApp />;
}
