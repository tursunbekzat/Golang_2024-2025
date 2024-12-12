import api from "./api";

// Payment Services
export const fetchPayments = async () => api.get("/payments").then((res) => res.data);

export const fetchPaymentById = async (id) =>
  api.get(`/payments/${id}`).then((res) => res.data);

export const createPayment = async (paymentData) =>
  api.post("/payments", paymentData).then((res) => res.data);

export const deletePayment = async (id) =>
  api.delete(`/payments/${id}`).then((res) => res.status === 204);
