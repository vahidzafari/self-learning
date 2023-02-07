package main

import "fmt"

// A variable declaration can be done at the package level or a function level or a block level. The scope of a variable defines from where that variable can be accessed and also the lifetime of the variable. Golang variables can be divided into two categories based on the scope.
// -Local Variable
// -Global Variable

// Local Variable:
// -Local variables are variables which are defined within a block or a function level
// -An example of a block is a for loop or a range loop etc.
// -These variables are only be accessed from within their block or function
// -These variables only live till the end of the block or a function in which they are declared. After that,
// they are Garbage Collected.
// -A local once declared cannot be redeclared within the same block or function.

// Global Variable
// -A variable will be global within a package if it is declared at the top of a file outside the scope
// of any function or block.
// -If this variable name starts with a lowercase letter then it can be accessed from within the package
// which contains this variable definition.
// -If the variable name starts with an uppercase letter then it can be accessed from an outside different
// package other than which it is declared.
// -Global variable are available throughout the lifetime of a program

var global_variable = "global_variable"

func main() {
	testGlobal()
}

func testGlobal() {
	fmt.Println("global_variable", global_variable)

	var local_variable = 456
	fmt.Println("local_variable", local_variable)
}
