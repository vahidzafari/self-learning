package main

import "fmt"

// A type assertion is a mechanism for working with the underlying concrete value of
// an interface. This mainly happens because interfaces are virtual data types without
// their own values—interfaces just define behavior and do not hold data of their own.
// Type switches use switch blocks for data types and allow you to differentiate between
// type assertion values, which are data types, and process each data type the way you want.
// On the other hand, in order to use the empty interface in type switches, you need to use
// type assertions.

// Type assertions use the x.(T) notation, where x is an interface type and T is a type,
// and help you extract the value that is hidden behind the empty interface. For a type
// assertion to work, x should not be nil and the dynamic type of x should be identical
// to the T type.

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func Teststruct(x interface{}) {
	// type switch
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret type")
	case Entry:
		fmt.Println("Entry type")
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}

func main() {
	A := Entry{100, "F2", Secret{"myPassword"}}
	Teststruct(A)
	Teststruct(A.F3)
	Teststruct("A string")
	Learn(12.23)
	Learn('€')
}
