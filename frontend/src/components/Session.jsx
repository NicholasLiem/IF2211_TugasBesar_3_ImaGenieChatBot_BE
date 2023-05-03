import React, { useEffect, useRef, useState } from "react";
import { Container, Text } from "@chakra-ui/react";
import { BsChatLeft, BsTrash } from "react-icons/bs";
import { Palette } from "../assets/palette";
import { useGlobalContext } from "../context";

const Session = ({ id, fetchSessions }) => {
  const containerRef = useRef(null);
  const [firstQuestion, setFirstQuestion] = useState("");
  // const [isLoading, setIsLoading] = useState(true);
  // const [isError, setIsError] = useState(false);
  const {   
    selectedId,
    setSelectedId,
  } = useGlobalContext();
  
  const getTitle = () => {
    const titles = firstQuestion.split(" ");
    let total_length = 0;
    let answers = "";
    for (let i = 0; i < titles.length; i++) {
      if (total_length + titles[i].length * 12 < 180) {
        answers += titles[i] + " ";
        total_length += titles[i].length * 12;
      } else {
        break;
      }
    }
    answers += "...";
    return answers;
  };
  const fetchData = async () => {
    try {
      const link = `http://localhost:5000/chat-sessions/${id}/messages`;
      const response = await fetch(link);
      const data = await response.json();
      data[0].sender === "user"
        ? setFirstQuestion(data[0].text)
        : setFirstQuestion(data[1].text);
     
    } catch (error) {
      
    }
  };

  const handleDelete = async () => {
    await fetch(`http://localhost:5000/chat-sessions/${id}`, {
      method: "DELETE",
    });
    setSelectedId(null);
    fetchSessions();
  };
  const style = { color: "white", fontSize: "16px", marginTop: "0.5em" };
  useEffect(() => {
    fetchData();
  }, [id, setSelectedId]);

  return (
    <Container
      color={"white"}
      alignItems={"flex-start"}
      alignContent={"flex-start"}
      display={"flex"}
      flexDirection={"row"}
      p={0}
      pl={5}
      py={5}
      w={"90%"}
      borderWidth={"2px"}
      borderColor={Palette.dark}
      _hover={{ borderColor: Palette.white, borderRadius: "md" }}
      cursor={"pointer"}
      bgColor={selectedId === id ? Palette.gray : ""}
      ref={containerRef}
      onClick={() => setSelectedId(id)}
      position={"relative"}
    >
      <BsChatLeft style={style} />
      <Text ml={2}>{getTitle() ? getTitle() : "New chat"}</Text>
      {selectedId === id && (
        <BsTrash
          style={{ position: "absolute", right: 15, fontSize: "20px" }}
          onClick={handleDelete}
        />
      )}
    </Container>
  );
};

export default Session;
