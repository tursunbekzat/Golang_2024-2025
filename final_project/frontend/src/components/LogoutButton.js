import React from "react";
import { logout } from "../services/auth";
import { useNavigate } from "react-router-dom";

const LogoutButton = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate("/login"); // Redirect to the login page
  };

  return <button onClick={handleLogout}>Logout</button>;
};

export default LogoutButton;
