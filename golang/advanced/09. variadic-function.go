package main

import (
	"fmt"
	"os"
)

// The general ideas and rules behind variadic functions are as follows:
// - Variadic functions use the pack operator, which consists of a ..., followed by
// a data type. So, for a variadic function to accept a variable number of int
// values, the pack operator should be ...int.
// - The pack operator can only be used once in any given function.
// - The variable that holds the pack operation is a slice and, therefore, is accessed
// as a slice inside the variadic function.
// - The variable name that is related to the pack operator is always last in the list
// of function parameters.
// - When calling a variadic function, you should put a list of values separated
// by , in the place of the variable with the pack operator or a slice with the
// unpack operator.

// The pack operator can also be used with an empty interface. In fact, most functions
// in the fmt package use ...interface{} to accept a variable number of arguments of
// all data types. You can find the source code of the latest implementation of fmt at
// https://golang.org/src/fmt/.

func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)
	sum := float64(0)
	for _, a := range s {
		sum = sum + a
	}
	s[0] = -1000
	return sum
}

func everything(input ...interface{}) {
	fmt.Println(input...)
}

func main() {
	sum := addFloats("Adding numbers...", 1.1, 2.12, 3.14, 4, 5, -1,
		10)
	fmt.Println("Sum:", sum)
	s := []float64{1.1, 2.12, 3.14}

	sum = addFloats("Adding numbers...", s...)
	fmt.Println("Sum:", sum)
	everything(s)

	// Cannot directly pass []string as []interface{}
	// You have to convert it first!
	empty := make([]interface{}, len(os.Args[1:]))
	for i, v := range os.Args[1:] {
		empty[i] = v
	}
	everything(empty...)

	arguments := os.Args[1:]
	empty = make([]interface{}, len(arguments))
	for i := range arguments {
		empty[i] = arguments[i]
	}

	everything(empty...)
	// This will work!
	str := []string{"One", "Two", "Three"}
	everything(str, str, str)
}
