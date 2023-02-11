package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
)

// Go supports the processing of multiple JSON records as streams instead of individual records,
// which is faster and more efficient.
// The DeSerialize() function is used for reading input in the form of JSON records,
// decoding it, and putting it into a slice. The function writes the slice, which is of
// the interface{} data type and is given as a parameter, and gets its input from the
// buffer of the *json.Decoder parameter. The *json.Decoder parameter along with its
// buffer is defined in the main() function in order to avoid allocating it all the time and
// therefore losing the performance gains and efficiency of using this type—the same
// applies to the use of *json.Encoder.

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var DataRecords []Data

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

var MIN = 0
var MAX = 26

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}

func main() {
	// Create sample data
	var i int
	var t Data
	for i = 0; i < 2; i++ {
		t = Data{
			Key: getString(5),
			Val: random(1, 100),
		}
		DataRecords = append(DataRecords, t)
	}

	// bytes.Buffer is both an io.Reader and io.Writer
	buf := new(bytes.Buffer)

	encoder := json.NewEncoder(buf)
	encoder.SetIndent("", "\t")
	err := Serialize(encoder, DataRecords)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("After Serialize:", buf)

	decoder := json.NewDecoder(buf)
	var temp []Data
	err = DeSerialize(decoder, &temp)
	fmt.Println("After DeSerialize:")
	for index, value := range temp {
		fmt.Println(index, value)
	}
}
