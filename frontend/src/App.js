import React, { Component } from "react";
import "./static/App.css";
import { connect, sendMsg } from "./api";
import {Container } from "@chakra-ui/react";
import Siderbar from "./components/Siderbar";
import ChatBox from "./components/ChatBox";

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    return (
      <div className="App">
        <Container
          bgColor={"rgb(52 53 65)"}
          maxW={"100vw"}
          overflowX={"hidden"}
          minH={"100vh"}
          display={"flex"}
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
