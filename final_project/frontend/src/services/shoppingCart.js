import api from "./api";

// ShoppingCart Services
export const fetchShoppingCarts = async () =>
  api.get("/shopping-carts").then((res) => res.data);

export const fetchShoppingCartById = async (id) =>
  api.get(`/shopping-carts/${id}`).then((res) => res.data);

export const createShoppingCart = async (cartData) =>
  api.post("/shopping-carts", cartData).then((res) => res.data);

export const deleteShoppingCart = async (id) =>
  api.delete(`/shopping-carts/${id}`).then((res) => res.status === 204);
