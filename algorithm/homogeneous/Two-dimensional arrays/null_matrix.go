package main

import "fmt"

// A null or a zero matrix is a matrix entirely consisting of zero values

func main() {

	var matrix = [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	fmt.Println(matrix)
}
