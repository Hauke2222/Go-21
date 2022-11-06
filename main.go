// 21 is a two player game, the game is played by choosing a number (1, 2, or 3) to be added to the running total.
// The game is won by the player whose chosen number causes the running total to reach exactly 21.
// The running total starts at zero. One player will be the computer.
// Players alternate supplying a number to be added to the running total.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	userInput     int  = 0
	computerInput int  = 0
	count         int  = 0
	min           int  = 1
	max           int  = 3
	winner        bool = false
)

func main() {
	i := 10
	for {
		userTurn()
		computerTurn()
		fmt.Println(i)
		i++
		if winner == true {
			break
		}
	}

	computerTurn()
}
func userTurn() {
	fmt.Println("Type a number from 1 to 3: ")
	fmt.Scanln(&userInput)
	if (userInput-min)*(userInput-max) <= 0 {
		count += userInput
	}
	fmt.Println("Your number:", userInput)
	fmt.Println("The total is now:", count)

	checkForWin()
}
func computerTurn() {
	rand.Seed(time.Now().UnixNano())
	computerInput = rand.Intn(max-min+1) + min
	fmt.Println(computerInput)

	checkForWin()
}
func checkForWin() {
	if count == 21 {
		fmt.Println("Win")
		return
	}
}
