import React, { useRef, Fragment } from "react";
import { Container, Input, Text, Button, Stack, Radio, RadioGroup, Box, Image, Textarea} from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { IoIosPaperPlane } from "react-icons/io";
import {CgProfile} from "react-icons/cg"
import ImaGenieKelarImage from "../assets/genieProfile.png"
import UserImage from "../assets/user.png"
import { Palette } from "../assets/palette";

const ChatBox = ({ selectedId, setSelectedId, fetchSessions}) => {
  const [loading, setIsLoading] = useState(true);
  const [messages, setMessages] = useState([]);
  const containerRef = useRef();
  const [radioValue, setRadioValue] = useState("KMP");
  const [text, setText] = useState("");

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

  useEffect(() => {
    fetchMesagges();
  }, [selectedId]);

    useEffect(() => {
        containerRef.current?.lastChild?.focus();
        const lastChild = containerRef.current?.lastChild;
        lastChild && (lastChild.style.outline = "none");
    }, [messages]);



  if (loading) {
    return ;
  }

  const handleKeyDown = (event) => {
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault(); // Prevents submitting the form
      handleSubmit(event);
    }
  };

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
              patternType : radioValue,
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
              patternType : radioValue,
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
  
  const style = { fontSize: "1.5em", color: "white"};
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
      boxShadow={"2xl"}
    >
      { messages.length === 0 ?
        <Box
        minH={"75%"}
        minW={"80%"}
        display={"flex"}
        flexDirection={"column"}
        justifyContent={"center"}
        alignItems={"center"}
        opacity={"0.6"}
        >
            <Text 
            fontSize={"5xl"}
            fontWeight={"bold"}>
                Empty Chat Session
            </Text>
            <Text
            fontSize={"lg"}
            fontWeight={600}>
                Please kindly insert your question inside the chat box
            </Text>
        </Box>
        :
        
      <Container
        display={"flex"}
        // bgColor={Palette.white}
        flexDirection={"column"}
        minW={"80%"}
        mt={5}
        px={5}
        maxH={"75%"}
        overflowY={"scroll"}
        sx={{ 
          "::-webkit-scrollbar": {
            display:"none",
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
                borderWidth={0}
                color={"white"}
                maxW="55%"
                alignSelf={"flex-end"}
                alignItems={"flex-end"}
                py={5}
                px={5}
                gap={3}
                tabIndex={index + 1}
              >
                <Box 
                shadow={"xl"}
                display={"flex"}
                flexDirection={"row"}
                justifyContent={"center"}
                borderWidth={0}
                alignItems={"center"}
                bg={Palette.white}
                borderRadius={"30px"}>
                    <Text pr={6} pl={10} fontSize={"20px"} fontWeight={"bold"} left={0} color={"black"} >
                        You
                    </Text>
                    <Box bg={Palette.dark} borderRadius="30px" p={1} maxW="60px" h="auto">
                        <Image ngColor={"white"} src={UserImage} objectFit="cover" w="100%" h="100%"></Image>
                    </Box>
                </Box>
                
                <Text shadow={"xl"} color={"black"} borderRadius={"2xl"} py={6} px={10} bg={Palette.white} 
                fontSize={"16px"} fontWeight={650} maxW={"100%"} textAlign={"left"}
                transitionDuration={"0.15s"} transitionTimingFunction={"ease-in-out"}
                >
                    {message.text.split('\n').map((line, i) => <Fragment key={i}>{line}<br/></Fragment>)}
                </Text>
              </Box>
            );
          } else if (message.sender === "bot") {
            return (
              <Box
                display={"flex"}
                flexDirection={"column"}
                borderRadius={"lg"}
                color={"white"}
                maxW="55%"
                alignSelf={"flex-start"}
                alignItems={"flex-start"}
                py={5}
                px={5}
                gap={3}
                tabIndex={index + 1}
                __focus={{
                    outline : "none", 
                }}
              >
                <Box 
                display={"flex"}
                shadow={"xl"}
                flexDirection={"row"}
                justifyContent={"center"}
                alignItems={"center"}
                bg={Palette.blue}
                borderRadius={"30px"}>
                    <Box bg={Palette.dark} borderRadius="30px" p={1.5} maxW="60px" h="auto">
                        <Image ngColor={"white"} src={ImaGenieKelarImage}  />
                    </Box>
                    <Text pl={6} pr={10} fontSize={"20px"} fontWeight={"bold"} left={0}>
                        ImaGenieKelar
                    </Text>
                </Box>
                
                <Text shadow={"xl"} borderRadius={"2xl"} py={6} px={10} bg={Palette.blue} 
                fontSize={"16px"} fontWeight={650} maxW={"100%"} textAlign={"left"}
                transitionDuration={"0.15s"} transitionTimingFunction={"ease-in-out"}
                >
                    {message.text.split('\n').map((line, i) => <Fragment key={i}>{line}<br/></Fragment>)}
                </Text>
              </Box>
            );
          }
        })}
      </Container>
      }

        <form onSubmit={handleSubmit} style={{width: "100%",marginBottom:0}}>
            <Box bg={"#BFC8CF"} w={"100%"} display={"flex"} flexDirection={"row"} 
            justifyContent={"center"} alignItems={"center"} py={5} gap={5}>

                <RadioGroup marginRight={5} value={radioValue} justifySelf={"flex-start"}
                borderWidth={"2px"} p={4} borderRadius={"xl"} bgColor={"#525260"} colorScheme="white"
                onChange={() => {radioValue === "KMP" ? setRadioValue("BM") : setRadioValue("KMP")}} >
                    <Stack direction={"row"} gap={3}>
                        <Radio size="md" value='KMP'> <Text color={"white"}>  KMP </Text> </Radio>
                        <Radio size="md" value='BM'> <Text color={"white"}>  BM </Text> </Radio>
                    </Stack>
                </RadioGroup>
                
                <Textarea
                placeholder="Type your questions here..."
                m={0}
                w="60%"
                fontWeight={600}
                maxHeight={"200px"}
                height={"24px"}
                resize={"none"}
                autoSize={{ minRows: 1, maxRows: 2}}
                py={3}
                bgColor={"rgb(64,65,79)"}
                borderColor={"#FFFFFF"}
                borderWidth={"2.5px"}
                borderRadius={"3xl"}
                color="white"
                overflowY="hidden"
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
                value={text}
                onKeyDown={handleKeyDown}
                onChange={(e) => {setText(e.target.value)}}
                _focus={{
                    borderColor:"#FFFFFF",
                    borderWidth: "0",
                    outline: "none"
                }}
                id="textarea"
                />
                <Button 
                display={"flex"}
                justifyContent={"center"}
                alignItems={"center"}
                style={style} type="submit"
                py={6}
                borderColor={"#FFFFFF"}
                borderWidth={"2.5px"}
                borderRadius={"2xl"}
                bgColor={"rgb(64,65,79)"}
                pl={3}
                pr={4}
                _hover = {{
                    bgColor:"rgb(89,90,104)"
                }}
                >
                    <IoIosPaperPlane />
                </Button>

            </Box>
        </form>
    </Container>
  );
};

export default ChatBox;
