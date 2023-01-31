package main

import "fmt"

// Both arrays and slices limit you to using positive integers as indexes. Maps are
// powerful data structures because they allow you to use indexes of various data types
// as keys to look up your data as long as these keys are comparable. A practical rule of
// thumb is that you should use a map when you are going to need indexes that are not
// positive integer numbers or when the integer indexes have big gaps.

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("aMap:", aMap)
	aMap = nil
	fmt.Println("aMap:", aMap)
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}

	// range works with maps as well
	for key, v := range aMap {
		fmt.Println("key:", key, "value:", v)
	}

	aMap["test"] = 1
	fmt.Println("aMap:", aMap)
	fmt.Println("aMap len:", len(aMap))
	delete(aMap, "test")
	fmt.Println("aMap:", aMap)
}
