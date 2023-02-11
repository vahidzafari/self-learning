package main

// An algorithm is of linear complexity if the processing time or storage space is directly
// proportional to the number of input elements to be processed. In Big O notation, linear
// complexity is presented as O(n). String matching algorithms such as the Boyer-Moore and
// Ukkonen have linear complexity.

import (
	"fmt"
)

func main() {
	var m [10]int
	var k int
	for k = 0; k < 10; k++ {
		m[k] = k * 200
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}
