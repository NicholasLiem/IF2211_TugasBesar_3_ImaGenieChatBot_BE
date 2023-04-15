package calculator


type NumberStack struct {
	buffer[] *float64
	nEff int
}

func (ns *NumberStack) Push(n float64){
	ns.buffer = append(ns.buffer, &n)
	ns.nEff++
}

func (ns *NumberStack) Pop() float64{
	res := ns.buffer[0]
	ns.buffer = ns.buffer[1:]
	ns.nEff--
	return *res
}

func (ns *NumberStack) IsEmpty() bool {
	return ns.nEff == 0
}

func (ns *NumberStack) Top() float64 {
	return *ns.buffer[ns.nEff-1]
}

func (ns *NumberStack) Reset() {
	for (!ns.IsEmpty()){
		ns.Pop()
	}
	ns.nEff = 0
}



type OperatorStack struct {
	buffer[] *rune
	nEff int
}

func (os *OperatorStack) Push(s rune){
	os.buffer = append(os.buffer, &s)
	os.nEff++
}

func (os *OperatorStack) Pop() rune {
	res := os.buffer[os.nEff-1]
	os.buffer = os.buffer[:os.nEff-1]
	os.nEff--
	return *res
}

func (os *OperatorStack) IsEmpty() bool {
	return os.nEff == 0
}

func (os *OperatorStack) Top() rune {
	return *os.buffer[os.nEff-1]
}

func (os *OperatorStack) Reset() {
	for (!os.IsEmpty()){
		os.Pop()
	}
	os.nEff = 0
}