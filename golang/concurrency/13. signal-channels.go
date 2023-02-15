package main

import (
	"fmt"
	"sync"
	"time"
)

// A signal channel is one that is used just for signaling. Put simply, you can use
// a signal channel when you want to inform another goroutine about something.
// Signal channels should not be used for data transferring.

// Function A() is going to be blocked until channel a, which is passed as a parameter,
// is closed. Just before it ends, it closes channel b, which is passed as a parameter. This
// is going to unblock the next goroutine, which is going to be function B().

var wg sync.WaitGroup

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	time.Sleep(3 * time.Second)
	close(b)
}

func C(a, b chan struct{}) {
	<-a
	fmt.Println("C()!")
	close(b)
}

// This is the last function that is going to be executed. Therefore, although it is blocked,
// it does not close any channels before exiting. Additionally, being the last function
// means that it can be executed more than once, which is not true for functions A(), B()
// and C() because a channel can be closed only once.
func D(a chan struct{}) {
	<-a
	fmt.Println("D()!")
	wg.Done()
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})
	w := make(chan struct{})

	wg.Add(1)
	go func() {
		D(w)
	}()

	wg.Add(1)
	go func() {
		D(w)
	}()

	go A(x, y)

	wg.Add(1)
	go func() {
		D(w)
	}()

	go C(z, w)
	go B(y, z)

	wg.Add(1)
	go func() {
		D(w)
	}()

	// This triggers the process
	close(x)

	wg.Wait()
}
