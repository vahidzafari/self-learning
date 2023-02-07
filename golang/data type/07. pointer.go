package main

import "fmt"

// A pointer is the memory address of a variable. You need to dereference a pointer
// in order to get its value—dereferencing is performed using the * character in
// front of the pointer variable. Additionally, you can get the memory address of
// a normal variable using an & in front of it.
// The main benefit you get from pointers is that passing a variable to a function
// as a pointer (we can call that by reference) does not discard any changes you
// make to the value of that variable inside that function when the function returns.

// Apart from reasons of simplicity, there exist three more reasons for using pointers:
// - Pointers allow you to share data between functions.
// - Pointers are also very handy when you want to tell the difference between
// the zero value of a variable and a value that is not set (nil). This is
// particularly useful with structures because pointers (and therefore pointers
// to structures), can have the nil value, which means that you can compare a
// pointer to a structure with the nil value, which is not allowed for normal
// structure variables.
// - Having support for pointers and, more specifically, pointers to structures
// allows Go to support data structures such as linked lists and binary trees,
// which are widely used in computer science. Therefore, you are allowed to
// define a structure field of a Node structure as Next *Node, which is a pointer
// to another Node structure.

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)
	// Pointer to f
	fP := &f
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP)
	// The value of f changes
	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)
	// The value of f does not change
	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)

	// Check for empty structure
	var k *aStructure
	// This is nil because currently k points to nowhere
	fmt.Println(k)
	// Therefore you are allowed to do this:
	if k == nil {
		k = new(aStructure)
	}
	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil!")
	}
}
