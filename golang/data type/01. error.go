package main

import (
	"errors"
	"fmt"
	"os"
)

// Go provides a special data type for representing error conditions and error messages
// named error—in practice, this means that Go treats errors as values.
// Go follows the next convention about error values: if the value of an error variable is nil,
// then there was no error. As an example, let us consider strconv.Atoi(), which is used for
// converting a string value into an int value (Atoi stands for ASCII to Int). As specified by
// its signature, strconv.Atoi() returns (int, error). Having an error value of nil means that
// the conversion was successful and that you can use the int value if you want. Having an error
// value that is not nil means that the conversion was unsuccessful and that the string input
// is not a valid int value.

// Custom error message with errors.New()
func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is a custom error message")
	}
	return nil
}

// Custom error message with fmt.Errorf()
func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b,
			os.Getuid())
	}
	return nil
}

func main() {
	err := check(0, 0)
	if err.Error() == "this is a custom error message" {
		fmt.Println("Custom error detected!")
	}
	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}
}
