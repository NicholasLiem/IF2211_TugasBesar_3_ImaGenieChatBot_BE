import React, { useEffect, useRef, useState } from "react";
import { Container, Text } from "@chakra-ui/react";
import { BsChatLeft } from "react-icons/bs";

const Session = ({ id, setSelectedId }) => {
  const containerRef = useRef(null);
  const [content, setContent] = useState("");
  const [firstQuestion, setFirstQuestion] = useState("");
  const [isLoading, setIsLoading] = useState(true);
  const [isError, setIsError] = useState(false);
 
  const getTitle = () => {
    const titles = firstQuestion.split(" ")
    let total_length = 0
    let answers = ""
    for(let i = 0; i < titles.length;i++){
      if(total_length + titles[i].length*12 < 240){
        answers += (titles[i] + " ")
        total_length += titles[i].length * 12
      }else{
        break
      }
    }
    answers += "..."
    return(answers)
  };
  const fetchData = async () => {
    setIsLoading(true);
    try {
      const link = `http://localhost:5000/chat-sessions/${id}/messages`;
      const response = await fetch(link);
      const data = await response.json();
      data[0].sender === "user"
        ? setFirstQuestion(data[0].text)
        : setFirstQuestion(data[1].text);
      setContent(data);
      setIsLoading(false);
      setIsError(false);
    } catch (error) {
      setIsError(true);
      setIsLoading(false);
    }
  };
  const style = { color: "white", fontSize: "16px", marginTop: "0.5em" };
  useEffect(() => {
    fetchData();
  }, []);
  if (isLoading) {
    console.log("Loading di session...");
    return <div className="loading">Loading...</div>;
  }
  if (isError) {
    console.log("error...");
    console.log(id);
    return <div className="error">Error...</div>;
  }
  return (
    <Container
      color={"white"}
      alignItems={"center"}
      alignContent={"center"}
      display={"flex"}
      flexDirection={"row"}
      mt={4}
      p={0}
      pl={2}
      py={3}
      _hover={{ bgColor: "#2a2b32", borderRadius: "md" }}
      cursor={"pointer"}
      ref={containerRef}
      onClick={() => setSelectedId(id)}
    >
      <BsChatLeft style={style} />
      <Text ml={2}>{getTitle()}</Text>
    </Container>
  );
};

export default Session;
