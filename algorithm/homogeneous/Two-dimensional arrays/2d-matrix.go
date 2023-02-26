package main

import "fmt"

var matrix1 = [2][2]int{
	{4, 5},
	{1, 2},
}
var matrix2 = [2][2]int{
	{6, 7},
	{3, 4},
}

func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var n int
	var sum [2][2]int
	for m = 0; m < 2; m++ {
		for n = 0; n < 2; n++ {
			sum[n][m] = matrix1[n][m] + matrix2[n][m]
		}
	}
	return sum
}

func subtract(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var n int
	var sum [2][2]int
	for m = 0; m < 2; m++ {
		for n = 0; n < 2; n++ {
			sum[n][m] = matrix1[n][m] - matrix2[n][m]
		}
	}
	return sum
}

func multiply(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var n int
	var product [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			var productSum int = 0
			for n = 0; n < 2; n++ {
				productSum = productSum + matrix1[l][n]*matrix2[n][m]
			}
			product[l][m] = productSum
		}
	}
	return product
}

func transpose(matrix1 [2][2]int) [2][2]int {
	var m int
	var l int
	var transMatrix [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			transMatrix[l][m] = matrix1[m][l]
		}
	}
	return transMatrix
}

func determinant(matrix1 [2][2]int) float64 {

	var det float64

	det = det + float64(((matrix1[0][0] * matrix1[1][1]) - (matrix1[0][1] * matrix1[1][0])))

	return det
}

func inverse(matrix [2][2]int) [2][2]float64 {

	var det float64

	det = determinant(matrix)

	var invmatrix [2][2]float64
	invmatrix[0][0] = float64(matrix[1][1]) / det
	invmatrix[0][1] = float64(-1*matrix[0][1]) / det
	invmatrix[1][0] = float64(-1*matrix[1][0]) / det
	invmatrix[1][1] = float64(matrix[0][0]) / det

	return invmatrix

}

func main() {

	var matrix1 = [2][2]int{
		{4, 5},
		{1, 2}}
	var matrix2 = [2][2]int{
		{6, 7},
		{3, 4}}
	var sum [2][2]int
	sum = add(matrix1, matrix2)

	fmt.Println(sum)

	var difference [2][2]int
	difference = subtract(matrix1, matrix2)

	fmt.Println(difference)

	var product [2][2]int
	product = multiply(matrix1, matrix2)

	fmt.Println("product", product)

	determinantResult := determinant(product)
	fmt.Println("determinantResult", determinantResult)

	inverseResult := inverse(product)
	fmt.Println("inverseResult", inverseResult)
}
