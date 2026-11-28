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
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
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

	var maxValue int

	if guess < 10 {
		maxValue = 10
	} else {
		maxValue = guess
	}

	for turn := 1; turn <= maxTurns; turn++ {
		cpuguess := rand.Intn(maxValue) + 1

		if guess != cpuguess {
			continue
		}

		if turn == 1 {
			fmt.Println("You won on the first turn!!! WOWWW")
		} else {
			fmt.Println("You won!!!")
		}
		return
	}

	fmt.Println("You lost! :(")

}
