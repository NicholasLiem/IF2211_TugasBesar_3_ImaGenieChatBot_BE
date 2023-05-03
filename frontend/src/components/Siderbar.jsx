import React from "react";
import { Container, Text } from "@chakra-ui/react";
import { IoIosAdd } from "react-icons/io";
import Session from "./Session";
import { Palette } from "../assets/palette";
import { useGlobalContext } from "../context";

const Siderbar = ({
  sessions,
  fetchSessions,
}) => {
  const {setSelectedId} = useGlobalContext()
  const style = {fontSize: "24px" };
  return (
    <Container
      maxW={"20%"}
      display={"flex"}
      flexDirection={"column"}
      bgColor={Palette.dark}
      minH={"100vh"}
      left={0}
      top={0}
      m={0}
      px={1}
      boxShadow={"2xl"}
    >
      <Container
        top={0}
        border={"1px solid "+Palette.white}
        borderRadius={"md"}
        my={8}
        w={"70%"}
        py={4}
        color={Palette.white}
        cursor={"pointer"}
        display={"flex"}
        flexWrap={"wrap"}
        flexDirection={"row"}
        transitionDuration={"0.2s"}
        transitionTimingFunction={"ease-in-out"}
        _hover={{ borderColor:Palette.dark,  w:"90%", bgColor: Palette.white, color: Palette.dark}}
        onClick={() => setSelectedId(null)}
        boxShadow={"2xl"}
      >
        <IoIosAdd style={style} />
        <Text ml={2} fontWeight={700}>
          New Chat
        </Text>
      </Container>
      <Container
        display={"flex"}
        flexDirection={"column"}
        mb={8}
        px={0}
        overflowY={"auto"}
        sx={{
          "::-webkit-scrollbar": {
            width: "5px",
          },
          "::-webkit-scrollbar-track": {
            background: "rgb(108,110,124)",
          },
          "::-webkit-scrollbar-thumb": {
            background: "rgba(217,217,227,.8)",
          },
        }}
      >
        {sessions.map((item) => {
          return (
            <Session
              id={item.id}
              fetchSessions = {fetchSessions}
            />
          );
        })}
      </Container>
    </Container>
  );
};

export default Siderbar;
