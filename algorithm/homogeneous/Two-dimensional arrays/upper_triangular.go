package main

import "fmt"

// An upper triangular matrix consists of elements with a value of zero below the main
// diagonal.

func main() {

	var matrix = [3][3]int{
		{1, 1, 1},
		{0, 1, 1},
		{0, 0, 1},
	}

	fmt.Println(matrix)
}
