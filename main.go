package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	fmt.Print("Enter a name: ")

	var name string
	var input string
	var userScore int
	var computerScore int
	var i int

	rounds := 3
	index := 0

	_, err := fmt.Scan(&name)

	if err != nil {
		fmt.Println("Error reading name", err)
	}

	fmt.Printf("Start the game as %s \n", name)

	for {

		for {
			fmt.Print("Choose between 1 or 0: \n")
			_, err := fmt.Scan(&input)

			if err != nil {
				fmt.Println("Error reading name", err)
			}

			i, _ := strconv.Atoi(input)

			if i <= 1 && i >= 0 {
				break
			}
		}

		myrand := random(0, 2)
		if myrand == i {
			userScore = userScore + 1
			fmt.Printf("You won the round. %s %d - %d Computer \n", name, userScore, computerScore)
		} else {
			computerScore = computerScore + 1
			fmt.Printf("You lost the round. %s %d - %d Computer \n", name, userScore, computerScore)
		}

		index = index + 1
		if rounds == index {
			if userScore > computerScore {
				fmt.Println("You won the game.")
			} else {
				fmt.Println("Game over!!!")
			}

			fmt.Printf("Computer Score: %d \n", computerScore)
			fmt.Printf("%s Score: %d \n", name, userScore)
			break
		}
	}

}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
