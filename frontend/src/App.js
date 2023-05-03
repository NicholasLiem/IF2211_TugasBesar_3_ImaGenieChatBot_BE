import React, {useEffect, useState } from "react";
import "./static/App.css";
import { Container} from "@chakra-ui/react";
import Siderbar from "./components/Siderbar";
import ChatBox from "./components/ChatBox";
import SplashScreen from "./components/SplashScreen";
import HelpPage from "./components/Help";
import { Palette } from "./assets/palette";

function App() {
  const [sessions, setSessions] = useState([]);
  const fetchSessions = async () => {
    try {
      const response = await fetch("http://localhost:5000/chat-sessions");
      const data = await response.json();
      setSessions(data.reverse());
    } catch (error) {}
  };
  useEffect(()=>{
    fetchSessions()
  },[])

  return (
    <div className="App">
      {/* <SplashScreen /> */}
      <HelpPage />
      <Container
        bgColor={Palette.lightGray}
        maxW={"100vw"}
        maxH={"100vh"}
        display={"flex"}
        overflowY={"hidden"}
        flexDirection={"row"}
        p={0}
        m={0}
      >
        <Siderbar
          sessions={sessions}
          fetchSessions={fetchSessions}
        />
        <ChatBox
          fetchSessions={fetchSessions}
          setSessions={setSessions}
        />
      </Container>
    </div>
  );
}

export default App;
