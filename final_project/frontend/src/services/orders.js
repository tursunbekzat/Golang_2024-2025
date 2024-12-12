import api from "./api";

// Order Services
export const fetchOrders = async () => api.get("/orders").then((res) => res.data);

export const fetchOrderById = async (id) =>
  api.get(`/orders/${id}`).then((res) => res.data);

export const createOrder = async (orderData) =>
  api.post("/orders", orderData).then((res) => res.data);

export const updateOrder = async (id, orderData) =>
  api.put(`/orders/${id}`, orderData).then((res) => res.data);

export const deleteOrder = async (id) =>
  api.delete(`/orders/${id}`).then((res) => res.status === 204);
