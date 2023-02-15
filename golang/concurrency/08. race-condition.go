package main

import "fmt"

// A data race condition is a situation where two or more running elements, such as
// threads and goroutines, try to take control of or modify a shared resource or shared
// variable of a program. Strictly speaking, a data race occurs when two or more
// instructions access the same memory address, where at least one of them performs a
// write (change) operation. If all operations are read operations, then there is no race
// condition.
// Using the -race flag when running or building Go source files executes the Go race
// detector, which makes the compiler create a modified version of a typical executable
// file. This modified version can record all accesses to shared variables as well as all
// synchronization events that take place, including calls to sync.Mutex and sync.

// go run -race channels.go

func printer(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}

func main() {
	// This is an unbuffered channel
	var ch chan bool = make(chan bool)
	// Write 5 values to channel with a single goroutine
	go printer(ch, 5)
	// IMPORTANT: As the channel c is closed,
	// the range loop is going to exit on its own.
	for val := range ch {
		fmt.Print(val, " ")
	}
	fmt.Println()
	for i := 0; i < 15; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}
