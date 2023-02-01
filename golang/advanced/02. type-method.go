package main

import (
	"fmt"
	"os"
	"strconv"
)

// A type method is a function that is attached to a specific data type.
// Having a data type called ar2x2, you can create a type method named FunctionName
// for it as follows:

// func (a ar2x2) FunctionName(parameters) <return values> {
// ...
// }

// The (a ar2x2) part is what makes the FunctionName() function a type method
// because it associates FunctionName() with the ar2x2 data type. No other data type
// can use that function. However, you are free to implement FunctionName() for other
// data types or as a regular function. If you have a ar2x2 variable named varAr, you
// can invoke FunctionName() as varAr.FunctionName(...), which looks like selecting
// the field of a structure variable.
// You are not obligated to develop type methods if you do not want to. In fact, each
// type method can be rewritten as a regular function. Therefore, FunctionName() can
// be rewritten as follows:

// func FunctionName(a ar2x2, parameters...) <return values> {
// ...
// }

// Have in mind that under the hood, the Go compiler does turn methods into regular
// function calls with the self value as the first parameter. However, interfaces require
// the use of type methods to work.

type ar2x2 [2][2]int

func (a *ar2x2) Add(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
}

func (a *ar2x2) Subtract(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] - b[i][j]
		}
	}
}

func (a *ar2x2) Multiply(b ar2x2) {
	a[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	a[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	a[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	a[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]
}

func main() {
	if len(os.Args) != 9 {
		fmt.Println("Need 8 integers")
		return
	}
	k := [8]int{}
	for index, i := range os.Args[1:] {
		v, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		k[index] = v
	}
	a := ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b := ar2x2{{k[4], k[5]}, {k[6], k[7]}}
	a.Add(b)
	fmt.Println("a+b:", a)
	a.Subtract(a)
	fmt.Println("a-a:", a)
	a = ar2x2{{k[0], k[1]}, {k[2], k[3]}}

	a.Multiply(b)
	fmt.Println("a*b:", a)
	a = ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b.Multiply(a)
	fmt.Println("b*a:", b)
}

// type T1 struct {
// 	Firstname string
// }

// func (t T1) set() {
// 	t.Firstname = "*Vahid"
// }

// func (t *T1) setPtr() {
// 	t.Firstname = "Vahid"
// }

// func (t *T1) print() {
// 	fmt.Println("Firstname: ", t.Firstname)
// }

// func main() {
// 	t1 := new(T1)
// 	t1.set()
// 	t1.print()
// 	t1.setPtr()
// 	t1.print()
// }
