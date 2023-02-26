package main

import "fmt"

// A Boolean matrix is a matrix that consists of elements in the mth row and the nth column with
// a value of 1. A matrix can be modified to be a Boolean matrix by making the values in
// the mth row and the nth column equal to 1. The changeMatrix method
// transforms the input matrix in to a Boolean matrix by changing the row and column values
// from 0 to 1 if the cell value is 1. The method takes the input matrix as the parameter and
// returns the changed matrix.

func changeMatrix(matrix [3][3]int) [3][3]int {
	var i int
	var j int
	var Rows [3]int
	var Columns [3]int
	var matrixChanged [3][3]int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if matrix[i][j] == 1 {
				Rows[i] = 1
				Columns[j] = 1
			}
		}
	}
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if Rows[i] == 1 || Columns[j] == 1 {
				matrixChanged[i][j] = 1
			}
		}
	}
	return matrixChanged
}

func printMatrix(matrix [3][3]int) {
	var i int
	var j int
	//var k int
	for i = 0; i < 3; i++ {

		for j = 0; j < 3; j++ {

			fmt.Printf("%d", matrix[i][j])

		}
		fmt.Printf("\n")
	}

}

func main() {

	var matrix = [3][3]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	printMatrix(matrix)

	matrix = changeMatrix(matrix)

	printMatrix(matrix)

}
