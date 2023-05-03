package test


import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/extra"
	"fmt"
)

func testRPS() {
	r := extra.RPSGame{}

	var test = []string{
		"Gunting",
		"Batu",
		"Kertas",
		"BATU",
		"KERTAS",
		"hehe",
	}

	for _, t := range test {
		r.PlayGame(t)
		fmt.Println(r.GetMessage())
	}
}
