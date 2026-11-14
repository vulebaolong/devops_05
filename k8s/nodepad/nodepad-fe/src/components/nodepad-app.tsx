"use client";

import styles from "@/app/home.module.css";
import { errorMessage, Note } from "@/lib/api";
import { AuthMode, nodepadService } from "@/lib/nodepad.service";
import { ActionIcon, Avatar, Button, Group, Loader, Menu, Modal, Pagination, PasswordInput, Textarea, TextInput } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import dayjs from "dayjs";
import Image from "next/image";
import Link from "next/link";
import { useState } from "react";
import { toast } from "sonner";

export default function NodepadApp() {
    const queryClient = useQueryClient();
    const [opened, modal] = useDisclosure(false);
    const [mode, setMode] = useState<AuthMode>("login");
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");
    const [created, setCreated] = useState<Note | null>(null);
    const [page, setPage] = useState(1);
    const [editing, setEditing] = useState<Note | null>(null);
    const [auth, setAuth] = useState({ email: "", password: "" });

    const me = useQuery({ queryKey: ["me"], queryFn: nodepadService.getMe, retry: false });
    const notes = useQuery({
        queryKey: ["notes", page],
        queryFn: () => nodepadService.getNotes(page, 6),
        enabled: !!me.data,
    });
    const createNote = useMutation({
        mutationFn: () => nodepadService.createNote({ title: title.trim() || null, content }),
        onSuccess: (note) => {
            setCreated(note); setTitle(""); setContent("");
            queryClient.invalidateQueries({ queryKey: ["notes"] });
            toast.success("Note đã được lưu");
        },
        onError: (error) => toast.error(errorMessage(error)),
    });
    const authenticate = useMutation({
        mutationFn: () => nodepadService.authenticate(mode, auth),
        onSuccess: async () => {
            modal.close(); setAuth({ email: "", password: "" });
            await queryClient.invalidateQueries({ queryKey: ["me"] });
            toast.success(mode === "login" ? "Đăng nhập thành công" : "Tài khoản đã sẵn sàng");
        },
        onError: (error) => toast.error(errorMessage(error)),
    });
    const logout = useMutation({
        mutationFn: nodepadService.logout,
        onSuccess: () => {
            queryClient.setQueryData(["me"], null);
            queryClient.removeQueries({ queryKey: ["notes"] });
            setCreated(null);
            toast.success("Đã đăng xuất");
        },
        onError: (error) => toast.error(errorMessage(error)),
    });
    const updateNote = useMutation({
        mutationFn: (note: Note) => nodepadService.updateNote(note.id, { title: note.title, content: note.content }),
        onSuccess: () => { setEditing(null); queryClient.invalidateQueries({ queryKey: ["notes"] }); toast.success("Đã cập nhật note"); },
        onError: (error) => toast.error(errorMessage(error)),
    });
    const deleteNote = useMutation({
        mutationFn: nodepadService.deleteNote,
        onSuccess: () => { queryClient.invalidateQueries({ queryKey: ["notes"] }); toast.success("Đã xóa note"); },
        onError: (error) => toast.error(errorMessage(error)),
    });
    const shareLink = (note: Note) => `${window.location.origin}/notes/public/${note.shareKey}`;
    const copy = async (note: Note) => { await navigator.clipboard.writeText(shareLink(note)); toast.success("Đã sao chép liên kết"); };

    return <div className={styles.shell}>
        <div className={styles.orb} /><div className={styles.orbTwo} />
        <header className={styles.header}>
            <Link className={styles.brand} href="/" aria-label="Nodepad — Trang chủ">
                <Image className={styles.logo} src="/logo.png" alt="Nodepad" width={40} height={40} priority />
                <span>nodepad</span>
            </Link>
            <div className={styles.nav}>
                {me.data ? <Menu position="bottom-end" width={245} shadow="xl" radius="md">
                    <Menu.Target>
                        <button className={styles.profileButton} type="button">
                            <Avatar className={styles.avatar} size={34} radius="xl">{me.data.email.charAt(0).toUpperCase()}</Avatar>
                            <span className={styles.profileCopy}><b>{me.data.email.split("@")[0]}</b><small>{me.data.email}</small></span>
                            <span className={styles.chevron}>⌄</span>
                        </button>
                    </Menu.Target>
                    <Menu.Dropdown>
                        <Menu.Label>Tài khoản Nodepad</Menu.Label>
                        <Menu.Item leftSection="✓">Đã đồng bộ an toàn</Menu.Item>
                        <Menu.Divider />
                        <Menu.Item color="red" leftSection="↪" disabled={logout.isPending} onClick={() => logout.mutate()}>Đăng xuất</Menu.Item>
                    </Menu.Dropdown>
                </Menu> : <>
                    <Button variant="subtle" color="gray" onClick={() => { setMode("login"); modal.open(); }}>Đăng nhập</Button>
                    <Button className={styles.primary} onClick={() => { setMode("register"); modal.open(); }}>Tạo tài khoản</Button>
                </>}
            </div>
        </header>
        {me.data && <section className={styles.memberBar}>
            <div className={styles.memberGreeting}>
                <div className={styles.statusIcon}>✓</div>
                <div><span>Không gian của bạn</span><strong>Chào mừng trở lại, {me.data.email.split("@")[0]}</strong></div>
            </div>
            <div className={styles.memberStats}>
                <div><b>{notes.data?.totalItems ?? "—"}</b><span>note đã lưu</span></div>
                <div className={styles.syncState}><i /><span>Đồng bộ trực tuyến</span></div>
                <Button variant="light" color="violet" component="a" href="#write">+ Note mới</Button>
            </div>
        </section>}
        <main className={styles.main}>
            <section className={styles.hero}>
                <div className={styles.eyebrow}><span className={styles.pulse} /> {me.data ? "Tự động lưu vào thư viện của bạn" : "Không cần đăng nhập · Riêng tư theo liên kết"}</div>
                <h1 className={styles.title}><span className={styles.gradientText}>{me.data ? "Không gian của bạn." : "Ý tưởng đến nhanh"}</span><br />{me.data ? "Mọi note, một nơi." : "Ghi lại còn nhanh hơn."}</h1>
                <p className={styles.subtitle}>{me.data ? "Viết điều đang nghĩ, Nodepad sẽ giữ chúng an toàn và sẵn sàng trên mọi thiết bị của bạn." : "Một không gian tối giản để viết, lưu và chia sẻ suy nghĩ. Không quảng cáo, không xao nhãng — chỉ bạn và những điều đáng nhớ."}</p>
            </section>
            <section className={styles.composer} id="write"><div className={styles.composerInner}>
                <TextInput className={styles.titleInput} value={title} onChange={(e) => setTitle(e.currentTarget.value)} placeholder="Tiêu đề note (không bắt buộc)" maxLength={200} />
                <Textarea className={styles.contentInput} value={content} onChange={(e) => setContent(e.currentTarget.value)} placeholder="Bắt đầu viết điều gì đó tuyệt vời..." autosize={false} />
                <div className={styles.composerFooter}>
                    <span className={styles.hint}>{me.data ? "Note này sẽ được lưu vào tài khoản của bạn" : "Bạn sẽ nhận được một liên kết để chia sẻ"}</span>
                    <Button className={styles.primary} loading={createNote.isPending} disabled={!content.trim()} onClick={() => createNote.mutate()}>Lưu & chia sẻ →</Button>
                </div>
            </div></section>
            {created && <div className={styles.success}><div className={styles.successText}><b>Liên kết của bạn đã sẵn sàng</b><span>{shareLink(created)}</span></div><Button variant="light" color="teal" onClick={() => copy(created)}>Sao chép link</Button></div>}
            {me.data && <section className={styles.notesSection}>
                <div className={styles.sectionHead}><div><h2>Thư viện của bạn</h2><p>{notes.data?.totalItems ?? 0} note được lưu an toàn</p></div>{notes.isFetching && <Loader size="sm" color="violet" />}</div>
                <div className={styles.grid}>
                    {notes.data?.items.map((note) => <article className={styles.noteCard} key={note.id}>
                        <div className={styles.noteTop}><h3>{note.title || "Note không tiêu đề"}</h3><div className={styles.cardActions}>
                            <ActionIcon variant="subtle" color="gray" aria-label="Chia sẻ" onClick={() => copy(note)}>↗</ActionIcon>
                            <ActionIcon variant="subtle" color="violet" aria-label="Sửa" onClick={() => setEditing({ ...note })}>✎</ActionIcon>
                            <ActionIcon variant="subtle" color="red" aria-label="Xóa" onClick={() => confirm("Xóa note này?") && deleteNote.mutate(note.id)}>×</ActionIcon>
                        </div></div><p>{note.content}</p><span className={styles.noteMeta}>Chỉnh sửa {dayjs(note.updatedAt).fromNow()}</span>
                    </article>)}
                    {!notes.isLoading && !notes.data?.items.length && <div className={styles.empty}>Note đầu tiên của bạn sẽ xuất hiện ở đây.</div>}
                </div>
                {!!notes.data?.totalPages && notes.data.totalPages > 1 && <div className={styles.pagination}><Pagination value={page} onChange={setPage} total={notes.data.totalPages} color="violet" /></div>}
            </section>}
        </main>
        <Modal opened={opened} onClose={modal.close} centered title={<span className={styles.modalTitle}>{mode === "login" ? "Chào mừng trở lại" : "Tạo tài khoản"}</span>} overlayProps={{ backgroundOpacity: .72, blur: 8 }} radius="lg">
            <p className={styles.authCopy}>{mode === "login" ? "Mở lại thư viện note của bạn trên mọi thiết bị." : "Lưu trữ, chỉnh sửa và quản lý mọi note ở một nơi."}</p>
            <form className={styles.form} onSubmit={(e) => { e.preventDefault(); authenticate.mutate(); }}>
                <TextInput label="Email" type="email" required value={auth.email} onChange={(e) => setAuth({ ...auth, email: e.currentTarget.value })} placeholder="you@example.com" />
                <PasswordInput label="Mật khẩu" required minLength={5} value={auth.password} onChange={(e) => setAuth({ ...auth, password: e.currentTarget.value })} placeholder="Ít nhất 5 ký tự" />
                <Button className={styles.primary} type="submit" loading={authenticate.isPending}>{mode === "login" ? "Đăng nhập" : "Bắt đầu miễn phí"}</Button>
                <div className={styles.switchAuth}>{mode === "login" ? "Chưa có tài khoản?" : "Đã có tài khoản?"} <button type="button" className={styles.linkButton} onClick={() => setMode(mode === "login" ? "register" : "login")}>{mode === "login" ? "Đăng ký" : "Đăng nhập"}</button></div>
            </form>
        </Modal>
        <Modal opened={!!editing} onClose={() => setEditing(null)} centered title={<span className={styles.modalTitle}>Chỉnh sửa note</span>} radius="lg">
            {editing && <div className={styles.form}>
                <TextInput label="Tiêu đề" value={editing.title ?? ""} onChange={(e) => setEditing({ ...editing, title: e.currentTarget.value })} />
                <Textarea label="Nội dung" minRows={10} autosize value={editing.content} onChange={(e) => setEditing({ ...editing, content: e.currentTarget.value })} />
                <Group justify="flex-end"><Button variant="default" onClick={() => setEditing(null)}>Hủy</Button><Button className={styles.primary} loading={updateNote.isPending} onClick={() => updateNote.mutate(editing)}>Lưu thay đổi</Button></Group>
            </div>}
        </Modal>
    </div>;
}
