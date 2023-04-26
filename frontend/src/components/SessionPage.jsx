import React, { useEffect, useState } from "react";

const SessionPage = ({ id }) => {
  const [messages, setMessages] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isError, setIsError] = useState(false);
  const fetchSessions = async () => {
    setIsLoading(true);
    try {
      const response = await fetch(
        `http://localhost:5000/chat-messages/${id}/messages`
      );
      const data = await response.json();
      console.log(data);
      setIsLoading(false);
      setIsError(false);
    } catch (error) {
      setIsError(true);
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchSessions();
  }, []);

  if (isLoading) {
    return <div className="loading">Loading...</div>;
  }

  if (isError) {
    return <div className="error">Error....</div>;
  }

  return <div></div>;
};

export default SessionPage;
