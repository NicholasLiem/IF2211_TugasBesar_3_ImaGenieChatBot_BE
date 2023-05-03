package extra

import (
	"math/rand"
	"time"
	"strings"
)

var RPSLIST = []string{"Batu", "Gunting", "Kertas"}

type RPSGame struct {
	userInput string
	programOutput string
	status int
	valid bool
	message string
}

func (r *RPSGame) PlayGame(s string) {
	r.Reset()
	r.userInput = s
	// Using time for randomize
	rand.Seed(time.Now().UnixNano()) 

	r.CheckUserInput()
	if (r.valid){
		var num int = rand.Intn(3)
		r.programOutput = RPSLIST[num]
		r.UpdateGameStatus()
		if (r.status == 1){
			r.message = "Program uses <"+r.programOutput+">. Luck is on your side, you have won the game."
		} else if (r.status == 2){
			r.message = "Program uses <"+r.programOutput+">. The game ends with a draw."
		} else {
			r.message = "Program uses <"+r.programOutput+">. You might have a skill issue, you lost the game."
		}
	} else {
		r.status = -1
		r.message = "Input \""+r.userInput+"\" is invalid. Please insert a correct one."
	}
}

func (r *RPSGame) CheckUserInput() {
	for _, s := range RPSLIST {
		if (strings.ToLower(s) == strings.ToLower(r.userInput)){
			r.valid = true
			return
		}
	}
	r.valid = false
}

func (r *RPSGame) Reset(){
	r.userInput = ""
	r.programOutput = ""
	r.message = ""
	r.valid = true
	r.status = 0
}

func (r *RPSGame) UpdateGameStatus() {
	// 0 => User loses
	// 1 => User wins
	// 2 => The game ends with a draw
	userIdx := 0
	programIdx := 0
	for i, s := range RPSLIST{
		if (strings.ToLower(r.programOutput) == strings.ToLower(s)){
			programIdx = i
		}
		if (strings.ToLower(r.userInput) == strings.ToLower(s)){
			userIdx = i
		}
	}

	if (userIdx == programIdx){
		r.status = 2
		return
	}
	if (userIdx == 2){
		if (programIdx == 0){ r.status = 1; return} else { r.status = 0; return}
	} else if (userIdx == 0){
		if (programIdx == 1){ r.status = 1; return} else { r.status = 0; return}
	} else {
		if (programIdx == 2){ r.status = 1; return} else { r.status = 0; return}
	}
}

func (r *RPSGame) GetMessage() string {
	return r.message
}

