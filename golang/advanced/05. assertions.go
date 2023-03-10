package main

import "fmt"

// Trying to extract the concrete value from an interface using a type assertion can have two outcomes:
// - If you use the correct concrete data type, you get the underlying value without any issues
// - If you use an incorrect concrete data type, your program will panic

func returnNumber() interface{} {
	return 12
}

func main() {
	anInt := returnNumber()
	number := anInt.(int)
	number++
	fmt.Println(number)

	// The next statement would fail because there
	// is no type assertion to get the value:
	// anInt++
	// The next statement fails but the failure is under
	// control because of the ok bool variable that tells
	// whether the type assertion is successful or not
	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful: ", value)
	} else {
		fmt.Println("Type assertion failed!")
	}

	// The next statement is successful but
	// dangerous because it does not make sure that
	// the type assertion is successful.
	// It just happens to be successful
	i := anInt.(int)
	fmt.Println("i:", i)

	// The following will PANIC because anInt is not bool
	_ = anInt.(bool)
}
