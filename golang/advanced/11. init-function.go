package main

import "fmt"

// Each Go package can optionally have a private function named init() that is
// automatically executed at the beginning of execution time—init() runs when the
// package is initialized at the beginning of program execution.
// The fact that the init() function is a private function by design means that it cannot
// be called from outside the package in which it is contained. The init() function
// has the following characteristics:
// - init() takes no arguments.
// - init() returns no values.
// - The init() function is optional.
// - The init() function is called implicitly by Go.
// - You can have an init() function in the main package. In that case, init() is
// executed before the main() function. In fact, all init() functions are always
// executed prior to the main() function.
// - A source file can contain multiple init() functions—these are executed in
// the order of declaration.
// - The init() function or functions of a package are executed only once, even
// if the package is imported multiple times.
// - Go packages can contain multiple files. Each source file can contain one or
// more init() functions.

// There are some exceptions where the use of init() makes sense:
// - For initializing network connections that might take time prior to the
// execution of package functions or methods.
// - For initializing connections to one or more servers prior to the execution of
// package functions or methods.
// - For creating required files and directories.
// - For checking whether required resources are available or not.

// As an example, if a main package imports package A and package A depends on package B,
// then the following will take place:
// - The process starts with main package.
// - The main package imports package A.
// - Package A imports package B.
// - The global variables, if any, in package B are initialized.
// - The init() function or functions of package B, if they exist, run. This is the
// first init() function that gets executed.
// - The global variables, if any, in package A are initialized.
// - The init() function or functions of package A, if there are any, run.
// - The global variables in the main package are initialized.
// - The init() function or functions of main package, if they exist, run.
// - The main() function of the main package begins execution.

func init() {
	fmt.Println("init function")
}

func main() {
	fmt.Println("main function")
}
