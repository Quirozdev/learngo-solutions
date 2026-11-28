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
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

func main() {

	const maxTurns = 5

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a value")
		return
	}

	var guess int
	var err error
	var option string

	if len(args) > 2 {
		option = args[1]
		guess, err = strconv.Atoi(args[2])

	} else {
		guess, err = strconv.Atoi(args[1])
	}

	if err != nil || guess <= 0 {
		fmt.Println("Please provide a valid integer value larger than 0")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		cpuguess := rand.Intn(guess) + 1

		if option == "-v" {
			fmt.Printf("%d ", cpuguess)
		}

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
