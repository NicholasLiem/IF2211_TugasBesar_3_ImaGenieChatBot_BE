package main

import (
	"fmt"
	"github.com/NicholasLiem/Tubes3_Testing/calculator"
)

func main () {
	c := &calculator.Calculator{}

	strArr := [12] string {
		"6+3-5*7",
		"(8*4)/(4-2)",
		"(/3+4)", 
		"(9*3)^(1/3)*2",
		"(4+7+.",
		"((9+7)))",
		"()()(5)",
		"9.73432-2.345*2",
		"6+7/0",
		"6..4+3",
		"   5+6  *9",   // Ignoring whitespace test
		"(4+5)2",
		}

	for _,e := range strArr {
		c.InsertInput(e)
		fmt.Println("Input: ", c.GetInput())
		c.Calculate()
		if(c.IsValid()){
			fmt.Println("Solution: ", c.GetSolution())
		} else {
			fmt.Println(c.GetErrorMessage())
		}
	}



	// ns := &calculator.NumberStack{}
	// ns.Push(5)
	// ns.Push(65354.678)
	// ns.Pop()
	// ns.Display()

	// os := &calculator.OperatorStack{}
	// os.Push('v')
	// os.Push('g')
	// os.Push('t')
	// os.Pop()
	// os.Display()

}
