import {Container, Box, Text} from "@chakra-ui/react"
import React, {useState} from "react";


const HelpPage = () => {

    const [layerPos, setLayerPos] = useState(-3000)
    const [boxPos, setBoxPos] = useState(-3000)
    const [text, setText] = useState("?")

    return (
        <Container>
            <Box
                h={"30px"}
                w={"30px"}  
                bg={"transparent"}      
                position={"absolute"}
                top={"8"}
                right={"8"}
                color={"#FFFFFF"}
                borderRadius={"15px"}
                borderWidth={"3px"}
                fontWeight={"bold"}
                _hover={{
                    bgColor:"#FFFFFF",
                    color: "rgb(52 53 65)",
                }}
                zIndex={6}
                onClick={() => { 
                    text === "?" ? setText("X") : setText("?");
                    layerPos === 0 ? setLayerPos(-3000) : setLayerPos(0);
                    boxPos === "20vw" ? setBoxPos(-3000) : setBoxPos("20vw");
            }}
            >
                {text}
            </Box>

            {/* Layer */}
            <Box
            position={"absolute"}
            bgColor={"#000000"}
            opacity={0.7}
            top={0}
            left={layerPos}
            zIndex={5}
            w={"100%"}
            h={"100%"}
            transitionDuration={"0.5s"}
            transitionTimingFunction={"ease-in-out"}
            />

            {/* Content */}
            <Box
            position={"absolute"}
            bgColor={"#AFAFBF"}
            borderRadius={"10px"}
            top={"10vh"}
            left={boxPos}
            zIndex={6}
            w={"60vw"}
            h={"80vh"}
            transitionDuration={"1s"}
            transitionTimingFunction={"ease-in-out"}
            padding={"20px"}
            > 
                <Text fontSize={"3xl"} fontWeight={600}>
                    Instructions
                </Text>
            </Box>
                
        </Container>
    );
}



export default HelpPage;