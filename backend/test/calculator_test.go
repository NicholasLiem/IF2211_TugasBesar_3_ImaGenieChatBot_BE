package test

import (
	// "fmt"
	// "github.com/NicholasLiem/Tubes3_ImagineKelar/calculator"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/date"
)

func main() {

	// Calculator Testing
	// c := &calculator.Calculator{}

	// strArr := [] string {
	// 	"6+3-5*7",
	// 	"(8*4)/(4-2)",
	// 	"(/3+4)",
	// 	"(9*3)^(1/3)*2",
	// 	"(4+7+.",
	// 	"((9+7)))",
	// 	"()()(5)",
	// 	"9.73432-2.345*2",
	// 	"6+7/0",
	// 	"6..4+3",
	// 	"   5+6  *9",   // Ignoring whitespace test
	// 	"(4+5)2",
	// 	"(((6.3)))",
	// 	"6(45+4)",
	// 	"4^5-4*7/3",
	// 	"((4))+5-7",
	// 	"(((2^10))",
	// 	"(45-4)(5*3)",
	// 	"(4-0.345)",
	// 	"0.000",
	// 	" 7+ 4 / 5*",
	// 	"+4-5*4",
	// 	}

	// for _,e := range strArr {
	// 	c.InsertInput(e)
	// 	fmt.Println("Input: ", c.GetInput())
	// 	c.Calculate()
	// 	if(c.IsValid()){
	// 		fmt.Println("Solution: ", c.GetSolution())
	// 	} else {
	// 		fmt.Println(c.GetErrorMessage())
	// 	}
	// }

	d := &date.Date{}

	strArr := []string{
		"09/10/2003",
		"01/10/2003",
		"07/03/2003",
		"09-10-2023",
		"9/18/2004",
		"30/02/1500",
		"29/02/2000",
		"29/02/1900",
		"1/1/1",
		"23/4/2023",
		"0/5/2012",
		"hehe/5/2012",
		"25/5/2017",
	}

	for _, s := range strArr {
		d.GetDayFromDate(s)
		d.DisplayDate()
	}

}
