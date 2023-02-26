package main

import "fmt"

// A lower triangular matrix consists of elements that have a value of zero above the main
// diagonal.

func main() {

	var matrix = [3][3]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 1},
	}

	fmt.Println(matrix)
}
