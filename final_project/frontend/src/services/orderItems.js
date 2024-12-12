import api from "./api";

// OrderItems Services
export const fetchOrderItems = async () =>
  api.get("/order-items").then((res) => res.data);

export const fetchOrderItemsByOrderId = async (orderId) =>
  api.get(`/order-items/${orderId}`).then((res) => res.data);

export const createOrderItem = async (orderItemData) =>
  api.post("/order-items", orderItemData).then((res) => res.data);

export const deleteOrderItem = async (id) =>
  api.delete(`/order-items/${id}`).then((res) => res.status === 204);
