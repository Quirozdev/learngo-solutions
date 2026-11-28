// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {
	const maxTurns = 5

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a value")
		return
	}

	guess, err := strconv.Atoi(args[1])

	if err != nil || guess <= 0 {
		fmt.Println("Please provide a valid integer value larger than 0")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		cpuguess := rand.Intn(guess) + 1

		if guess == cpuguess {
			if turn == 1 {
				fmt.Println("You won on the first turn!!! WOWWW")
			} else {
				fmt.Println("You won!!!")
			}
			return
		}
	}

	fmt.Println("You lost! :(")
}
