"use client";

import styles from "@/app/home.module.css";
import { errorMessage } from "@/lib/api";
import { nodepadService } from "@/lib/nodepad.service";
import { Button, Loader, Text } from "@mantine/core";
import { useQuery } from "@tanstack/react-query";
import dayjs from "dayjs";
import Image from "next/image";
import Link from "next/link";
import { useParams } from "next/navigation";

export default function PublicNotePage() {
    const { shareKey } = useParams<{ shareKey: string }>();
    const note = useQuery({
        queryKey: ["public-note", shareKey],
        queryFn: () => nodepadService.getPublicNote(shareKey),
    });

    return <main className={styles.publicWrap}>
        <article className={styles.publicCard}>
            <Link className={styles.brand} href="/" aria-label="Nodepad — Trang chủ">
                <Image className={styles.logo} src="/logo.png" alt="Nodepad" width={40} height={40} priority />
                <span>nodepad</span>
            </Link>
            {note.isLoading && <div className={styles.empty}><Loader color="violet" /></div>}
            {note.isError && <div className={styles.empty}>{errorMessage(note.error)}<br /><br /><Button component={Link} href="/" variant="light" color="violet">Tạo note mới</Button></div>}
            {note.data && <>
                <h1>{note.data.title || "Note không tiêu đề"}</h1>
                <Text c="dimmed" size="xs">Được tạo {dayjs(note.data.createdAt).format("DD/MM/YYYY · HH:mm")}</Text>
                <div className={styles.publicContent}>{note.data.content}</div>
                <Button component={Link} href="/" variant="light" color="violet" mt={35}>Viết note của bạn →</Button>
            </>}
        </article>
    </main>;
}
