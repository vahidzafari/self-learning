package main

import "fmt"

// Go supports integer, floating-point, and complex number values in various versions
// depending on the memory space they consume—this saves memory and computing
// time. Integer data types can be either signed or unsigned, which is not the case for
// floating point numbers.

func main() {
	// Complex
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	fmt.Printf("Type of c3: %T\n", c3)

	x := 12
	k := 5
	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x)

	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)

	divFloat := float64(x) / float64(k)
	fmt.Println("divFloat", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)
}
