import React, { Component } from "react";
import "./static/App.css";
import {Container } from "@chakra-ui/react";
import Siderbar from "./components/Siderbar";
import ChatBox from "./components/ChatBox";


class App extends Component {
  render() {
    return (
      <div className="App">
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
          <Siderbar />
          <ChatBox />
        </Container>
      </div>
    );
  }
}

export default App;
