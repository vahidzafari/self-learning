package main

import "fmt"

// The defer keyword postpones the execution of a function until the surrounding function returns.
// Usually, defer is used in file I/O operations to keep the function call that closes an
// opened file close to the call that opened it, so that you do not have to remember to
// close a file that you have opened just before the function exits.

// It is very important to remember that deferred functions are executed in last in, first
// out (LIFO) order after the surrounding function has been returned. Putting it simply,
// this means that if you defer function f1() first, function f2() second, and function
// f3() third in the same surrounding function, then when the surrounding function is
// about to return, function f3() will be executed first, function f2() will be executed
// second, and function f1() will be the last one to get executed.

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
