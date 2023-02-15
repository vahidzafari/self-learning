package main

import "fmt"

// These channels allow us to put jobs in a queue quickly in order to be able to deal
// with more requests and process requests later on. Moreover, you can use buffered
// channels as semaphores in order to limit the throughput of your application.

// The presented technique works as follows: all incoming requests are forwarded to
// a channel, which processes them one by one. When the channel is done processing
// a request, it sends a message to the original caller saying that it is ready to process
// a new one. So, the capacity of the buffer of the channel restricts the number of
// simultaneous requests that it can keep.

func main() {
	// numbers cannot store more than 5 integers
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		// This is where the processing takes place
		case numbers <- i * i:
			fmt.Println("About to process", i)
		default:
			fmt.Print("No space for ", i, " ")
		}
	}
	fmt.Println()

	for {
		select {
		case num := <-numbers:
			fmt.Print("*", num, " ")
		default:
			fmt.Println("Nothing left to read!")
			return
		}
	}
}
