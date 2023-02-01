package main

import (
	"fmt"
	"sort"
)

// An interface is a Go mechanism for defining behavior that is implemented using a
// set of methods. Interfaces play a key role in Go and can simplify the code of your
// programs when they have to deal with multiple data types that perform the same
// task—recall that fmt.Println() works for almost all data types. But remember,
// interfaces should not be unnecessarily complex. If you decide to create your own
// interfaces, then you should begin with a common behavior that you want to be used
// by multiple data types.

// Interfaces work with methods on types (or type methods), which are like functions
// attached to given data types, which in Go are usually structures (although we can
// use any data type we want).
// The empty interface is defined as just interface{}. As the empty interface has no
// methods, it means that it is already implemented by all data types.

// In a more formal way, a Go interface type defines (or describes) the behavior of
// other types by specifying a set of methods that need to be implemented for supporting
// that behavior. For a data type to satisfy an interface, it needs to implement all the
// type methods required by that interface. Therefore, interfaces are abstract types
// that specify a set of methods that need to be implemented so that another type
// can be considered an instance of the interface. So, an interface is two things: a set
// of methods and a type.

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type S1 struct {
	F1 int
	F2 string
	F3 int
}

// We want to sort S2 records based on the value of F3.F1,
// Which is equivalent to S1.F1 as F3 is an S1 structure
type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2slice []S2

// Implementing sort.Interface for S2slice
func (a S2slice) Len() int {
	return len(a)
}

// What field to use when comparing
func (a S2slice) Less(i, j int) bool {
	return a[i].F3.F1 < a[j].F3.F1
}

func (a S2slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	data := []S2{
		S2{1, "One", S1{1, "S1_1", 10}},
		S2{2, "Two", S1{2, "S1_1", 20}},
		S2{-1, "Two", S1{-1, "S1_1", -20}},
	}
	fmt.Println("Before:", data)
	sort.Sort(S2slice(data))
	fmt.Println("After:", data)
	// Reverse sorting works automatically
	sort.Sort(sort.Reverse(S2slice(data)))
	fmt.Println("Reverse:", data)
}
