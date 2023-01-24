package main

import (
	"fmt"
	"os"
	"strconv"
)

// Usually, user input is given in the form of command-line arguments to the executable
// file.  By default, command-line arguments in Go are stored in the os.Args slice. Go
// also offers  the flag package for parsing command-line arguments.
// The first command-line argument stored in the os.Args slice is always the name of
// the executable.

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var min, max float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		if i == 1 {
			min = n
			max = n
			continue
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
