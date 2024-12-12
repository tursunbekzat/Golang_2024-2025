import api from "./api";

// Session Services
export const fetchSessions = async () =>
  api.get("/sessions").then((res) => res.data);

export const createSession = async (sessionData) =>
  api.post("/sessions", sessionData).then((res) => res.data);

export const deleteSession = async (id) =>
  api.delete(`/sessions/${id}`).then((res) => res.status === 204);
