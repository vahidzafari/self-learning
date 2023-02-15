package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

// The main purpose of the context package is to define the Context type and support
// cancellation. Yes, you heard that right; there are times when, for some reason, you
// want to abandon what you are doing. However, it would be very helpful to be able
// to include some extra information about your cancellation decisions. The context
// package allows you to do exactly that.

// The Context type is an interface with four methods named Deadline(), Done(),
// Err(), and Value(). The good news is that you do not need to implement all of these
// functions of the Context interface—you just need to modify a Context variable using
// methods such as context.WithCancel(), context.WithDeadline(), and context.
// WithTimeout()

// The f1() function creates and executes a goroutine. The time.Sleep() call simulates
// the time it would take a real goroutine to do its job. In this case it is 4 seconds, but
// you can put any time period you want. If the c1 context calls the Done() function in
// less than 4 seconds, the goroutine will not have enough time to finish.
func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	// The WithCancel() method returns a copy of parent context with a new Done channel.
	// Notice that the cancel variable, which is a function, is one of the return values of
	// context.CancelFunc(). The context.WithCancel() function uses an existing Context
	// and creates a child with cancellation. The context.WithCancel() function also
	// returns a Done channel that can be closed, either when the cancel() function is called,
	// as shown in the preceding code, or when the Done channel of the parent context is
	// closed.

	go func() {
		time.Sleep(4 * time.Second)
		cancel()

	}()

	select {
	case <-c1.Done():
		fmt.Println("f1() Done:", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}
	return
}

// The cancel variable in f2() comes from context.WithTimeout(), which requires two
// parameters: a Context parameter and a time.Duration parameter. When the timeout
// period expires the cancel() function is called automatically.
func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()
	select {
	case <-c2.Done():
		fmt.Println("f2() Done:", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

// The cancel variable in f3() comes from context.WithDeadline(). context.
// WithDeadline() requires two parameters: a Context variable and a time in the future
// that signifies the deadline of the operation. When the deadline passes, the cancel()
// function is called automatically.
func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()
	select {
	case <-c3.Done():
		fmt.Println("f3() Done:", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a delay!")
		return
	}
	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delay:", delay)
	f1(delay)
	f2(delay)
	f3(delay)
}
