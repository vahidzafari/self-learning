package main

import "fmt"

// Go supports constants, which are variables that cannot change their values.
// In Go const can be either of type string, numeric, boolean, and characters.
// Constants in Go are defined with the help of the const keyword. Generally speaking,
// constants can be either global or local variables. The main benefit you get from using
// constants in your programs is the guarantee that their value will not change during
// program execution. Strictly speaking, the value of a constant variable is defined at
// compile time, not at runtime—this means that it is included in the binary executable.
// A constant declared within an inner having a same name as constant declared in the outer
// scope will shadow the constant in outer scope.

// const <identifier> <type> = <value>;

type Digit int
type Power2 int

const GLOBAL = 123
const PI = 3.1415926
const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

// The constant generator iota is used for declaring a sequence of related values that
// use incrementing numbers without the need to explicitly type each one of them.

func main() {
	const (
		Zero  Digit = iota // 0
		One                // 1
		Two                // 2
		Three              // 3
		Four               // 4
	)

	const (
		p2_0 Power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)

	fmt.Println("GLOBAL:", GLOBAL)
	const GLOBAL = 456
	fmt.Println("GLOBAL:", GLOBAL)
}
