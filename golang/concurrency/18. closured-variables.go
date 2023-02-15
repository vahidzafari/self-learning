package main

import (
	"fmt"
	"time"
)

// We are going to talk about closured variables, which are variables
// inside closures, and the go statement. Notice that closured variables in goroutines are
// evaluated when the goroutine actually runs and when the go statement is executed
// in order to create a new goroutine. This means that closured variables are going to be
// replaced by their values when the Go scheduler decides to execute the relevant code.

// The program mostly prints the number 21, which is the last value of the variable of
// the for loop and not the other numbers. As i is a closured variable, it is evaluated
// at the time of execution. As the goroutines begin but wait for the Go scheduler to
// allow them to get executed, the for loop ends, so the value of i that is being used is
// 21. Lastly, the same issue also applies to Go channels, so be careful.

func main() {
	// Problem
	// for i := 0; i <= 20; i++ {
	// 	go func() {
	// 		fmt.Print(i, " ")
	// 	}()
	// }

	// Solution
	// for i := 0; i <= 20; i++ {
	// 	i := i
	// 	go func() {
	// 		fmt.Print(i, " ")
	// 		}()
	// 	}

	// Solution
	for i := 0; i <= 20; i++ {
		go func(x int) {
			fmt.Print(x, " ")
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println()
}
