package test

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/extra"
	"fmt"
)

func testRandomPick (){
	r := extra.RandomPick{}

	var test = []string{
		"Satu Dua Tiga",
		"hehe huhu haha hihi",
		"Juan1 Juan2 Juan3",
		"Makan Minum Mandi Tidur Gawe",
		"SatuDoang",
		"Ninu Nana Nono",
	}
	
	var num = []string{
		"1",
		"2",
		"3",
		"0",
		"3",
		"loh",
	}

	for i  := range test {
		r.Pick(num[i], test[i])
		fmt.Println(r.GetMessage())
	}
}