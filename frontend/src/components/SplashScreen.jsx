import {Box, Image, Text} from "@chakra-ui/react"
import React, {useState} from "react";
import GenieLamp from "../assets/genieLamp.png"
import GenieLampOpen from "../assets/genieLampOpen.png"
import Cloud from "../assets/cloud.png"

const SplashScreen  = () =>{
    const [top, setTop] = useState(0);
    const splashImage = {GenieLamp, GenieLampOpen};
    const [logo, setLogo] = useState(splashImage.GenieLamp);
    const [pos, setPos] = useState("30%")
    const [rot, setRot] = useState("rotate(0def)")

    return (
        <Box
        position={"absolute"}
        display={"flex"}
        flexDirection={"column"}
        bgColor={"rgb(52 53 65)"}
        w = "100vw"
        h = "100vh"
        zIndex={10}
        justifyContent={"center"}
        alignItems={"center"}
        top = {top}
        transitionDuration={"0.8s"}
        transitionTimingFunction={"ease-in-out"}
        onClick={() => setTop(-2000)}
        >
            <Box
            position={"relative"}
            width = "40vw"
            height = "30vh"
            display={"flex"}
            justifyContent={"center"}
            alignContent={"center"}
            onMouseEnter={() => {
                setLogo(splashImage.GenieLampOpen);
                setPos("10%");
                setRot("rotate(-10deg)");
            }}
            onMouseLeave={() => {
                setLogo(splashImage.GenieLamp)
                setPos("20%");
                setRot("rotate(0deg)");
            }}
            >
                <Image src={logo} w={"30vw"} 
                    transitionDuration = "0.3s"
                    transitionTimingFunction="ease-in-out"
                    zIndex={10}
                    transform={rot}
                    _hover={{
                        transitionDuration :"0.3s",
                        transitionTimingFunction:"ease-in-out",
                    }}
                    />
                <Image position="absolute"
                bottom={0} left={pos} zIndex={12}
                src = {Cloud} h = "20vh" 
                transitionDuration = "0.6s"
                transitionTimingFunction="ease-in-out"
                _hover={{
                    transitionDuration :"0.6s",
                    transitionTimingFunction:"ease-in-out",
                }}/>
                <Image position="absolute"
                bottom={0} right={pos} zIndex={8}
                src = {Cloud} h = "20vh" 
                transitionDuration = "0.9s"
                transitionTimingFunction="ease-in-out"
                _hover={{
                    transitionDuration :"0.9s",
                    transitionTimingFunction:"ease-in-out",
                }}/>
            </Box>

            <Box color={"#FFFFFF"} fontWeight={"600"}>
                <Text
                fontSize= "6xl"
                transitionDuration = "0.3s"
                transitionTimingFunction="ease-in-out"
                _hover={{
                    fontSize:"7xl",
                    transitionDuration :"0.3s",
                    transitionTimingFunction:"ease-in-out",
                }}>
                    ImaGenieKelar Bot
                </Text>
                <Text fontSize={"xl"}
                transitionDuration = "0.3s"
                transitionTimingFunction="ease-in-out"
                _hover={{
                    fontSize:"2xl",
                    transitionDuration :"0.3s",
                    transitionTimingFunction:"ease-in-out",
                }}> 
                    The gateway to cast all your problems away
                </Text>
            </Box>

        </Box>

    );
}

export default SplashScreen;