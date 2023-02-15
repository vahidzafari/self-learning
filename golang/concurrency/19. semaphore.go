package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/sync/semaphore"
)

// A semaphore is a construct that can limit or control the access to
// a shared resource. As we are talking about Go, a semaphore can limit the access of
// goroutines to a shared resource but originally, semaphores were used for limiting
// access to threads. Semaphores can have weights that limit the number of threads or
// goroutines that can have access to a resource.
// The process is supported via the Acquire() and Release() methods, which are
// defined as follows:
// func (s *Weighted) Acquire(ctx context.Context, n int64) error
// func (s *Weighted) Release(n int64)
// The second parameter of Acquire() defines the weight of the semaphore.

// Maximum number of goroutines
var Workers = 4
var sem = semaphore.NewWeighted(int64(Workers))

func worker(n int) int {
	square := n * n
	time.Sleep(time.Second)
	return square
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need #jobs!")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Where to store the results
	var results = make([]int, nJobs)

	// Needed by Acquire()
	ctx := context.TODO()

	// 	In this part, we try to acquire the semaphore as many times as the number of jobs
	// defined by nJobs. If nJobs is bigger than Workers, then the Acquire() call is going to
	// block and wait for Release() calls in order to unblock.
	for i := range results {
		err = sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Println("Cannot acquire semaphore:", err)
			break
		}

		go func(i int) {
			defer sem.Release(1)
			temp := worker(i)
			// No race conditions here - each goroutine writes
			// to a different slice element
			results[i] = temp
		}(i)
	}

	// Acquire all of the tokens
	// This is similar to Wait()
	// It blocks until all workers have finished
	err = sem.Acquire(ctx, int64(Workers))
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range results {
		fmt.Println(k, "->", v)
	}
}
