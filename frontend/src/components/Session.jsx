import React, {useEffect, useRef, useState} from "react";
import { Container, Text } from "@chakra-ui/react";
import { BsChatLeft } from "react-icons/bs";

const Session = ({ title }) => {
  const containerRef = useRef(null)
  const [content, setContent] = useState("")
  const style = { color: "white", fontSize: "16px", marginTop: "0.5em" };
  useEffect(()=>{
    const titles = title.split(" ")
    let title_length = 0
    let answers =""
    for(let i = 0; i < titles.length;i++){
      if(title_length > containerRef.current.offsetWidth - 15){
        break
      }else{
        title_length += titles[i].length * 12
        answers += titles[i]
      }
    }
    setContent(answers)
  },[])
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
      _hover={{ bgColor: "#2a2b32", borderRadius:"md" }}
      cursor={"pointer"}
      ref = {containerRef}
    >
      <BsChatLeft style={style} />
      <Text ml={2}>{content}</Text>
    </Container>
  );
};

export default Session;
