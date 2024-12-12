import React, { useEffect, useState } from "react";
import { fetchSessions, deleteSession } from "../services/sessions";

const SessionsList = () => {
  const [sessions, setSessions] = useState([]);

  useEffect(() => {
    const loadSessions = async () => {
      const data = await fetchSessions();
      setSessions(data);
    };
    loadSessions();
  }, []);

  const handleDelete = async (id) => {
    await deleteSession(id);
    setSessions(sessions.filter((session) => session.session_id !== id));
  };

  return (
    <div>
      <h1>Sessions</h1>
      <ul>
        {sessions.map((session) => (
          <li key={session.session_id}>
            User: {session.user_id}, Expires: {session.expires_at}{" "}
            <button onClick={() => handleDelete(session.session_id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default SessionsList;
