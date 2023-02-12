package main

import (
	"fmt"
	"sync"
)

// A channel is a communication mechanism that, among other things, allows
// goroutines to exchange data. Firstly, each channel allows the exchange of a particular
// data type, which is also called the element type of the channel, and secondly, for
// a channel to operate properly, you need someone to receive what is sent via the
// channel. You should declare a new channel using make() and the chan keyword
// (make(chan int)), and you can close a channel using the close() function. You can
// declare the size of a channel by writing something like make(chan int, 1).

// A pipeline is a virtual method for connecting goroutines and channels so that the
// output of one goroutine becomes the input of another goroutine using channels
// to transfer your data. One of the benefits that you get from using pipelines is that
// there is a constant data flow in your program, as no goroutine or channel has to wait
// for everything to be completed in order to start their execution. Additionally, you
// use fewer variables and therefore less memory space because you do not have to
// save everything as a variable. Finally, the use of pipelines simplifies the design of
// the program and improves its maintainability.

func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

func printer(ch chan bool) {
	ch <- true
}

func main() {
	// This channel is buffered with a size of 1. This means that as soon as we fill that
	// buffer, we can close the channel and the goroutine is going to continue its execution
	// and return. A channel that is unbuffered has a different behavior: when you try to
	// send a value to that channel, it blocks forever because it is waiting for someone to
	// fetch that value. In this case, we definitely want a buffered channel in order to avoid
	// any blocking.
	c := make(chan int, 1)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)

	fmt.Println("Read:", <-c)
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

	waitGroup.Wait()

	var ch chan bool = make(chan bool)
	for i := 0; i < 5; i++ {
		go printer(ch)
	}

	// Range on channels
	// IMPORTANT: As the channel c is not closed,
	// the range loop does not exit by its own.
	n := 0
	for i := range ch {
		fmt.Println(i)
		if i == true {
			n++
		}
		if n > 2 {
			fmt.Println("n:", n)
			close(ch)
			break
		}
	}

	// When trying to read from a closed channel, we get the zero value of its data type
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
