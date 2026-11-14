import axios from "axios";

export type User = { id: string; email: string; createdAt: string; updatedAt: string; deletedAt: string | null };
export type Note = { id: string; userId: string | null; shareKey: string; shareUrl: string; title: string | null; content: string; createdAt: string; updatedAt: string; deletedAt: string | null };
export type NoteList = { page: number; pageSize: number; totalItems: number; totalPages: number; items: Note[] };
export type ApiError = { code?: string; message?: string };

export const api = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8888/api",
    withCredentials: true,
    headers: { "Content-Type": "application/json" },
});

export function errorMessage(error: unknown) {
    if (axios.isAxiosError<ApiError>(error)) return error.response?.data?.message ?? "Không thể kết nối đến máy chủ.";
    return "Đã có lỗi xảy ra.";
}
