package main

import (
	"fmt"
	"sync"
)

// It is not enough to create multiple goroutines—you also need to wait for them
// to finish before the main() function ends.
// The synchronization process begins by defining a sync.WaitGroup variable and
// using the Add(), Done() and Wait() methods.

// type WaitGroup struct {
// 	noCopy noCopy
// 	state1 [3]uint32
// }

// Each call to sync.Add() increases a counter in the state1 field, which is an array
// with three uint32 elements. Notice that it is really important to call sync.Add()
// before the go statement in order to prevent any race conditions.
// When each goroutine finishes its job, the sync.Done() function should be executed
// in order to decrease the same counter by one. Behind the scenes, sync.Done() runs
// a Add(-1) call. The Wait() method waits until that counter becomes 0 in order to
// return. The return of Wait() inside the main() function means that main() is going
// to return and the program ends.

func main() {
	count := 10
	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)

	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
