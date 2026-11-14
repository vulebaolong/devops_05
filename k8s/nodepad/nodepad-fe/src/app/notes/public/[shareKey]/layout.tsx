import type { Metadata } from "next";

type PublicNote = {
    title: string | null;
    content: string;
};

type LayoutProps = {
    children: React.ReactNode;
    params: Promise<{ shareKey: string }>;
};

function excerpt(content: string) {
    const normalized = content.replace(/\s+/g, " ").trim();
    return normalized.length > 150 ? `${normalized.slice(0, 147)}...` : normalized;
}

async function getPublicNote(shareKey: string): Promise<PublicNote | null> {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8888/api";
    try {
        const response = await fetch(`${apiUrl}/notes/public/${encodeURIComponent(shareKey)}`, {
            next: { revalidate: 60 },
        });
        if (!response.ok) return null;
        return response.json();
    } catch {
        return null;
    }
}

export async function generateMetadata({ params }: Omit<LayoutProps, "children">): Promise<Metadata> {
    const { shareKey } = await params;
    const note = await getPublicNote(shareKey);
    const title = note?.title?.trim() || "Ghi chú được chia sẻ";
    const description = note?.content ? excerpt(note.content) : "Đọc ghi chú được chia sẻ riêng tư qua Nodepad.";
    const url = `/notes/public/${encodeURIComponent(shareKey)}`;

    return {
        title,
        description,
        alternates: { canonical: url },
        openGraph: {
            type: "article",
            locale: "vi_VN",
            siteName: "Nodepad",
            url,
            title: `${title} | Nodepad`,
            description,
            images: [{
                url: "/opengraph-image",
                width: 1200,
                height: 630,
                alt: "Nodepad — Ghi chú online nhanh, riêng tư và miễn phí",
            }],
        },
        twitter: {
            card: "summary_large_image",
            title: `${title} | Nodepad`,
            description,
            images: ["/twitter-image"],
        },
        robots: {
            index: false,
            follow: false,
        },
    };
}

export default function PublicNoteLayout({ children }: LayoutProps) {
    return children;
}
