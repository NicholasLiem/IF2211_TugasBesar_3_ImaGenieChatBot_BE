import React, { Component } from "react";
import "./static/App.css";
import { connect, sendMsg } from "./api";
import { Button} from "@chakra-ui/react";

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
            <Button 
            colorScheme='blue' 
            onClick={this.send} 
            rounded='10px'>Send Message</Button>
        </div>
    );
  }

}

export default App;

