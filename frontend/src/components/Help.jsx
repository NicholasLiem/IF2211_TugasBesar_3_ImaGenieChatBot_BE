import {Container, Box, Text} from "@chakra-ui/react"
import React, {useState} from "react";
import { Palette } from "../assets/palette";

const InstructionList = [
    {
        "title" : "Fitur Penambahan Pertanyaan",
        "query" : "(Tambahkan | Add) pertanyaan <Pertanyaan> dengan jawaban <Jawaban>",
        "example" : "Tambahkan pertanyaan mengapa bumi berbentuk balok? dengan jawaban karena Patrick adalah kuda laut",
        "desc" : "Program menerima pasangan pertanyaan dan jawaban yang diinput user untuk dimasukkan ke dalam database.",
    },
    {
        "title" : "Fitur Pengubahan Jawaban ",
        "query" : "(Ubah | Update) pertanyaan <Pertanyaan> dengan jawaban <Jawaban>",
        "example" : "Ubah pertanyaan mengapa bumi berbentuk balok? dengan jawaban karena jerapah suka makan ayam geprek",
        "desc" : "Program mengubah jawaban dari pertanyaan tertentu yang terdapat pada database menjadi jawaban baru yang diinput oleh user.",
    },
    {
        "title" : "Fitur Penghapusan Pertanyaan",
        "query" : "(Hapus | Delete) pertanyaan <Pertanyaan>",
        "example" : "Hapus pertanyaan mengapa bumi berbentuk balok?",
        "desc" : "Program menghapus pasangan pertanyaan dan jawaban dari jenis pertanyaan yang diinput oleh user.",
    },
    {
        "title" : "Fitur Kalkulator",
        "query" : "(Hitunglah | Berapakah) <Input>",
        "example" : "Hitunglah 2+4*8, Berapakah (((8*7)+6)-5)",
        "desc" : "Program mengeluarkan hasil dari perhitungan yang diinput oleh user. Jika input hitungan invalid, program akan menampilkan pesan error.",
    },
    {
        "title" : "Fitur Tanggal",
        "query" : "Hari apakah tanggal <Input>",
        "example" : "Hari apakah tanggal 26/7/2023",
        "desc" : "Program mengeluarkan nama hari dari tanggal yang telah diinput oleh user. Jika tanggal tersebut invalid, program akan menampilkan pesan error.",
    },
    {
        "title" : "Fitur Permainan Suit",
        "query" : "Mainkan suit dengan (Batu | Gunting | Kertas)",
        "example" : "Mainkan suit dengan Gunting",
        "desc" : "Program akan bermain permainan suit dengan user. Hasil yang dikeluarkan program didapatkan dengan cara acak.",
    },
    {
        "title" : "Fitur Pemilihan Acak",
        "query" : "Pilih <Angka> dari <Input1> <Input2> <Input-n>",
        "example" : "Pilih 2 dari Nicholas Nathania Juan, Pilih 1 dari Haha Hihi Huhu",
        "desc" : "Program akan memilih satu dari berbagai pilihan yang diinput oleh user.",
    },
]

const InstructionInfo = ({title, query, example, desc}) => {
    return (
        <Container
            bgColor={"#FFFFFF"}
            padding={10}
            marginY = {10}
            minW={"70%"}
            borderRadius={15}
            borderWidth={"2px"}
            borderColor={"#404040"}
            shadow={"xl"}
            transitionDuration="0.2s"
            transitionTimingFunction="ease-in-out"
            _hover={{
                paddingY: "20",
                transitionDuration:"0.2s",
                transitionTimingFunction:"ease-in-out"
            }}>
            <Text fontSize={"2xl"} fontWeight={"bold"}>
                {title}
            </Text>
            <Text fontSize={"l"} fontWeight={500} fontStyle={"italic"}>
                {query}
            </Text>
            <Text>
                {"Contoh: "}{example}
            </Text>
            <Text marginTop={2}>
                {desc}
            </Text>
        </Container>
    );
}


const HelpPage = () => {

    const [layerPos, setLayerPos] = useState(-3000)
    const [boxPos, setBoxPos] = useState(-3000)
    const [text, setText] = useState("?")
    const [color, setColor] = useState(Palette.dark)

    return (
        <Container>
            <Box
                h={"30px"}
                w={"30px"}  
                bg={"transparent"}      
                position={"absolute"}
                top={"8"}
                right={"8"}
                color={color}
                bgColor={"transparent"}
                borderColor={color}
                borderRadius={"15px"}
                borderWidth={"3px"}
                fontWeight={"bold"}
                _hover={{
                    bgColor:"#FFFFFF",
                    color: "rgb(52 53 65)",
                }}
                zIndex={6}
                onClick={() => { 
                    color === Palette.white ? setColor(Palette.dark) : setColor(Palette.white);
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
            opacity={0.4}
            top={0}
            left={layerPos}
            zIndex={5}
            w={"100%"}
            h={"100%"}
            transitionDuration={"0.5s"}
            transitionTimingFunction={"ease-in-out"}
            onClick={() => { 
                color === Palette.white ? setColor(Palette.dark) : setColor(Palette.white);
                text === "?" ? setText("X") : setText("?");
                layerPos === 0 ? setLayerPos(-3000) : setLayerPos(0);
                boxPos === "20vw" ? setBoxPos(-3000) : setBoxPos("20vw");
            }}
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
                <Text color={"#2F2F3F"} fontSize={"4xl"} fontWeight={"bold"} marginBottom={4}>
                    Instructions
                </Text >
                <Container 
                bg={"#9898A8"}
                borderRadius={"5px"}
                overflowY={"scroll"}
                scrollBehavior={"smooth"}
                minW="100%"
                h={"60vh"}
                sx={{
                    "::-webkit-scrollbar": {
                        width: "8px",
                    },
                    "::-webkit-scrollbar-track": {
                        background: "rgb(68,70,84)",
                    },
                    "::-webkit-scrollbar-thumb": {
                        background: "#FFFFFF",
                    },
                    }}>
                {InstructionList.map((data) => {
                    return (
                        <InstructionInfo title = {data.title} query={data.query} example= {data.example} desc={data.desc} />
                    )
                })}
                </Container>
            </Box>
                
        </Container>
    );
}



export default HelpPage;