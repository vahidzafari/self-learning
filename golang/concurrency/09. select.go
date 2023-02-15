package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// The select keyword is really important because it allows you to listen to multiple
// channels at the same time. A select block can have multiple cases and an optional
// default case, which mimics the switch statement. It is good for select blocks to
// have a timeout option just in case. Last, a select without any cases (select{})
// waits forever.

// A select statement is not evaluated sequentially, as all of its channels are examined
// simultaneously. If none of the channels in a select statement are ready, the select
// statement blocks (waits) until one of the channels is ready. If multiple channels of a
// select statement are ready, then the Go runtime makes a random selection from
// the set of these ready channels.

func gen(min, max int, createNumber chan int, end chan bool) {
	time.Sleep(time.Second)
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("Ended!")
			// return
		case <-time.After(4 * time.Second):
			fmt.Println("time.After()!")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		gen(0, 2*n, createNumber, end)
		wg.Done()
	}()

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	end <- true
	wg.Wait()
	fmt.Println("Exiting...")
}
