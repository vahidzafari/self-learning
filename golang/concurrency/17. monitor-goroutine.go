package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// Go comes with built-in synchronization features that allow a single goroutine to
// own a shared piece of data. This means that other goroutines must send messages
// to this single goroutine that owns the shared data, which prevents the corruption of
// the data. Such a goroutine is called a monitor goroutine. In Go terminology, this is
// sharing by communicating instead of communicating by sharing.

// When you have a read request, the read() function attempts to read
// from the readValue channel, which is controlled by the monitor() function.
// This returns the current value of the value variable. On the other hand, when you
// want to change the stored value, you call set(). This writes to the writeValue
// channel, which is also handled by the same select statement. As a result, no one
// can deal with the shared variable without using the monitor() function, which is
// in charge.

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}
func read() int {
	return <-readValue
}

// The monitor() function contains the logic of the program with the endless for loop
// and the select statement. The first case receives data from the writeValue channel,
// sets the value variable accordingly, and prints that new value. The second case sends
// the value of the value variable to the readValue channel. As all traffic goes through
// monitor() and its select block, there is no way to have a race condition because
// there is a single instance of monitor() running.
func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give an integer!")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Going to create %d random numbers.\n", n)
	rand.Seed(time.Now().Unix())
	go monitor()

	var wg sync.WaitGroup
	for r := 0; r < n; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}

	wg.Wait()
	fmt.Printf("\nLast value: %d\n", read())
}
