package main

import "fmt"

// A backtracking algorithm solves a problem by constructing the solution incrementally.
// Multiple options are evaluated, and the algorithm chooses to go to the next component of
// the solution through recursion. Backtracking can be a chronological type or can traverse the
// paths, depending on the problem that you are solving.

// Backtracking is an algorithm that finds candidate solutions and rejects a candidate on the
// basis of its feasibility and validity. Backtracking is useful in scenarios such as finding a
// value in an unordered table. It is faster than a brute force algorithm, which rejects a large
// number of solutions in an iteration. Constraint satisfaction problems such as parsing, rules
// engine, knapsack problems, and combinatorial optimization are solved using backtracking.

// The following is an example of a backtracking algorithm. The problem is to identify the
// combinations of elements in an array of 10 elements whose sum is equal to 18. The
// findElementsWithSum method recursively tries to find the combination. Whenever the
// sum goes beyond the k target, it backtracks.

func findElementsWithSum(arr [10]int, combinations [19]int, size int, k int,
	addValue int, l int, m int) int {
	var num int = 0
	if addValue > k {
		return -1
	}
	if addValue == k {
		num = num + 1
		var p int = 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}
	var i int
	for i = l; i < size; i++ {
		//fmt.Println(" m", m)
		combinations[m] = l
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return num
}

func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var addedSum int = 18
	var combinations [19]int
	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
	// fmt.Println(check)//var check2 bool = findElement(arr,9)
	// fmt.Println(check2)
}
