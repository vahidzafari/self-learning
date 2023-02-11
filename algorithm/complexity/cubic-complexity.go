package main

// In the case of cubic complexity, the processing time of an algorithm is proportional to the
// cube of the input elements. The complexity of the following algorithm is 10*10*10 = 1,000.
// The three loops have a maximum of 10. The cubic complexity for a matrix update is O(n^3).

import (
	"fmt"
)

func main() {
	var k, l, m int
	var arr [10][10][10]int

	for k = 0; k < 10; k++ {
		for l = 0; l < 10; l++ {
			for m = 0; m < 10; m++ {
				arr[k][l][m] = 1
				fmt.Println("Element value ", k, l, m, " is", arr[k][l][m])
			}
		}
	}
}
