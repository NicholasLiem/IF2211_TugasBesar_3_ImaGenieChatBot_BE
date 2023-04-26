import React, { useEffect, useState } from "react";
import { Container, Text } from "@chakra-ui/react";
import { IoIosAdd } from "react-icons/io";
import Session from "./Session";

const Siderbar = ({setSelectedId}) => {
  const [sessions, setSessions] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setIsError] = useState(false);

  const handleIdChange = (id) =>{
    setSelectedId(id)
  }
  const fetchSessions = async () => {
    setIsLoading(true);
    try {
      const response = await fetch("http://localhost:5000/chat-sessions");
      const data = await response.json();
      console.log("Selesai fetch")
      setSessions(data);
      setIsLoading(false);
      setIsError(false);
    } catch (error) {
      console.log(error)
      setIsError(true);
      setIsLoading(false);
    }
  };
  useEffect(() => {
    fetchSessions();
  }, []);
  if(isLoading){
    console.log("Loading...")
    return <div className="loading">Loading...</div>
  }

  if(error){
    console.log("Error sidebar")
    return <div className="error">Error...</div>
  }

  const style = { color: "white", fontSize: "24px" };
  return (
    <Container
      maxW={"20%"}
      display={"flex"}
      flexDirection={"column"}
      flexWrap={"wrap"}
      bgColor={"rgb(32,33,35)"}
      minH={"100vh"}
      left={0}
      top={0}
      m={0}
      px={1}
    >
      <Container

        top={0}
        border={"1px solid #d9d9e3"}
        borderRadius={"md"}
        mt={4}
        py={3}
        cursor={"pointer"}
        display={"flex"}
        flexWrap={"wrap"}
        flexDirection={"row"}
        _hover={{ opacity: "0.75", brightness: "1.3" }}
      >
        <IoIosAdd style={style} />
        <Text color={"white"} ml={2}>
          New Chat
        </Text>
      </Container>
      <Container
        display={"flex"}
        flexDirection={"column"}
        flexWrap={"wrap"}
        p={0}
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
      >
        {sessions.map((item) => {
          return <Session id={item.id} setSelectedId = {handleIdChange} />;
        })}
      </Container>
    </Container>
  );
};

export default Siderbar;
