import React, { useRef } from "react";
import { Container, Input, Text, Button, Stack, Radio, RadioGroup, Box, Wrap} from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { IoIosPaperPlane } from "react-icons/io";

const ChatBox = ({ selectedId, setSelectedId, fetchSessions}) => {
  const [loading, setIsLoading] = useState(true);
  const [messages, setMessages] = useState([]);
  const containerRef = useRef();
  const fetchMesagges = async () => {
    setIsLoading(true);
    try {
      const response = await fetch(
        `http://localhost:5000/chat-sessions/${selectedId}/messages`
      );
      const data = await response.json();
      data.sort((a,b) =>a.createdDate - b.createdDate)
      setMessages(data);
      setIsLoading(false);
    } catch (error) {
      setMessages([]);
      setIsLoading(false);
    }
  };
  const [text, setText] = useState("");
  const [radioValue, setRadioValue] = useState("KMP");

  useEffect(() => {
    fetchMesagges();
  }, [selectedId]);

  useEffect(() => {
    containerRef.current?.lastChild?.focus();
  }, [messages]);
  if (loading) {
    return <div className="loading">Chatbox....</div>;
  }
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
        fetchMesagges();
        setIsLoading(false);
      } catch (error) {

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
        setSelectedId(id);
        try {
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
          try {
            const response = await fetch(
              `http://localhost:5000/chat-sessions/${id}/messages`
            );
            const data = await response.json();
            setMessages(data);
          } catch (error) {
            setMessages([]);
            setIsLoading(false);
          }
          setIsLoading(false);
        } catch (error) {

        }
      } catch (error) {
        setIsLoading(false);
      }
    }
    setText("");
  };
  
  const style = { fontSize: "1.5em", color: "white", marginBottom: 2 };
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
        // bgColor={"rgb(102 103 115)"}
        flexDirection={"column"}
        minW={"70%"}
        mt={10}
        px={5}
        h={"80%"}
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
              <Box
                display={"flex"}
                flexDirection={"column"}
                borderRadius={"lg"}
                bgColor={"#9496A4"}
                color={"white"}
                w="55%"
                alignSelf={"flex-end"}
                alignItems={"flex-end"}
                py={8}
                px={10}
                gap={2}
                my={3}
                tabIndex={index + 1}
                transitionDuration={"0.8s"}
                transitionTimingFunction={"ease-in-out"}
              >
                <Text fontSize={"20px"} fontWeight={"bold"} left={0}>
                    You:
                </Text>
                <Text fontSize={"16px"} fontWeight={650} maxW={"100%"} textAlign={"right"} mx="auto">
                {message.text}{" "}
                </Text>
              </Box>
            );
          } else if (message.sender === "bot") {
            return (
              <Box
                display={"flex"}
                flexDirection={"column"}
                borderRadius={"lg"}
                bgColor={"#545664"}
                color={"white"}
                w="55%"
                alignSelf={"flex-start"}
                alignItems={"flex-start"}
                py={8}
                px={10}
                gap={2}
                my={3}
                tabIndex={index + 1}
                transitionDuration={"0.8s"}
                transitionTimingFunction={"ease-in-out"}
              >
                <Text fontSize={"20px"} fontWeight={"bold"} left={0}>
                    ImaGenieKelar:
                </Text>
                <Text fontSize={"16px"} fontWeight={650} maxW={"100%"} textAlign={"left"}>
                  {message.text}{" "}
                </Text>
              </Box>
            );
          }
        })}
      </Container>

        <form onSubmit={handleSubmit} style={{width: "100%", justifyContent:"center", alignItems:"center", gap:"0"}}>
            <RadioGroup marginRight={10} value={radioValue} justifySelf={"flex-start"}
            borderWidth={"2px"} p={4} borderRadius={"xl"} bgColor={"#525260"} colorScheme="white"
            onChange={() => {radioValue === "KMP" ? setRadioValue("BM") : setRadioValue("KMP")}} >
                <Stack direction={"row"} gap={3}>
                    <Radio size="md" value='KMP'> <Text color={"white"}>  KMP </Text> </Radio>
                    <Radio size="md" value='BM'> <Text color={"white"}>  BM </Text> </Radio>
                </Stack>
            </RadioGroup>
            
            <Input
            placeholder="Type your questions here..."
            m={0}
            w="60%"
            py={6}
            alignSelf={"center"}
            bgColor={"rgb(64,65,79)"}
            borderWidth={"2.5px"}
            borderLeftRadius={"3xl"}
            borderRightRadius={0}
            borderRightWidth={"0.5px"}
            color="white"
            value={text}
            onChange={(e) => setText(e.target.value)}
            />
            <Button 
            display={"flex"}
            justifyContent={"center"}
            alignItems={"center"}
            style={style} type="submit"
            py={6}
            borderColor={"FFFFFF"}
            borderWidth={"2.5px"}
            borderLeftRadius={0}
            borderRightRadius={"3xl"}
            borderLeftWidth={0}
            bgColor={"rgb(64,65,79)"}
            pl={3}
            pr={4}
            _hover = {{
                bgColor:"rgb(89,90,104)"
            }}
            >
                <IoIosPaperPlane />
            </Button>
        </form>
    </Container>
  );
};

export default ChatBox;
