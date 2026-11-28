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
	"math"
	"math/rand"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func main() {
	const maxTurns = 5

	args := os.Args

	if len(args) < 3 {
		fmt.Println("Please provide two guesses")
		return
	}

	firstguess, err1 := strconv.Atoi(args[1])
	secondguess, err2 := strconv.Atoi(args[2])

	if err1 != nil || err2 != nil || firstguess <= 0 || secondguess <= 0 {
		fmt.Println("Please provide valid guesses larger than 0")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		cpuguess := rand.Intn(int(math.Max(float64(firstguess), float64(secondguess)))) + 1

		if firstguess != cpuguess && secondguess != cpuguess {
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
