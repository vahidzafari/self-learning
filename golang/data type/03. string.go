package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// A Go string is just a collection of bytes and can be accessed as a whole or as an array.
// A single byte can store any ASCII character—however, multiple bytes are usually needed
// for storing a single Unicode character.

// A rune is an int32 value that is used for representing a single Unicode code point, which
// is an integer value that is used for representing single Unicode characters or, less
// frequently, providing formatting information.

func main() {
	aString := "Hello World! €"
	aUint := []byte("Hello World! €")
	fmt.Println("First character", string(aString[0]))
	fmt.Println("First character integer", aString[0])
	fmt.Printf("Type %T\n", aString)
	fmt.Printf("aUint Type %T\n", aUint)

	// Print an existing string as runes
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	// Print an existing string as characters
	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	// Runes
	// A rune
	r := '€'
	fmt.Println("As an int32 value:", r)
	// Convert Runes to text
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	num := 100
	input := strconv.Itoa(num)
	fmt.Println("strconv.Itoa()", input)
	input = strconv.FormatInt(int64(num), 10)
	fmt.Println("strconv.FormatInt()", input)
	input = string(num)
	fmt.Println("string()", input)

	sL := "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}
