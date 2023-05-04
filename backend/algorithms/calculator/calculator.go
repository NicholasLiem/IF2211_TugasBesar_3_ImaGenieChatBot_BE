package calculator

import (
	"fmt"
	"math"
	"strconv"
)

type Calculator struct {
	input string
	nDeque *NumberDeque
	oDeque *OperatorDeque
	lastIsNum bool
	currentVal float64
	commaVal int
	valid bool
	solution float64
	errorMessage string
	lastIsCP bool //CP = closing parenthesis
	lastIsMinusNum bool // For minus that is in the value of a num
	minusVal int 
}

func (c *Calculator) InsertInput(input string) {
	c.input = input
}



func (c *Calculator) ResetCalculator(){
	c.nDeque = &NumberDeque{}
	c.oDeque = &OperatorDeque{}
	c.nDeque.Reset()
	c.oDeque.Reset()
	c.lastIsNum = false
	c.currentVal = 0
	c.commaVal = 0
	c.valid = true
	c.solution = 0
	c.errorMessage = ""
	c.lastIsCP = false
	c.lastIsMinusNum = false
	c.minusVal = 0
}



func IsOperand(c rune) bool{
	arr :=[10] rune {rune('0'), rune('1'),rune('2'), rune('3'), rune('4'), rune('5'), rune('6'), rune('7'), rune('8'), rune('9')}
	for i := range arr {
		if c == arr[i] {
			return true
		}
	}
	return false
}

func IsOperator(c rune) bool {
	arr := [5] rune {rune('+'), rune('-'), rune('*'), rune('/'), rune('^')}
	for i := range arr {
		if c == arr[i] {
			return true
		}
	}
	return false
}

func IsExtra(c rune) bool {
	return c == rune('(') || c == rune(')') || c == rune('.')
}


func Precedence(c rune) int {
	if c == rune('^') {
		return 3
	} else if (c == rune('*') || c == rune('/')){
		return 2
	} else if (c == rune('+') || c == rune('-')){
		return 1
	} else {
		return -1
	}
}

func Operate(a float64, b rune, c float64) (float64, bool) {
	if (b == rune('+')){
		return a + c, true
	} else if (b == rune('-')){
		return a - c, true
	} else if (b == rune('*')){
		return a * c, true
	} else if (b == rune('/')){
		if (c != 0){
			return a / c, true
		} else {
			return 0, false
		}
	} else {
		return math.Pow(a,c), true
	}
}

func (c *Calculator) GetCurrentValToDeque() {
	c.commaVal = 0
	c.nDeque.InsertLast(c.currentVal)
	c.currentVal = 0
}

func (c *Calculator) SetInvalid(msg string){
	c.valid = false
	c.errorMessage = msg
}

