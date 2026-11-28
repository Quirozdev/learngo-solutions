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
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Printf("Please provide min and max.\n")
		return
	}

	min, err := strconv.Atoi(args[1])
	max, err2 := strconv.Atoi(args[2])

	if err != nil || err2 != nil {
		fmt.Println("wrong numbers")
		return
	}

	if min >= max {
		fmt.Printf("max (%d) must be larger than min (%d)\n", max, min)
		return
	}

	sum, i := 0, 0

	for {
		if i >= max {
			break
		}

		i++

		if i%2 != 0 {
			continue
		}

		sum += i
		fmt.Print(i)

		if i < max {
			fmt.Printf(" + ")
		}

		i++
	}

	fmt.Printf(" = %d\n", sum)
}
