package calculator

import (
	"fmt"
)


type NumberDeque struct {
	buffer[] float64
	nEff int
}

func (ns *NumberDeque) InsertLast(n float64){
	ns.buffer = append(ns.buffer, n)
	ns.nEff++
}

func (ns *NumberDeque) DeleteLast() float64{
	res := ns.buffer[ns.nEff-1]
	ns.buffer = ns.buffer[:ns.nEff-1]
	ns.nEff--
	return res
}

func (ns *NumberDeque) InsertFirst(n float64){
	ns.buffer =  append([]float64{n}, ns.buffer...)
	ns.nEff++
}

func (ns *NumberDeque) DeleteFirst() float64 {
	res := ns.buffer[0]
	ns.buffer = ns.buffer[1:]
	ns.nEff--
	return res
}

func (ns *NumberDeque) IsEmpty() bool {
	return ns.nEff == 0
}

func (ns *NumberDeque) Top() float64 {
	return ns.buffer[ns.nEff-1]
}

func (ns *NumberDeque) Reset() {
	for (!ns.IsEmpty()){
		ns.DeleteLast()
	}
	ns.nEff = 0
}

func (ns *NumberDeque) Display(){
	for i, n := range ns.buffer {
		if (i == ns.nEff-1){
			fmt.Println(n)
		} else {
			fmt.Print(n , " ")
		}
	}
}

func (ns *NumberDeque) GetNeff() int{
	return ns.nEff
}



type OperatorDeque struct {
	buffer[] rune
	nEff int
}

func (os *OperatorDeque) InsertLast(s rune){
	os.buffer = append(os.buffer, s)
	os.nEff++
}

func (os *OperatorDeque) DeleteLast() rune {
	res := os.buffer[os.nEff-1]
	os.buffer = os.buffer[:os.nEff-1]
	os.nEff--
	return res
}


func (os *OperatorDeque) InsertFirst(n rune){
	os.buffer =  append([]rune{n}, os.buffer...)
	os.nEff++
}

func (os *OperatorDeque) DeleteFirst() rune {
	res := os.buffer[0]
	os.buffer = os.buffer[1:]
	os.nEff--
	return res
}

func (os *OperatorDeque) IsEmpty() bool {
	return os.nEff == 0
}

func (os *OperatorDeque) Top() rune {
	return os.buffer[os.nEff-1]
}

func (os *OperatorDeque) Reset() {
	for (!os.IsEmpty()){
		os.DeleteLast()
	}
	os.nEff = 0
}

func (os *OperatorDeque) Display(){
	for i, n := range os.buffer {
		if (i == os.nEff-1){
			fmt.Println(string(n))
		} else {
			fmt.Print(string(n) , " ")
		}
	}
}

func (os *OperatorDeque) GetNeff() int{
	return os.nEff
}