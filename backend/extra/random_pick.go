package extra

import (
	"math/rand"
	"strings"
	"time"
	"strconv"
	"fmt"
)

type RandomPick struct {
	amount int
	buffer []string
	message string
	valid bool
}

func (r *RandomPick) Pick(n string, s string) {
	r.Reset()
	nPick, err := strconv.Atoi(n)
	if (err != nil || nPick < 0){
		r.valid = false
		r.message = "The amount to choose is invalid. Please insert a correct one."
		return
	}
	r.amount = nPick
	r.buffer = strings.Split(s, " ")
	if (r.amount > len(r.buffer)){
		r.valid = false
		r.message = "The amount to choose is greater than the amount of choices. Please make sure to insert the right one."
		return
	}
	fmt.Println(r.amount, r.buffer)
	r.GetOutput()
	r.valid = true

}

func (r *RandomPick) GetOutput() {
	if (r.amount == 0){
		r.message = "Program does not choose anything."
	} else if (r.amount == len(r.buffer)){
		r.message = "Program chooses every given options."
	} else {
		var temp = []int{}
		for  len(temp) < r.amount{
			rand.Seed(time.Now().UnixNano())
			num := rand.Intn(len(r.buffer))
			if (!IsInArray(num, temp)){
				temp = append(temp, num)
			}

		}
		r.message = "From "+strconv.Itoa(len(r.buffer))+ " given choises. Program decides to choose "
		for i, s := range temp{
			tempStr := ""
			if (i == len(temp)-1){
				tempStr = "<"+ r.buffer[s]+">."
			} else {
				tempStr = "<"+r.buffer[s]+">, "
			}
			r.message +=  tempStr
		}

	}
}

func IsInArray(n int, arr []int) bool {
	for _, e := range arr {
		if (n == e){
			return true;
		}
	}
	return false;
}

func (r *RandomPick) GetMessage() string{
	return r.message
}

func (r *RandomPick) Reset() {
	r.amount = 0
	r.buffer = []string{}
	r.message = ""
	r.valid = true
}