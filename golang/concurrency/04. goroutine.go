package main

import (
	"fmt"
	"time"
)

// You can define, create, and execute a new goroutine using the go keyword followed
// by a function name or an anonymous function. The go keyword makes the function
// call return immediately, while the function starts running in the background as a
// goroutine and the rest of the program continues its execution. You cannot control
// or make any assumptions about the order in which your goroutines are going to be
// executed because that depends on the scheduler of the OS, the Go scheduler, and
// the load of the OS.

func main() {
	count := 10
	fmt.Printf("Going to create %d goroutines.\n", count)
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
