package main

import "fmt"

// When using a channel as a function parameter, you can specify its direction; that is,
// whether it is going to be used for sending or receiving data.  You will not be able
// to send data accidentally to a channel from which you should only receive data or
// receive data from a channel to which you should only be sending data. As a result,
// if you declare that a channel function parameter is going to be used for reading
// only and you try to write to it, you get an error message that will most likely save
// you from nasty bugs in the future.

func printer(ch chan<- bool) {
	ch <- true
}

// This function accepts a channel parameter that is available for writing only
func writeToChannel(c chan<- int, x int) {
	fmt.Println("1", x)
	c <- x
	fmt.Println("2", x)
}

// The channel parameter of this function is available for reading only
func f2(out <-chan int, in chan<- int) {
	x := <-out
	fmt.Println("Read (f2):", x)
	in <- x
	return
}
func main() {}
