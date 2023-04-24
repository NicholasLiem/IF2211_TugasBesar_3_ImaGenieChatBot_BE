import React from "react";
import { Container, Icon, Text } from "@chakra-ui/react";
import { IoIosAdd } from "react-icons/io";
import Session from "./Session";
const data = [
  "Kenapa matahari terbit?",
  "Kenapa air enak?",
  "Berapa 1 + 1?",
  "Siapa pencipta lampu?",
  "Apakah alien nyata?",
  "Gimana cara stop tubes?",
  "Kok stima tubesnya banyak?",
  "Gimana cara jalan ke itb?",
  "Coba tes kalimat nya",
];
const Siderbar = () => {
  const style = { color: "white", fontSize: "24px" };
  return (
    <Container
      w={"280px"}
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
          "::-webkit-scrollbar" :{
            width: "5px"
          },
          "::-webkit-scrollbar-track" :{
            background: "rgb(68,70,84)"
          },
          "::-webkit-scrollbar-thumb" :{
            background :"rgba(217,217,227,.8)"
          }
        }}
      >
        {data.map((item) => {
          return <Session title={item} />;
        })}
      </Container>
    </Container>
  );
};

export default Siderbar;
