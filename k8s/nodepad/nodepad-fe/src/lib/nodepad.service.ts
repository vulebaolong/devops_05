import { api, Note, NoteList, User } from "./api";
import { API_ENDPOINTS } from "./api-endpoints";

export type AuthMode = "login" | "register";
export type AuthInput = { email: string; password: string };
export type CreateNoteInput = { title: string | null; content: string };
export type UpdateNoteInput = Pick<Note, "title" | "content">;

export const nodepadService = {
    async getMe() {
        return (await api.get<User>(API_ENDPOINTS.auth.me)).data;
    },

    async authenticate(mode: AuthMode, input: AuthInput) {
        return (await api.post(API_ENDPOINTS.auth[mode], input)).data;
    },

    async logout() {
        return (await api.post(API_ENDPOINTS.auth.logout)).data;
    },

    async getNotes(page: number, pageSize: number) {
        return (await api.get<NoteList>(API_ENDPOINTS.notes.root, {
            params: { page, pageSize },
        })).data;
    },

    async getPublicNote(shareKey: string) {
        return (await api.get<Note>(API_ENDPOINTS.notes.public(shareKey))).data;
    },

    async createNote(input: CreateNoteInput) {
        return (await api.post<Note>(API_ENDPOINTS.notes.root, input)).data;
    },

    async updateNote(id: string, input: UpdateNoteInput) {
        return (await api.patch<Note>(API_ENDPOINTS.notes.detail(id), input)).data;
    },

    async deleteNote(id: string) {
        return api.delete<Note>(API_ENDPOINTS.notes.detail(id));
    },
};
