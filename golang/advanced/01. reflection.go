package main

import (
	"fmt"
	"reflect"
)

// You might be wondering how you can find out the names of the fields of a structure
// at execution time. In such cases, you need to use reflection. Apart from enabling you
// to print the fields and the values of a structure, reflection also allows you to explore
// and manipulate unknown structures like the ones created from decoding JSON data.

// Reflection allows you to dynamically learn the type of an arbitrary object along with
// information about its structure. Go provides the reflect package for working with reflection.
// Now, reflect.Value is used for storing values of any type, whereas reflect.Type is used
// for representing Go types. There exist two functions named reflect.TypeOf() and reflect.ValueOf()
// that return the reflect.Type and reflect.Value values, respectively. Note that reflect.TypeOf()
// returns the actual type of variable—if we are examining a structure, it returns the
// name of the structure.

// As structures are really important in Go, the reflect package offers the reflect.
// NumField() method for listing the number of fields in a structure as well as the
// Field() method for getting the reflect.Value value of a specific field of a structure.
// The reflect package also defines the reflect.Kind data type, which is used for
// representing the specific data type of a variable: int, struct, etc. The documentation
// of the reflect package lists all possible values of the reflect.Kind data type. The
// Kind() function returns the kind of a variable.

type Secret struct {
	Username string
	Password string
}
type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"Mihalis",
		"Tsoukalos"}}

	r := reflect.ValueOf(A)
	fmt.Println("String value:", r.String())

	iType := r.Type()

	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)
	for i := 0; i < r.NumField(); i++ {

		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())

		// Check whether there are other structures embedded in Record
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		// Need to convert it to string in order to compare it
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}
		// Same as before but using the internal value
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}
}
