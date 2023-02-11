package main

import "fmt"

// A brute force algorithm solves a problem based on the statement and the problem
// definition. Brute force algorithms for search and sort are sequential search and selection
// sort. Exhaustive search is another brute force algorithm where the solution is in a set of
// candidate solutions with definitive properties. The space in which the search happens is a
// state and combinatorial space, which consists of permutations, combinations, or subsets.

// Brute Force algorithms are known for wide applicability and simplicity in solving complex
// problems. Searching, string matching, and matrix multiplication are some scenarios where
// they are used. Single computational tasks can be solved using brute force algorithms. They
// do not provide efficient algorithms. The algorithms are slow and non-performant.

func findElement(arr [10]int, k int) bool {
	var i int
	for i = 0; i < 10; i++ {
		if arr[i] == k {
			return true
		}
	}
	return false
}

func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var check bool = findElement(arr, 10)
	fmt.Println(check)
	var check2 bool = findElement(arr, 9)
	fmt.Println(check2)
}
