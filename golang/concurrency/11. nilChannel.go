package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// nil channels always block! Therefore, you should use them when you want that
// behavior on purpose!

// The send() function keeps sending random numbers to the c channel. Do not
// confuse channel c, which is a channel function parameter, with channel t.C, which is
// part of timer t—you can change the name of the c variable but not the name of the C
// field. When the time of timer t expires, the timer sends a value to the t.C channel.

// This triggers the execution of the relevant branch of the select statement, which
// assigns the value nil to channel c, prints the value of the sum variable and wg.Done()
// is executed, which is going to unblock wg.Wait() found in the main() function.
// Additionally, as c becomes nil, it stops/blocks send() from sending any data to it.

var wg sync.WaitGroup

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)
	for {
		select {
		case input := <-c:
			sum = sum + input
		case <-t.C:
			c = nil
			fmt.Println(sum)
			wg.Done()
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	rand.Seed(time.Now().Unix())
	wg.Add(1)
	go add(c)
	go send(c)
	wg.Wait()
}
