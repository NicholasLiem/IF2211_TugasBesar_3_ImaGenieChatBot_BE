package main

import (
	"fmt"
	"github.com/NicholasLiem/Tubes3_Testing/calculator"
)

func main () {
	c := &calculator.Calculator{}
	c.InsertInput("5+4+6")
	fmt.Println(c.GetInput())

}
