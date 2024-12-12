import api from "../services/api";

// Login Service
export const login = async (credentials) => {
  const response = await api.post("/login", credentials);
  const token = response.data.token;

  // Store the token in localStorage
  localStorage.setItem("token", token);

  return token;
};
export const logout = () => {
    localStorage.removeItem("token"); // Remove the token from localStorage
  };