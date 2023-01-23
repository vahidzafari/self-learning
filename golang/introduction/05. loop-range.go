package main

import "fmt"

// Go supports for loops as well as the range keyword for iterating over all the
// elements of arrays, slices, and maps. Go provides support for the for keyword only,
// instead of including direct support for while loops. However, depending on how you
// write a for loop, it can function as a while loop or an infinite loop. Moreover,
// for loops can implement the functionality of JavaScript's forEach function when
// combined with the range keyword.
// You can also create for loops with variables and conditions. A for loop can be exited
// with a break keyword and you can skip the current iteration with the continue
// keyword. When used with range, for loops allow you to visit all the elements of a
// slice or an array without knowing the size of the data structure.

func main() {
	// Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	// For loop used as while loop
	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// This is a slice but range also works with arrays
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value: ", v)
	}
}
