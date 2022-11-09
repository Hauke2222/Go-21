// 21 is a two player game, the game is played by choosing a number (1, 2, or 3) to be added to the running total.
// The game is won by the player whose chosen number causes the running total to reach exactly 21.
// The running total starts at zero. One player will be the computer.
// Players alternate supplying a number to be added to the running total.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	userInput     int = 0
	computerInput int = 0
	count         int = 0
	min           int = 1
	max           int = 3
)

func main() {
	for {
		userTurn()
		computerTurn()
	}
}
func userTurn() {
	fmt.Println("User turn:")

	for {
		fmt.Println("Type a number from 1 to 3: ")
		fmt.Scanln(&userInput)
		if (userInput-min)*(userInput-max) <= 0 && count+computerInput <= 21 {
			count += userInput
			fmt.Println("Your number:", userInput)
			fmt.Println("The total is now:", count)
			fmt.Println("---------------------------------")
			break
		} else if !((userInput-min)*(userInput-max) <= 0) {
			fmt.Println("Invalid number, try again")
		}
	}
	checkForWin("User")
}
func computerTurn() {
	for {
		rand.Seed(time.Now().UnixNano())
		computerInput = rand.Intn(max-min+1) + min

		if count+computerInput > 21 {
			computerInput = rand.Intn(max-min+1) + min
		} else {
			count += computerInput
			fmt.Println("Computer turn:")
			fmt.Println("The computer choose number:", computerInput)
			fmt.Println("The total is now:", count)
			fmt.Println("---------------------------------")
			break
		}
	}
	checkForWin("Computer")
}
func checkForWin(player string) {
	if count == 21 {
		fmt.Println("The", player, "has won.")
		os.Exit(1)
	}
}
