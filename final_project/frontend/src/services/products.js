import api from "./api";

// Product Services
export const fetchProducts = async () =>
  api.get("/products").then((res) => res.data);

export const fetchProductById = async (id) =>
  api.get(`/products/${id}`).then((res) => res.data);

export const createProduct = async (productData) =>
  api.post("/products", productData).then((res) => res.data);

export const updateProduct = async (id, productData) =>
  api.put(`/products/${id}`, productData).then((res) => res.data);

export const deleteProduct = async (id) =>
  api.delete(`/products/${id}`).then((res) => res.status === 204);
