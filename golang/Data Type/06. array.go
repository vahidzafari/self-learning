package main

import (
	"fmt"
	"sort"
)

// When defining an array variable, you must define its size. Otherwise, you
// should put [...] in the array declaration and let the Go compiler find out
// the length for you. So you can create an array with 4 string elements either
// as [4]string{"Zero", "One", "Two", "Three"} or as [...]string{"Zero",
// "One", "Two", "Three"}. If you put nothing in the square brackets, then a
// slice is going to be created instead. The (valid) indexes for that particular
// array are 0, 1, 2, and 3.
// You cannot change the size of an array after you have created it.
// When you pass an array to a function, what is happening is that Go creates
// a copy of that array and passes that copy to that function—therefore any
// changes you make to an array inside a function are lost when the function
// returns.

func main() {
	// Create an empty slice
	aSlice := make([]float64, 3)

	// The capacity shows how much a slice can be expanded without the need to allocate
	// more memory and change the underlying array.
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	bSlice := []float64{}

	// Both length and capacity are 0 because bSlice is empty
	fmt.Println(bSlice, len(bSlice), cap(bSlice))

	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
	emptyTwoD := make([][]int, 2)

	fmt.Println(twoD, len(twoD))
	fmt.Println(emptyTwoD, len(emptyTwoD))

	// Visiting all elements of a 2D slice
	// with a double for loop
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	// Byte slice
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)
	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b)

	i := 2
	// Replace element at index i with last element
	aSlice[i] = aSlice[len(aSlice)-1]
	// Remove last element
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("After 2nd deletion:", aSlice)

	//	Go offers the copy() function for copying an existing array to a slice or an existing
	// slice to another slice. However, the use of copy() can be tricky because the
	// destination slice is not auto-expanded if the source slice is bigger than the destination
	// slice. Additionally, if the destination slice is bigger than the source slice, then copy()
	// does not empty the elements from the destination slice that did not get copied.

	a1 := []int{1}
	a2 := []int{2, 3}
	a5 := []int{10, 11, 12, 13, 14}
	// len(a5) > len(a1)
	copy(a1, a5)
	fmt.Println("a1", a1)
	fmt.Println("a5", a5)
	// len(a2) < len(a5) -> OK
	copy(a5, a2)
	fmt.Println("a2", a2)
	fmt.Println("a5", a5)

	// Sorting slices
	sInts := []int{1, 0, 2, -3, 4, -20}
	sFloats := []float64{1.0, 0.2, 0.22, -3, 4.1, -0.1}
	sStrings := []string{"aa", "a", "A", "Aa", "aab", "AAa"}

	fmt.Println("sInts original:", sInts)
	sort.Ints(sInts)
	fmt.Println("sInts:", sInts)
	sort.Sort(sort.Reverse(sort.IntSlice(sInts)))
	fmt.Println("Reverse:", sInts)
	fmt.Println("sFloats original:", sFloats)
	sort.Float64s(sFloats)
	fmt.Println("sFloats:", sFloats)
	sort.Sort(sort.Reverse(sort.Float64Slice(sFloats)))
	fmt.Println("Reverse:", sFloats)
	sort.Sort(sort.Reverse(sort.StringSlice(sStrings)))
	fmt.Println("Reverse:", sStrings)
}
