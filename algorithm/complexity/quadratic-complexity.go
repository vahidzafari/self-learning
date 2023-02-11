package main

// An algorithm is of quadratic complexity if the processing time is proportional to the square
// of the number of input elements. In the following case, the complexity of the algorithm is
// 10*10 = 100. The two loops have a maximum of 10. The quadratic complexity for a
// multiplication table of n elements is O(n^2).

import (
	"fmt"
)

func main() {
	var k, l int
	for k = 1; k <= 10; k++ {
		fmt.Println(" Multiplication Table", k)
		for l = 1; l <= 10; l++ {
			var x int = l * k
			fmt.Println(x)
		}
	}
}
