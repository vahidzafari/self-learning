package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// A mutex variable, which is an abbreviation of mutual exclusion variable, is mainly used
// for thread synchronization and for protecting shared data when multiple writes can
// occur at the same time. A mutex works like a buffered channel with a capacity of
// one, which allows at most one goroutine to access a shared variable at any given
// time. This means that there is no way for two or more goroutines to be able to update
// that variable simultaneously. Go offers the sync.Mutex and sync.RWMutex data types.

// A critical section of a concurrent program is the code that cannot be executed
// simultaneously by all processes, threads, or, in this case, goroutines. It is the code
// that needs to be protected by mutexes. All of the interesting work is done
// by the sync.Lock() and sync.Unlock() functions, which can lock and unlock a
// sync.Mutex variable, respectively. Locking a mutex means that nobody else can
// lock it until it has been released using the sync.Unlock() function.

var m sync.Mutex
var v1 int

func change(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1 == 10 {
		v1 = 0
		fmt.Print("* ")
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup

	fmt.Printf("%d ", read())
	for i := 0; i < numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf("-> %d", read())
		}(i)
	}

	wg.Wait()
	fmt.Printf("-> %d\n", read())
}
