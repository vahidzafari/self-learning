package main

import (
	"fmt"
)

// You can declare a new variable using the var keyword followed by the variable name,
// followed by the desired data type. If you want, you can follow that declaration with
// = and an initial value for your variable. If there is an initial value given, you can
// omit the data type and the compiler will guess it for you.

// var <identifier> <type>;
// var <identifier> = <value>;
// var <identifier> <type> = <value>;

// if no initial value is given to a variable, the Go compiler will automatically initialize
// that variable to the zero value of its data type.

// There is also the := notation, which can be used instead of a var declaration. :=
// defines a new variable by inferring the data of the value that follows it.
// The short assignment statement can be used in place of a var declaration with an
// implicit type. the var keyword is mostly used for  declaring global or local variables
// without an initial value. The reason for the former is that every statement that exists
// outside of the code of a function must begin with a keyword such as func or var.
// Last, you might need to use var when you want to be explicit about the data type.

var Global int = 1234
var AnotherGlobal = -5678

func main() {
	i := Global + AnotherGlobal
	fmt.Printf("Global=%d, i=%d\n", Global, i)
}
