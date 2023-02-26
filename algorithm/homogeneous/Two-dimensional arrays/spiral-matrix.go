package main

import "fmt"

// A spiral matrix is an arrangement of n x n integers in which integers are arranged spirally
// in sequentially increasing order. A spiral matrix is an old toy algorithm. The spiral order is
// maintained using four loops, one for each corner of the matrix.

func PrintSpiral(n int) []int {
	var left int
	var top int
	var right int
	var bottom int
	left = 0
	top = 0
	right = n - 1
	bottom = n - 1
	var size int
	size = n * n
	var s []int
	s = make([]int, size)
	var i int
	i = 0
	for left < right {
		var c int
		for c = left; c <= right; c++ {
			s[top*n+c] = i
			i++
		}
		top++
		var r int
		for r = top; r <= bottom; r++ {
			s[r*n+right] = i
			i++
		}
		right--
		if top == bottom {
			break
		}
		for c = right; c >= left; c-- {
			s[bottom*n+c] = i
			i++
		}
		bottom--
		for r = bottom; r >= top; r-- {
			s[r*n+left] = i
			i++
		}
		left++
	}
	s[top*n+left] = i
	return s
}

func main() {
	var n int
	n = 5
	var length int
	length = 2
	var i int

	var sketch int
	for i, sketch = range PrintSpiral(n) {
		fmt.Printf("%*d ", length, sketch)
		if i%n == n-1 {
			fmt.Println("")
		}
	}
}
