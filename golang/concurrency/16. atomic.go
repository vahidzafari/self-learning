package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// An atomic operation is an operation that is completed in a single step relative
// to other threads or, in this case, to other goroutines. This means that an atomic
// operation cannot be interrupted in the middle of it. The Go Standard library offers
// the atomic package, which, in some simple cases, can help you to avoid using a
// mutex. With the atomic package, you can have atomic counters accessed by multiple
// goroutines without synchronization issues and without worrying about race
// conditions.

type atomCounter struct {
	val int64
}

// This is a helper function that returns the current value of an int64 atomic variable
// using atomic.LoadInt64().
func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	X := 100
	Y := 4
	var waitGroup sync.WaitGroup
	counter := atomCounter{}
	for i := 0; i < X; i++ {

		waitGroup.Add(1)
		go func(no int) {
			defer waitGroup.Done()
			for i := 0; i < Y; i++ {
				atomic.AddInt64(&counter.val, 1)
			}

		}(i)
	}
	waitGroup.Wait()
	fmt.Println(counter.Value())
}