func (c *Calculator) Calculate() {
	c.ResetCalculator()
	// Reading and calculate input
	for _, char := range c.input{
		// Number
		if (IsOperand(char)){
			if (!c.lastIsCP){
				str := string(char)
				digit,_ := strconv.Atoi(str)
				if (c.commaVal > 0){
					// Case if the digit is for after-comma value
					c.currentVal += float64(digit)/ math.Pow(10, float64(c.commaVal))
					c.commaVal++
				} else {
					// Case if the digit is not for after-comma value
					c.currentVal = c.currentVal*10 + float64(digit)
					if (c.lastIsMinusNum){
						c.currentVal *= math.Pow(-1, float64(c.minusVal))
						c.minusVal = 0
						c.lastIsMinusNum = false
					}
				}
				c.lastIsNum = true
				c.lastIsCP = false
			} else {
				c.SetInvalid("Invalid input, number after closing parenthesis detected.")
				break
			}
		} else if (IsExtra(char)){
			if (char == rune('.')){
				if (c.lastIsNum){
					if (c.commaVal > 0){
						c.SetInvalid("Invalid input, number with more than one comma detected.")
						break
					} else {
						c.commaVal = 1
						c.lastIsNum = false
						c.lastIsCP = false
					}
				} else {
					c.SetInvalid("Invalid input, non-number before comma detected.")
					break
				}
			} else if (char == rune('(')){
				if (c.lastIsNum){
					if (c.lastIsCP){
						c.SetInvalid("Invalid input, closing parenthesis right before opening parenthesis detected.")
						break
					} else {
						c.SetInvalid("Invalid input, number before opening parenthesis detected.")
						break
					}
				} else {
					c.commaVal = 0
					c.oDeque.InsertLast(char)
					c.lastIsNum = false
					c.lastIsCP = false
				}
			} else if (char == rune(')')){
				if (!c.lastIsNum){
					// () is still allowable
					if !c.oDeque.IsEmpty() && c.oDeque.Top() == rune('('){
						c.SetInvalid("Invalid input, empty parenthesis detected.")
						break
					}  else {
						c.SetInvalid("Invalid input, operator before closing parenthesis detected.")
						break
					}
				} else {
					// Case handling for cases like (5), (((7.568)))
					// If there is operation between "(" and ")", then these steps will be skipped
					c.GetCurrentValToDeque()
					for (!c.oDeque.IsEmpty() && c.oDeque.Top() != rune('(')) {
						int1 := c.nDeque.DeleteLast()
						int2 := c.nDeque.DeleteLast()
						op := c.oDeque.DeleteLast()
						res, valid := Operate(int2, op, int1)
						if (!valid) {
							c.SetInvalid("Invalid input, division by 0 detected.")
							break
						} else {
							// Works as if the all between parenthesis is somesort of number
							c.nDeque.InsertLast(res)
						}
					}
					if (c.oDeque.IsEmpty()){
						c.SetInvalid("Invalid input, extra closing parenthesis detected.")
						break
					} else {
						c.oDeque.DeleteLast()
						c.currentVal = c.nDeque.DeleteLast()
					}

					// Ini memang sengaja dibikin tetap c.lastIsNum = true
					// c.lastIsNum = false
					c.lastIsCP = true
				}
			}
		} else if (IsOperator((char))){
			if (char == rune('-') && !c.lastIsNum){
				c.lastIsMinusNum = true
				c.minusVal++
			} else if (c.lastIsNum){
				c.GetCurrentValToDeque()
				if (c.oDeque.IsEmpty()){
					c.oDeque.InsertLast(char)
				} else {
					if (Precedence(char) >= Precedence(c.oDeque.Top())){
						c.oDeque.InsertLast(char)
					} else {
						for (!c.oDeque.IsEmpty() && (Precedence(char) < Precedence(c.oDeque.Top())) ){
							int1 := c.nDeque.DeleteLast()
							int2 := c.nDeque.DeleteLast()
							op := c.oDeque.DeleteLast()
							res, valid := Operate(int2, op, int1)
							if (!valid) {
								c.SetInvalid("Invalid input, division by 0 detected.")
								break
							} else {
								c.nDeque.InsertLast(res)
							}
						}
						c.oDeque.InsertLast(char)
					}
				}
				c.lastIsNum = false
				c.lastIsCP = false
			} else {
				c.SetInvalid("Invalid input, non-number before operator detected.")
				break
			}
		} else if (char == rune(' ')){
			continue
		} else {
			c.SetInvalid("Invalid input, invalid symbol detected.")
			break
		}

		// Only for testing
		// fmt.Println("curChar: ", string(char))
		// fmt.Print("nDeque: "); c.nDeque.Display()
		// fmt.Print("oDeque: "); c.oDeque.Display()
		// fmt.Print("currentVal: "); c.DisplayCurrentVal()
	}

	// Calculate the rest after the reading input is finished
	if (c.valid){
		// No need to progress further if it is not valid
		if (!c.lastIsNum){
			c.SetInvalid("Invalid input, equation is ended by operator detected.")
		} else {
			c.GetCurrentValToDeque()
			// fmt.Print("nDeque: "); c.nDeque.Display()
			// fmt.Print("oDeque: "); c.oDeque.Display()
			// fmt.Print("currentVal: "); c.DisplayCurrentVal()
			for (!c.oDeque.IsEmpty()){
				if (c.nDeque.nEff >= 2){
					if (!c.IsAllSamePrecedence()){
						int1 := c.nDeque.DeleteLast()
						int2 := c.nDeque.DeleteLast()
						op := c.oDeque.DeleteLast()
						res, valid := Operate(int2, op, int1)
						if (!valid) {
							c.SetInvalid("Invalid input, division by 0 detected.")
							break
						} else {
							c.nDeque.InsertLast(res)
						}
					} else {
						int1 := c.nDeque.DeleteFirst()
						int2 := c.nDeque.DeleteFirst()
						op := c.oDeque.DeleteFirst()
						res, valid := Operate(int1, op, int2)
						if (!valid) {
							c.SetInvalid("Invalid input, division by 0 detected.")
							break
						} else {
							c.nDeque.InsertFirst(res)
						}
					}

				} else {
					c.SetInvalid("Invalid input, extra opening parenthesis detected.")
					break
				}
			}
		}
	
		// Save the result if still valid
		if (c.valid){
			c.solution = c.nDeque.Top()
		}
	}
	
}

func (c *Calculator) GetInput() string {
	return c.input
}

func (c *Calculator) IsValid() bool {
	return c.valid == true
}

func (c *Calculator) GetErrorMessage() string {
	return c.errorMessage
}

func (c *Calculator) GetSolution() float64 {
	return c.solution
}

func (c *Calculator) DisplayCurrentVal(){
	fmt.Println(c.currentVal)
}

func (c *Calculator) IsAllSamePrecedence() bool {
	prec := Precedence(c.oDeque.Top())
	for _, o := range c.oDeque.buffer {
		if (Precedence(o) != prec){
			return false
		}
	}
	return true
}