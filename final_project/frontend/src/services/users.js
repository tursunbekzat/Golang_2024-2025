import api from "./api";

// User Services
export const fetchUsers = async () => api.get("/users").then((res) => res.data);

export const fetchUserById = async (id) =>
  api.get(`/users/${id}`).then((res) => res.data);

export const createUser = async (userData) =>
  api.post("/users", userData).then((res) => res.data);

export const updateUser = async (id, userData) =>
  api.put(`/users/${id}`, userData).then((res) => res.data);

export const deleteUser = async (id) =>
  api.delete(`/users/${id}`).then((res) => res.status === 204);
