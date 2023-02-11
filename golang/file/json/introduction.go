package main

import (
	"encoding/json"
	"fmt"
)

// The Go standard library includes encoding/json, which is for working with JSON data.
// Additionally, Go allows you to add support for JSON fields in Go structures using
// tags, which is the subject of the Structures and JSON subsection. Tags control
// the encoding and decoding of JSON records to and from Go structures.

// ## Using Marshal() and Unmarshal()
// Both the marshaling and unmarshaling of JSON data are important procedures for working with
// JSON data using Go structures. Marshaling is the process of converting a Go structure into
// a JSON record. You usually want that for transferring JSON data via computer networks or
// for saving it on disk. Unmarshaling is the process of converting a JSON record given as
// a byte slice into a Go structure.

// Imagine that you have a Go structure that you want to convert into a JSON record
// without including any empty fields— you can do that with the use of omitempty

// Last, imagine that you have some sensitive data on some of the fields of a Go
// structure that you do not want to include in the JSON records. You can do that by
// including the "-" special value in the desired json: structure tags.

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"creationyear,omitempty"`
	Pass    string `json:"-"`
}

func main() {
	useall := UseAll{Name: "Mike", Surname: "Tsoukalos", Year: 2021}
	// Regular Structure
	// Encoding JSON data -> Convert Go Structure to JSON record with fields
	t, err := json.Marshal(&useall)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t)
	}
	// Decoding JSON data given as a string
	str := `{"username": "M.", "surname": "Ts", "created":2020}`
	// Convert string into a byte slice
	jsonRecord := []byte(str)
	// Create a structure variable to store the result
	temp := UseAll{}
	err = json.Unmarshal(jsonRecord, &temp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type: %T with value %v\n", temp, temp)
	}
}
