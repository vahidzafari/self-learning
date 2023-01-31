package main

import "fmt"

// When you define a new structure, you group a set of values into a single data type,
// which allows you to pass and receive this set of values as a single entity. A structure
// has fields, and each field has its own data type, which can even be another structure
// or slice of structures. Additionally, as a structure is a new data type, it is defined
// using the type keyword followed by the name of the structure and ending with the
// struct keyword, which signifies that we are defining a new structure.

type Entry struct {
	Name    string
	Surname string
	Year    int
}

// There exist two ways to work with structure variables. The first one is as regular
// variables and the second one is as pointer variables that point to the memory
// address of a structure. As a result, there exist two main ways to create a new structure
// variable using a function. The first one returns a regular structure variable whereas
// the second one returns a pointer to a structure. Each one of these two ways has two
// variations. The first variation returns a structure instance that is initialized by the Go
// compiler, whereas the second variation returns a structure instance that is initialized
// by the user.

// Additionally, you can create new structure instances using the new() keyword: pS := new(Entry).
// The new() keyword has the following properties:
// - It allocates the proper memory space, which depends on the data type, and then it zeroes it
// - It always returns a pointer to the allocated memory
// - It works for all data types except channel and map

// Initialized by Go
func zeroS() Entry {
	return Entry{}
}

// Initialized by the user
func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}
	return Entry{Name: N, Surname: S, Year: Y}
}

// Initialized by Go - returns pointer
func zeroPtoS() *Entry {
	return &Entry{}
}

// Initialized by the user - returns pointer
func initPtoS(N, S string, Y int) *Entry {
	if len(S) == 0 {
		return &Entry{Name: N, Surname: "Unknown", Year: Y}
	}
	return &Entry{Name: N, Surname: S, Year: Y}
}

func main() {
	s1 := zeroS()
	p1 := zeroPtoS()
	fmt.Println("s1:", s1, "p1:", *p1)
	s2 := initS("Mihalis", "Tsoukalos", 2020)
	p2 := initPtoS("Mihalis", "Tsoukalos", 2020)
	fmt.Println("s2:", s2, "p2:", *p2)
	fmt.Println("Year:", s1.Year, s2.Year, p1.Year, p2.Year)
	pS := new(Entry)
	fmt.Println("pS:", pS)
}
