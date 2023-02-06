package main

import (
	"fmt"
	"os"
)

var VERSION string

// The main function says that if there is a command-line argument and its value is
// version, print the version message with the help of the VERSION variable.
func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "version" {
			fmt.Println("Version:", VERSION)
		}
	}
}
