import React, { useRef } from "react";
import { Container, Input, Text } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { IoIosPaperPlane } from "react-icons/io";

const ChatBox = ({ selectedId, setSelectedId, fetchSessions, setSessions }) => {
  const [loading, setIsLoading] = useState(true);
  const [messages, setMessages] = useState([]);
  const containerRef = useRef();
  const fetchMesagges = async () => {
    console.log("Masuk fetch messages", selectedId);
    setIsLoading(true);
    try {
      const response = await fetch(
        `http://localhost:5000/chat-sessions/${selectedId}/messages`
      );
      const data = await response.json();
      setMessages(data);
      setIsLoading(false);
      console.log("Selesia update messages");
    } catch (error) {
      setMessages([]);
      setIsLoading(false);
    }
  };
  const [text, setText] = useState("");

  useEffect(() => {
    fetchMesagges();
  }, [selectedId]);

  useEffect(() => {
    containerRef.current?.lastChild?.focus();
  }, [messages]);
  if (loading) {
    return <div className="loading">Chatbox....</div>;
  }
  const addMessages = async (e) => {};
  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    if (selectedId) {
      try {
        await fetch(
          `http://localhost:5000/chat-sessions/${selectedId}/messages`,
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              text: text,
            }),
          }
        ).then((response) => response.json());
        setIsLoading(false);
        fetchMesagges();
      } catch (error) {
        console.log(error);
      }
    } else {
      let id = null;
      try {
        const response = await fetch("http://localhost:5000/chat-sessions", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
        });
        const data = await response.json();
        id = data.session_id;
        setSelectedId(() => {
          return id;
        });
        console.log("Fetchnya selesai di dalem chatbox", id);
        try {
          console.log(selectedId);
          await fetch(`http://localhost:5000/chat-sessions/${id}/messages`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              text: text,
            }),
          }).then((response) => response.json());
          fetchSessions();
          fetchMesagges();
          setIsLoading(false);
        } catch (error) {
          console.log(error);
        }
      } catch (error) {
        setIsLoading(false);
      }
    }

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
        flexDirection={"row"}
        flexWrap={"wrap"}
        minW={"100%"}
        mt={10}
        p={0}
        maxH={"75%"}
        overflowY={"scroll"}
        sx={{
          "::-webkit-scrollbar": {
            width: "5px",
          },
          "::-webkit-scrollbar-track": {
            background: "rgb(68,70,84)",
          },
          "::-webkit-scrollbar-thumb": {
            background: "rgba(217,217,227,.8)",
          },
        }}
        ref={containerRef}
      >
        {messages.map((message, index) => {
          if (message.sender === "user") {
            return (
              <Container
                display={"flex"}
                flexDirection="column"
                color={"white"}
                minW="100%"
                padding={6}
                gap={4}
                tabIndex={index + 1}
              >
                <Text fontSize={"16px"} w="70%" textAlign={"left"} mx="auto">
                  {message.text}
                </Text>
              </Container>
            );
          } else if (message.sender === "bot") {
            return (
              <Container
                display={"flex"}
                flexDirection={"column"}
                bgColor={"#444654"}
                minW="100%"
                textAlign={"left"}
                padding={6}
                gap={4}
                tabIndex={index + 1}
              >
                <Text
                  color="white"
                  fontSize={"16px"}
                  w="70%"
                  textAlign={"left"}
                  mx="auto"
                >
                  {message.text}{" "}
                </Text>
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
