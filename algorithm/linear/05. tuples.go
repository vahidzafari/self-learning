package main

import "fmt"

// A tuple is a finite sorted list of elements. It is a data structure that groups data. Tuples are
// typically immutable sequential collections. The element has related fields of different
// datatypes. The only way to modify a tuple is to change the fields. Operators such as + and *
// can be applied to tuples. A database record is referred to as a tuple.
// A tuple can be assigned in a single statement, which is useful for swapping values.
// Lists usually contain values of the same data type, while tuples contain different data.

// gets the power series of integer a and returns tuple of square of a and cube of a
func powerSeries(a int) (int, int, error) {
	square := a * a
	cube := square * a
	return square, cube, nil
}

func main() {
	var square int
	var cube int
	square, cube, _ = powerSeries(3)
	fmt.Println("Square ", square, "Cube", cube)
}
