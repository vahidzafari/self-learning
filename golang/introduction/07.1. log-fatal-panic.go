package main

import (
	"log"
	"os"
)

// The log.Fatal() function is used when something erroneous has happened and you
// just want to exit your program as soon as possible after reporting that bad situation.
// The call to log.Fatal() terminates a Go program at the point where log.Fatal() was
// called after printing an error message. In most cases, this custom error message can
// be Not enough arguments, Cannot access file, or similar. Additionally, it returns
// back a non-zero exit code, which in UNIX indicates an error.

// There are situations where a program is about to fail for good and you want to
// have as much information about the failure as possible—log.Panic() implies that
// something really unexpected and unknown, such as not being able to find a file that
// was previously accessed or not having enough disk space, has happened. Analogous
// to the log.Fatal() function, log.Panic() prints a custom message and immediately
// terminates the Go program.

// Have in mind that log.Panic() is equivalent to a call to log.Print() followed
// by a call to panic(). panic() is a built-in function that stops the execution of the
// current function and begins panicking. After that, it returns to the caller function.
// On the other hand, log.Fatal() calls log.Print() and then os.Exit(1), which is an
// immediate way of terminating the current program.

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello World!")
	}
	log.Panic("Panic: Hello World!")
}
