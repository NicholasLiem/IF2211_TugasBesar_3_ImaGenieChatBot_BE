import React from "react";
import { Container, Input, calc } from "@chakra-ui/react";
import { useEffect } from "react";
const ChatBox = () => {

  return (
    <Container
      m={0}
      display={"flex"}
      flexDirection={"column"}
      position={"relative"}
      alignContent={"center"}
      alignItems={"center"}
      flexWrap={"wrap"}
      maxW = "100%"
      p = {0}
    >
      <Input
        placeholder="Type your questions here"
        bottom={5}
        mb={5}
        mx="auto"
        w="60vw"
        alignSelf={"center"}
        position={"absolute"}
        bgColor={"rgb(64,65,79)"}
        py={6}
        borderRadius={"md"}
        color="white"
      />
    </Container>
  );
};

export default ChatBox;
