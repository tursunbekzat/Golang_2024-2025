import api from "./api";

// Review Services
export const fetchReviews = async () => api.get("/reviews").then((res) => res.data);

export const createReview = async (reviewData) =>
  api.post("/reviews", reviewData).then((res) => res.data);

export const deleteReview = async (id) =>
  api.delete(`/reviews/${id}`).then((res) => res.status === 204);
