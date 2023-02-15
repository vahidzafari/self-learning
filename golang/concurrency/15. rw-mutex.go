package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.RWMutex is based on sync.Mutex with the necessary additions and improvements.
// So, you might ask, how does sync.RWMutex improve sync.Mutex? Although a single
// function is allowed to perform write operations with a sync. RWMutex mutex, you
// can have multiple readers owning a sync.RWMutex mutex—this means that read
// operations are usually faster with sync.RWMutex. However, there is one important
// detail that you should be aware of: until all of the readers of a sync.RWMutex
// mutex unlock that mutex, you cannot lock it for writing, which is the small price
// you have to pay for the performance improvement you get for allowing multiple readers.

// The functions that can help you to work with sync.RWMutex are RLock() and
// RUnlock(), which are used for locking and unlocking the mutex for reading
// purposes, respectively. The Lock() and Unlock() functions used in sync.Mutex
// should still be used when you want to lock and unlock a sync.RWMutex mutex for
// writing purposes. Finally, it should be apparent that you should not make changes to
// any shared variables inside an RLock() and RUnlock() block of code.

var Password *secret
var wg sync.WaitGroup

type secret struct {
	RWM      sync.RWMutex
	password string
}

func Change(pass string) {
	fmt.Println("Change() function")
	Password.RWM.Lock()

	fmt.Println("Change() Locked")
	time.Sleep(4 * time.Second)
	Password.password = pass
	Password.RWM.Unlock()

	fmt.Println("Change() UnLocked")
}

func show() {
	defer wg.Done()
	Password.RWM.RLock()
	fmt.Println("Show function locked!")
	time.Sleep(2 * time.Second)
	fmt.Println("Pass value:", Password.password)
	defer Password.RWM.RUnlock()
}

func main() {
	Password = &secret{password: "myPass"}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go show()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("123456")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("54321")
	}()

	wg.Wait()

	// Direct access to Password.password
	fmt.Println("Current password value:", Password.password)
}
