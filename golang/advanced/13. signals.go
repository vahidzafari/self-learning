package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// UNIX signals offer a very handy way of interacting asynchronously with your
// applications. However, UNIX signal handling requires the use of Go channels that
// are used exclusively for this task.

// A goroutine is the smallest executable Go entity. In order to create a new goroutine
// you have to use the go keyword followed by a predefined function or an anonymous
// function—the methods are equivalent. A channel in Go is a mechanism that among
// other things allows goroutines to communicate and exchange data.

// There exists a dedicated channel that receives all signals, as defined by the signal.
// Notify() function. Go channels can have a capacity—the capacity of this particular
// channel is 1 in order to be able to receive and keep one signal at a time. This makes
// perfect sense as a signal can terminate a program and there is no need to try to
// handle another signal at the same time. There is usually an anonymous function
// that is executed as a goroutine and performs the signal handling and nothing else.
// The main task of that goroutine is to listen to the channel for data. Once a signal is
// received, it is sent to that channel, read by the goroutine, and stored into a variable—
// at this point the channel can receive more signals. That variable is processed by a
// switch statement.

func handleSignal(sig os.Signal) {
	fmt.Println("handleSignal() Caught:", sig)
}

func main() {
	fmt.Printf("Process ID: %d\n", os.Getpid())
	sigs := make(chan os.Signal, 1)

	// If you want to handle a limited number of signals, instead of all of them, you should
	// replace the signal.Notify(sigs) statement with the next statement:
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGINFO)
	signal.Notify(sigs)
	start := time.Now()

	go func() {
		for {
			sig := <-sigs

			switch sig {
			case syscall.SIGINT:
				duration := time.Since(start)
				fmt.Println("Execution time:", duration)

			case syscall.SIGUSR1:
				handleSignal(sig)

				// do not use return here because the goroutine exits
				// but the time.Sleep() will continue to work!
				os.Exit(0)
			default:
				fmt.Println("Caught:", sig)
			}

		}
	}()

	for {
		fmt.Print("+")
		time.Sleep(10 * time.Second)
	}
}
