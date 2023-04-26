import React from "react";
import { Container, Input, Text } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { IoIosPaperPlane } from "react-icons/io";
import SessionPage from "./SessionPage";

const ChatBox = ({ selectedId }) => {
  const [loading, setIsLoading] = useState(true);
  const [messages, setMessages] = useState([]);
  const fetchMesagges = async () => {
    setIsLoading(true);
    try {
      const response = await fetch(
        `http://localhost:5000/chat-sessions/${selectedId}/messages`
      );
      const data = await response.json();
      setMessages(data);
      setIsLoading(false);
    } catch (error) {
      setMessages([]);
      setIsLoading(false);
    }
  };
  const [text, setText] = useState("");

  useEffect(() => {
    fetchMesagges();
  }, [selectedId]);
  if (loading) {
    return <div className="loading">Loading....</div>;
  }
  const handleSubmit = async (e) => {
    e.preventDefault();
    setText("");
  };
  const style = { fontSize: "2em", color: "white", marginBottom: 2 };
  return (
    <Container
      m={0}
      display={"flex"}
      flexDirection={"column"}
      position={"relative"}
      alignContent={"center"}
      alignItems={"center"}
      flexWrap={"wrap"}
      maxW="100%"
      p={0}
    >
      <Container
        display={"flex"}
        flexDirection={"column"}
        flexWrap={"wrap"}
        minW={"100%"}
        mt={10}
        p={0}
        minH={"100%"}
      >
        {messages.map((message) => {
          if (message.sender === "user") {
            return (
              <Container
                display={"flex"}
                flexDirection="column"
                color={"white"}
                minW="100%"
                m={0}
                textAlign={"left"}
                padding={6}
                paddingLeft={10}
              >
                <Text fontSize={"16px"} ml={10}> {message.text} </Text>
              </Container>
            );
          } else if(message.sender==="bot"){
            return (
              <Container
                display={"flex"}
                flexDirection={"column"}
                bgColor={"#444654"}
                minW="100%"
                textAlign={"left"}
                padding={6}
                paddingLeft={10}
              >
                <Text color="white" fontSize={"16px"} ml={10}>{message.text} </Text>
              </Container>
            );
          }
        })}
      </Container>

      <form onSubmit={handleSubmit}>
        <Input
          placeholder="Type your questions here"
          m={0}
          w="60vw"
          alignSelf={"center"}
          bgColor={"rgb(64,65,79)"}
          py={6}
          borderRadius={"md"}
          color="white"
          value={text}
          onChange={(e) => setText(e.target.value)}
        />
        <button style={style} type="submit">
          <IoIosPaperPlane />
        </button>
      </form>
    </Container>
  );
};

export default ChatBox;
