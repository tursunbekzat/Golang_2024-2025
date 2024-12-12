import React, { useEffect, useState } from "react";
import { fetchUsers, deleteUser } from "../services/users";

const UserList = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    const loadUsers = async () => {
      const data = await fetchUsers();
      setUsers(data);
    };
    loadUsers();
  }, []);

  const handleDelete = async (id) => {
    await deleteUser(id);
    setUsers(users.filter((user) => user.user_id !== id));
  };

  return (
    <div>
      <h1>Users</h1>
      <ul>
        {users.map((user) => (
          <li key={user.user_id}>
            {user.username} - {user.email}{" "}
            <button onClick={() => handleDelete(user.user_id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default UserList;
