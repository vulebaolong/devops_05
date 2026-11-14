import Provider from "@/components/provider/provider";
import { ColorSchemeScript } from "@mantine/core";
import type { Metadata, Viewport } from "next";
import { Comfortaa } from "next/font/google";

const comfortaa = Comfortaa({
    subsets: ["latin", "vietnamese", "cyrillic", "cyrillic-ext", "greek", "latin-ext"],
    variable: "--font-comfortaa",
    display: "swap",
});

export const metadata: Metadata = {
    metadataBase: new URL("https://nodepad.vulebaolong.com"),

    title: {
        default: "Nodepad — Ghi chú online nhanh, riêng tư và miễn phí",
        template: "%s | Nodepad",
    },

    description:
        "Viết ghi chú ngay không cần đăng nhập, nhận liên kết public để chia sẻ và đồng bộ mọi note khi có tài khoản.",

    applicationName: "Nodepad",
    category: "productivity",
    referrer: "origin-when-cross-origin",
    formatDetection: {
        email: false,
        address: false,
        telephone: false,
    },

    keywords: [
        "Nodepad",
        "ghi chú online",
        "sổ tay online",
        "notepad online",
        "ghi chú miễn phí",
        "chia sẻ ghi chú",
        "viết note online",
        "lưu ghi chú",
        "ứng dụng ghi chú",
    ],

    authors: [{ name: "Nodepad", url: "https://nodepad.vulebaolong.com" }],
    creator: "Nodepad",
    publisher: "Nodepad",
    manifest: "/site.webmanifest",

    icons: {
        icon: [
            {
                url: "/favicon.ico",
            },
            {
                url: "/favicon-16x16.png",
                sizes: "16x16",
                type: "image/png",
            },
            {
                url: "/favicon-32x32.png",
                sizes: "32x32",
                type: "image/png",
            },
        ],
        apple: [
            {
                url: "/apple-touch-icon.png",
                sizes: "180x180",
                type: "image/png",
            },
        ],
    },

    openGraph: {
        type: "website",
        locale: "vi_VN",
        url: "https://nodepad.vulebaolong.com",
        siteName: "Nodepad",
        title: "Nodepad — Ghi chú online nhanh và riêng tư",
        description:
            "Viết ngay không cần đăng nhập, nhận liên kết để chia sẻ và đồng bộ note trên mọi thiết bị.",
    },

    twitter: {
        card: "summary_large_image",
        title: "Nodepad — Ghi chú online nhanh và riêng tư",
        description:
            "Viết, lưu và chia sẻ ghi chú ngay trên trình duyệt.",
    },

    robots: {
        index: true,
        follow: true,
        googleBot: {
            index: true,
            follow: true,
            "max-image-preview": "large",
            "max-snippet": -1,
            "max-video-preview": -1,
        },
    },
};

export const viewport: Viewport = {
    width: "device-width",
    initialScale: 1,
    themeColor: "#080a12",
    colorScheme: "dark",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="vi" suppressHydrationWarning>
            <head>
                <ColorSchemeScript defaultColorScheme="dark" />
            </head>
            <body className={`${comfortaa.className}`}>
                <Provider>{children}</Provider>
            </body>
        </html>
    );
}
