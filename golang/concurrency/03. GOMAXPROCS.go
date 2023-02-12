package main

import (
	"fmt"
	"runtime"
)

// The GOMAXPROCS environment variable allows you to set the number of OS threads
// (CPUs) that can execute user-level Go code simultaneously. Starting with Go version
// 1.5, the default value of GOMAXPROCS should be the number of logical cores available in
// your machine. There is also the runtime.GOMAXPROCS() function, which allows you to
// set and get the value of GOMAXPROCS programmatically.

func main() {
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}
