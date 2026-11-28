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
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
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

		if guess != cpuguess {
			continue
		}

		switch n := rand.Intn(3); n {
		case 0:
			fmt.Println("YOU WON")
		case 1:
			fmt.Println("YOU'RE AWESOME")
		default:
			fmt.Println("YOU ARE A WINNER!")
		}
		return
	}

	switch n := rand.Intn(3); n {
	case 0:
		fmt.Println("LOSER!")
	case 1:
		fmt.Println("YOU LOST. TRY AGAIN?")
	default:
		fmt.Println("YOU DIDN'T GUESSED IT CORRECTLY!")
	}
}
