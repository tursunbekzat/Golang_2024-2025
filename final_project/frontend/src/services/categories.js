import api from "./api";

// Category Services
export const fetchCategories = async () =>
  api.get("/categories").then((res) => res.data);

export const fetchCategoryById = async (id) =>
  api.get(`/categories/${id}`).then((res) => res.data);

export const createCategory = async (categoryData) =>
  api.post("/categories", categoryData).then((res) => res.data);

export const updateCategory = async (id, categoryData) =>
  api.put(`/categories/${id}`, categoryData).then((res) => res.data);

export const deleteCategory = async (id) =>
  api.delete(`/categories/${id}`).then((res) => res.status === 204);
