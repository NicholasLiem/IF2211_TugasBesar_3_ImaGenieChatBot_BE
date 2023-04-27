import React, { Component, useEffect, useState } from "react";
import "./static/App.css";
import { Container, Box, Button } from "@chakra-ui/react";
import Siderbar from "./components/Siderbar";
import ChatBox from "./components/ChatBox";
import SplashScreen from "./components/SplashScreen";
import HelpPage from "./components/Help";

function App() {
  const [selectedId, setSelectedId] = useState(null);
  const [sessions, setSessions] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setIsError] = useState(false);
  const [isNew,setIsNew] = useState(false)
  const fetchSessions = async () => {
    setIsLoading(true);
    console.log("Masuk fetch sessions")
    try {
      const response = await fetch("http://localhost:5000/chat-sessions");
      const data = await response.json();
      console.log("Selesai fetch");
      setSessions(data);
      console.log(data)
      setIsLoading(false);
      setIsError(false);
    } catch (error) {
      console.log(error);
      setIsError(true);
      setIsLoading(false);
    }
  };
  useEffect(()=>{
    console.log("Awal - awal render")
    fetchSessions()
  },[])

  return (
    <div className="App">
      <HelpPage />
      <Container
        bgColor={"rgb(52 53 65)"}
        maxW={"100vw"}
        maxH={"100vh"}
        display={"flex"}
        overflowY={"hidden"}
        flexDirection={"row"}
        p={0}
        m={0}
      >
        <Siderbar
          setSelectedId={setSelectedId}
          sessions={sessions}
          isLoading={isLoading}
          error={error}
          fetchSessions={fetchSessions}
        />
        <ChatBox
          selectedId={selectedId}
          setSelectedId={setSelectedId}
          fetchSessions={fetchSessions}
          setSessions={setSessions}
          setIsNew = {setIsNew}
        />
      </Container>
    </div>
  );
}

export default App;
