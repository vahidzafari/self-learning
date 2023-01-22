package main

import (
	"fmt"
	"os"
	"strconv"
)

// Go supports the if/else and switch control structures.
// if statements use no parenthesis for embedding the conditions that need to be
// examined because Go does not use parentheses in general. As expected, if has
// support for else and else if statements.

// err := anyFunctionCall()
// if err != nil {
// // Do something if there is an error
// }

// The switch statement has two different forms. In the first form, the switch statement
// has an expression that is being evaluated, whereas in the second form, the switch
// statement has no expression to evaluate. In that case, expressions are evaluated in
// each case statement, which increases the flexibility of switch. The main benefit you
// get from switch is that when used properly, it simplifies complex and hard-to-read
// if-else blocks.

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}
	argument := os.Args[1]

	// With expression after switch
	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	default:
		fmt.Println("Value:", argument)
	}

	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to int:", argument)
		return
	}

	// No expression after switch
	switch {
	case value == 0:
		fmt.Println("Zero!")
	case value > 0:
		fmt.Println("Positive integer")
	case value < 0:
		fmt.Println("Negative integer")
	default:
		fmt.Println("This should not happen:", value)
	}
}
