import React, { Component, useState } from "react";
import "./static/App.css";
import { Container, Box, Button } from "@chakra-ui/react";
import Siderbar from "./components/Siderbar";
import ChatBox from "./components/ChatBox";
import SplashScreen from "./components/SplashScreen";
import HelpPage from "./components/Help";

function App() {
  const [selectedId, setSelectedId] = useState(null)

  const select = (id) =>{
    this.setSelectedId({id})
  }
  return (
    <div className="App">
      <SplashScreen />
      <HelpPage />
      <Container
        bgColor={"rgb(52 53 65)"}
        maxW={"100vw"}
        minH={"100vh"}
        display={"flex"}
        overflowY={"hidden"}
        flexDirection={"row"}
        p={0}
        m={0}
      >
        <Siderbar setSelectedId={setSelectedId} />
        <ChatBox selectedId={selectedId} />
      </Container>
    </div>
  );
}

export default App;
